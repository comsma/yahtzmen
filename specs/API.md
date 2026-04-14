# API Specification

## Public Routes

### GET /
Home page.

**Response:** `200 OK`
```json
{
  "products": [{ "id": "string", "name": "string", "price": "number", "image": "string" }]
}
```

---

### GET /products/:id
Product detail page.

**Params:**
- `id` — product ID

**Response:** `200 OK`
```json
{
  "id": "string",
  "name": "string",
  "description": "string",
  "price": "number",
  "stock": "number",
  "images": "string[]"
}
```

**Errors:**
- `404` — product not found

---

## Admin Routes

> Requires admin authentication.

### GET /admin/products
List all products.

**Response:** `200 OK`
```json
{
  "products": [
    { "id": "string", "name": "string", "price": "number", "stock": "number" }
  ]
}
```

---

### POST /admin/products
Create a product.

**Body:**
```json
{
  "name": "string",
  "description": "string",
  "price": "number",
  "stock": "number"
}
```

**Response:** `201 Created`

---

### POST /admin/products/:id/images
Upload one or more images for a product.

**Params:**
- `id` — product ID

**Body:** `multipart/form-data`
- `images` — one or more image files

**Response:** `200 OK`
```json
{
  "images": ["string"]
}
```

**Errors:**
- `404` — product not found
- `400` — no images provided or invalid file type

---

### PUT /admin/products/:id
Update a product.

**Params:**
- `id` — product ID

**Body:** same fields as POST, all optional

**Response:** `200 OK`

---

### DELETE /admin/products/:id
Delete a product.

**Params:**
- `id` — product ID

**Response:** `204 No Content`

---

### GET /admin/orders
List all orders.

**Response:** `200 OK`
```json
{
  "orders": [
    {
      "id": "string",
      "customer": { "id": "string", "name": "string", "email": "string" },
      "items": [{ "productId": "string", "quantity": "number", "price": "number" }],
      "total": "number",
      "status": "pending | fulfilled | cancelled",
      "createdAt": "string"
    }
  ]
}
```

---

### GET /admin/orders/:id
Get order detail.

**Params:**
- `id` — order ID

**Response:** `200 OK` — single order object (same shape as above)

**Errors:**
- `404` — order not found
