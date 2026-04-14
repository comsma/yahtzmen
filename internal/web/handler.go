package web

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/comsma/yahtzmen/internal/models"
	"github.com/comsma/yahtzmen/internal/storage"
	"github.com/comsma/yahtzmen/internal/utils"
	"github.com/labstack/echo/v5"
)

type Handler struct {
	q       *models.Queries
	storage *storage.Store
}

func NewHandler(db *sql.DB, store *storage.Store) *Handler {
	return &Handler{q: models.New(db), storage: store}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.Home)
	e.GET("/products/:id", h.Product)
	e.GET("/privacy", h.Privacy)

	admin := e.Group("/admin")
	admin.GET("/products", h.AdminListProducts)
	admin.POST("/products", h.AdminCreateProduct)
	admin.GET("/products/:id/images", h.AdminImageForm)
	admin.POST("/products/:id/images", h.AdminUploadImages)
	admin.PUT("/products/:id", h.AdminUpdateProduct)
	admin.DELETE("/products/:id", h.AdminDeleteProduct)
	admin.GET("/orders", h.AdminListOrders)
	admin.GET("/orders/:id", h.AdminGetOrder)
}

func (h *Handler) imageURLs(ctx context.Context, productID string) []string {
	images, _ := h.q.GetProductImages(ctx, productID)
	urls := make([]string, len(images))
	for i, img := range images {
		urls[i] = h.storage.PublicURL(img.ObjectKey)
	}
	return urls
}

// GET /
func (h *Handler) Home(c *echo.Context) error {
	rows, err := h.q.GetAllProducts(c.Request().Context())
	if err != nil {
		return err
	}

	type product struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Price string `json:"price"`
		Image string `json:"image"`
	}

	products := make([]product, len(rows))
	for i, r := range rows {
		image := ""
		if urls := h.imageURLs(c.Request().Context(), r.ID); len(urls) > 0 {
			image = urls[0]
		}
		products[i] = product{ID: r.ID, Name: r.Name, Price: r.Price, Image: image}
	}

	return c.Render(http.StatusOK, "home.gohtml", map[string]interface{}{"Products": products})
}

func (h *Handler) Privacy(c *echo.Context) error {

	return c.Render(http.StatusOK, "privacy.gohtml", nil)
}

// GET /products/:id
func (h *Handler) Product(c *echo.Context) error {
	row, err := h.q.GetProduct(c.Request().Context(), c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	type product struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Price       string   `json:"price"`
		Notes       string   `json:"notes"`
		Dimensions  string   `json:"dimensions"`
		Features    []string `json:"features"`
		Images      []string `json:"images"`
	}

	var features []string
	if row.Features.Valid && row.Features.String != "" {
		features = strings.Split(row.Features.String, ", ")
	}

	p := product{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description.String,
		Price:       row.Price,
		Notes:       row.Notes.String,
		Dimensions:  row.Dimensions.String,
		Features:    features,
		Images:      h.imageURLs(c.Request().Context(), row.ID),
	}
	return c.Render(http.StatusOK, "product.gohtml", map[string]interface{}{
		"Product": p,
	})
}

// GET /admin/products
func (h *Handler) AdminListProducts(c *echo.Context) error {
	rows, err := h.q.GetAllProducts(c.Request().Context())
	if err != nil {
		return err
	}

	type product struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Price string `json:"price"`
	}

	products := make([]product, len(rows))
	for i, r := range rows {
		products[i] = product{ID: r.ID, Name: r.Name, Price: r.Price}
	}

	return c.Render(http.StatusOK, "admin/products.gohtml", map[string]interface{}{"products": products})
}

// POST /admin/products
func (h *Handler) AdminCreateProduct(c *echo.Context) error {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       string `json:"price"`
		Notes       string `json:"notes"`
		Dimensions  string `json:"dimensions"`
		Features    string `json:"features"`
	}
	if err := c.Bind(&body); err != nil {
		return echo.ErrBadRequest
	}

	id, err := utils.GenerateId()
	if err != nil {
		return err
	}

	err = h.q.CreateProduct(c.Request().Context(), models.CreateProductParams{
		ID:          id,
		Name:        body.Name,
		Price:       body.Price,
		Description: sql.NullString{String: body.Description, Valid: body.Description != ""},
		Notes:       sql.NullString{String: body.Notes, Valid: body.Notes != ""},
		Dimensions:  sql.NullString{String: body.Dimensions, Valid: body.Dimensions != ""},
		Features:    sql.NullString{String: body.Features, Valid: body.Features != ""},
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": id})
}

// GET /admin/products/:id/images
func (h *Handler) AdminImageForm(c *echo.Context) error {
	return c.Render(http.StatusOK, "admin", map[string]interface{}{
		"ProductID": c.Param("id"),
		"Uploaded":  false,
	})
}

// POST /admin/products/:id/images
func (h *Handler) AdminUploadImages(c *echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return echo.ErrBadRequest
	}

	files := form.File["images"]
	if len(files) == 0 {
		return echo.ErrBadRequest
	}

	var urls []string
	for _, fh := range files {
		objectKey, err := utils.GenerateId()
		if err != nil {
			return err
		}

		f, err := fh.Open()
		if err != nil {
			return err
		}
		defer f.Close()

		if err := h.storage.Upload(c.Request().Context(), objectKey, f, fh.Size, fh.Header.Get("Content-Type")); err != nil {
			return err
		}

		if err := h.q.CreateProductImage(c.Request().Context(), models.CreateProductImageParams{
			ID:        objectKey,
			ProductID: c.Param("id"),
			ObjectKey: objectKey,
		}); err != nil {
			return err
		}

		urls = append(urls, h.storage.PublicURL(objectKey))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"images": urls})
}

// PUT /admin/products/:id
func (h *Handler) AdminUpdateProduct(c *echo.Context) error {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       string `json:"price"`
		Notes       string `json:"notes"`
		Dimensions  string `json:"dimensions"`
		Features    string `json:"features"`
	}
	if err := c.Bind(&body); err != nil {
		return echo.ErrBadRequest
	}

	err := h.q.UpdateProduct(c.Request().Context(), models.UpdateProductParams{
		ID:          c.Param("id"),
		Name:        body.Name,
		Price:       body.Price,
		Description: sql.NullString{String: body.Description, Valid: body.Description != ""},
		Notes:       sql.NullString{String: body.Notes, Valid: body.Notes != ""},
		Dimensions:  sql.NullString{String: body.Dimensions, Valid: body.Dimensions != ""},
		Features:    sql.NullString{String: body.Features, Valid: body.Features != ""},
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// DELETE /admin/products/:id
func (h *Handler) AdminDeleteProduct(c *echo.Context) error {
	if err := h.q.DeleteProduct(c.Request().Context(), c.Param("id")); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// GET /admin/orders
func (h *Handler) AdminListOrders(c *echo.Context) error {
	rows, err := h.q.GetAllOrders(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"orders": rows})
}

// GET /admin/orders/:id
func (h *Handler) AdminGetOrder(c *echo.Context) error {
	row, err := h.q.GetOrder(c.Request().Context(), c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, row)
}
