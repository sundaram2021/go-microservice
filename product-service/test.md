Here are the routes you can use to test the **product-service**. These routes will cover all the functionalities: listing products, viewing a specific product, adding a product, updating a product, and deleting a product. For the **add**, **update**, and **delete** routes, you'll need a valid JWT token from the **user-service** for authentication.

### 1. **List All Products** (Public Route)
- **Method**: `GET`
- **URL**: `http://localhost:8082/products`

This route does not require authentication, and anyone can access it.

```bash
curl -X GET http://localhost:8082/products
```

**Expected Response**:

```json
{
  "products": [
    {
      "ID": 1,
      "Name": "Laptop",
      "Description": "A high-end laptop",
      "Price": 1200.99,
      "Stock": 10
    },
    {
      "ID": 2,
      "Name": "Phone",
      "Description": "A new smartphone",
      "Price": 499.99,
      "Stock": 100
    }
  ]
}
```

### 2. **Get a Single Product by ID** (Public Route)
- **Method**: `GET`
- **URL**: `http://localhost:8082/products/:id`

Replace `:id` with the actual product ID you want to retrieve.

```bash
curl -X GET http://localhost:8082/products/1
```

**Expected Response**:

```json
{
  "product": {
    "ID": 1,
    "Name": "Laptop",
    "Description": "A high-end laptop",
    "Price": 1200.99,
    "Stock": 10
  }
}
```

### 3. **Add a New Product** (Protected Route)
- **Method**: `POST`
- **URL**: `http://localhost:8082/products`
- **Authorization**: Requires a valid JWT token in the `Authorization` header.

You need to obtain a JWT token from the **user-service**. Here's how you can pass the token:

```bash
curl -X POST http://localhost:8082/products \
-H "Authorization: Bearer <your_jwt_token>" \
-H "Content-Type: application/json" \
-d '{
  "name": "Tablet",
  "description": "A new tablet",
  "price": 799.99,
  "stock": 50
}'
```

**Expected Response**:

```json
{
  "message": "Product created successfully",
  "product": {
    "ID": 3,
    "Name": "Tablet",
    "Description": "A new tablet",
    "Price": 799.99,
    "Stock": 50
  }
}
```

### 4. **Update a Product** (Protected Route)
- **Method**: `PUT`
- **URL**: `http://localhost:8082/products/:id`
- **Authorization**: Requires a valid JWT token in the `Authorization` header.

Replace `:id` with the product ID you want to update. For example, to update the product with ID `1`:

```bash
curl -X PUT http://localhost:8082/products/1 \
-H "Authorization: Bearer <your_jwt_token>" \
-H "Content-Type: application/json" \
-d '{
  "name": "Updated Laptop",
  "description": "An updated high-end laptop",
  "price": 1099.99,
  "stock": 8
}'
```

**Expected Response**:

```json
{
  "message": "Product updated successfully",
  "product": {
    "ID": 1,
    "Name": "Updated Laptop",
    "Description": "An updated high-end laptop",
    "Price": 1099.99,
    "Stock": 8
  }
}
```

### 5. **Delete a Product** (Protected Route)
- **Method**: `DELETE`
- **URL**: `http://localhost:8082/products/:id`
- **Authorization**: Requires a valid JWT token in the `Authorization` header.

Replace `:id` with the product ID you want to delete.

```bash
curl -X DELETE http://localhost:8082/products/1 \
-H "Authorization: Bearer <your_jwt_token>"
```

**Expected Response**:

```json
{
  "message": "Product deleted successfully"
}
```

---

### Steps to Obtain the JWT Token from the User-Service

If you haven't already, you need to get a valid JWT token from the **user-service**.

#### 1. **Register a User (Optional, if not already done)**
- **Method**: `POST`
- **URL**: `http://localhost:8081/register`

```bash
curl -X POST http://localhost:8081/register \
-H "Content-Type: application/json" \
-d '{"username": "admin", "password": "password123"}'
```

#### 2. **Login to Get a JWT Token**
- **Method**: `POST`
- **URL**: `http://localhost:8081/login`

```bash
curl -X POST http://localhost:8081/login \
-H "Content-Type: application/json" \
-d '{"username": "admin", "password": "password123"}'
```

**Expected Response**:

```json
{
  "token": "<your_jwt_token>"
}
```

Copy the token from the response and use it in your **product-service** requests that require authentication.

---

### Summary of Routes:

| Method | Route               | Auth Required | Description                       |
|--------|---------------------|---------------|-----------------------------------|
| `GET`  | `/products`          | No            | List all products                 |
| `GET`  | `/products/:id`      | No            | Get details of a specific product |
| `POST` | `/products`          | Yes           | Add a new product                 |
| `PUT`  | `/products/:id`      | Yes           | Update an existing product        |
| `DELETE`| `/products/:id`     | Yes           | Delete a product                  |

Let me know if you encounter any issues during testing or need further adjustments!