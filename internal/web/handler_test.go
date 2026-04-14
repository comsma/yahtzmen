package web_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/comsma/yahtzmen/internal/storage"
	"github.com/comsma/yahtzmen/internal/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v5"
	"github.com/moby/moby/api/types/container"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	mysqlmodule "github.com/testcontainers/testcontainers-go/modules/mysql"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

func run(m *testing.M) int {
	ctx := context.Background()

	ctr, err := mysqlmodule.Run(ctx, "mysql:8.0",
		mysqlmodule.WithDatabase("yahtzmen"),
		mysqlmodule.WithUsername("root"),
		mysqlmodule.WithPassword("password"),
		testcontainers.CustomizeRequest(testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				HostConfigModifier: func(hc *container.HostConfig) {
					hc.AutoRemove = true
				},
			},
		}),
	)
	if err != nil {
		fmt.Printf("container start: %v\n", err)
		return 1
	}
	defer ctr.Terminate(ctx)

	dsn, err := ctr.ConnectionString(ctx, "parseTime=true")
	if err != nil {
		fmt.Printf("dsn: %v\n", err)
		return 1
	}

	testDB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("db open: %v\n", err)
		return 1
	}
	defer testDB.Close()

	if err := runMigrations(testDB); err != nil {
		fmt.Printf("migrations: %v\n", err)
		return 1
	}

	return m.Run()
}

func runMigrations(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	return goose.Up(db, "../../migrations")
}

func newTestServer() *echo.Echo {
	e := echo.New()
	store, _ := storage.New("localhost", "test-bucket", "http://localhost/test-bucket", "", "")
	h := web.NewHandler(testDB, store)
	h.RegisterRoutes(e)
	return e
}

func TestHome(t *testing.T) {
	e := newTestServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", rec.Code)
	}
	var body map[string]interface{}
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if _, ok := body["products"]; !ok {
		t.Fatal("missing products key")
	}
}

func TestProductNotFound(t *testing.T) {
	e := newTestServer()
	req := httptest.NewRequest(http.MethodGet, "/products/doesnotexist", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("want 404, got %d", rec.Code)
	}
}

func TestAdminCreateAndGetProduct(t *testing.T) {
	e := newTestServer()

	body := `{"name":"Test Hat","price":"49.99","description":"A fine hat"}`
	req := httptest.NewRequest(http.MethodPost, "/admin/products", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("create: want 201, got %d — %s", rec.Code, rec.Body.String())
	}

	var created map[string]string
	json.NewDecoder(rec.Body).Decode(&created)
	id := created["id"]
	if id == "" {
		t.Fatal("no id in response")
	}

	req = httptest.NewRequest(http.MethodGet, "/products/"+id, nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("get: want 200, got %d", rec.Code)
	}
	var product map[string]interface{}
	json.NewDecoder(rec.Body).Decode(&product)
	if product["id"] != id {
		t.Fatalf("want id %s, got %v", id, product["id"])
	}
}

func TestAdminUpdateProduct(t *testing.T) {
	e := newTestServer()

	createBody := `{"name":"Update Me","price":"10.00","description":"before"}`
	req := httptest.NewRequest(http.MethodPost, "/admin/products", strings.NewReader(createBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	var created map[string]string
	json.NewDecoder(rec.Body).Decode(&created)
	id := created["id"]

	updateBody := `{"name":"Updated","price":"20.00","description":"after"}`
	req = httptest.NewRequest(http.MethodPut, "/admin/products/"+id, strings.NewReader(updateBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("want 200, got %d — %s", rec.Code, rec.Body.String())
	}
}

func TestAdminDeleteProduct(t *testing.T) {
	e := newTestServer()

	createBody := `{"name":"Delete Me","price":"5.00"}`
	req := httptest.NewRequest(http.MethodPost, "/admin/products", strings.NewReader(createBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	var created map[string]string
	json.NewDecoder(rec.Body).Decode(&created)
	id := created["id"]

	req = httptest.NewRequest(http.MethodDelete, "/admin/products/"+id, nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("want 204, got %d", rec.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/products/"+id, nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("want 404 after delete, got %d", rec.Code)
	}
}

func TestAdminListOrders(t *testing.T) {
	e := newTestServer()
	req := httptest.NewRequest(http.MethodGet, "/admin/orders", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", rec.Code)
	}
	var body map[string]interface{}
	json.NewDecoder(rec.Body).Decode(&body)
	if _, ok := body["orders"]; !ok {
		t.Fatal("missing orders key")
	}
}

func TestAdminGetOrderNotFound(t *testing.T) {
	e := newTestServer()
	req := httptest.NewRequest(http.MethodGet, "/admin/orders/doesnotexist", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("want 404, got %d", rec.Code)
	}
}
