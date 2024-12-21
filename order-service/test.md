
### 1. **Create an Order** (Authenticated Route)

This route allows any authenticated user to create an order by specifying the product ID and quantity. The system will fetch the product details from the product service and calculate the total amount based on the price and quantity.

#### Request:
```bash
curl -X POST http://localhost:8083/orders \
-H "Authorization: Bearer <jwt_token>" \
-H "Content-Type: application/json" \
-d '{
  "user_id": 1,
  "product_id": 2,
  "quantity": 3
}'
```

#### Response:
- **Success**: `200 OK`
    ```json
    {
      "message": "Order created successfully",
      "order": {
        "ID": 1,
        "UserID": 1,
        "ProductID": 2,
        "Quantity": 3,
        "TotalAmount": 150.00,
        "Status": "Pending"
      }
    }
    ```

- **Error**: `401 Unauthorized` (if no or invalid token is provided)
    ```json
    {
      "error": "Invalid or expired token"
    }
    ```

---

### 2. **View Specific Order by ID** (Authenticated Route)

This route allows any authenticated user to view an order by its ID.

#### Request:
```bash
curl -X GET http://localhost:8083/orders/1 \
-H "Authorization: Bearer <jwt_token>"
```

#### Response:
- **Success**: `200 OK`
    ```json
    {
      "order": {
        "ID": 1,
        "UserID": 1,
        "ProductID": 2,
        "Quantity": 3,
        "TotalAmount": 150.00,
        "Status": "Pending"
      }
    }
    ```

- **Error**: `404 Not Found` (if the order does not exist)
    ```json
    {
      "error": "Order not found"
    }
    ```

---

### 3. **View All Orders** (Admin Only)

This route allows an admin to view all orders in the system. Regular users can use this endpoint to view only their own orders by passing a `user_id` query parameter.

#### Admin Request (View All Orders):
```bash
curl -X GET http://localhost:8083/orders \
-H "Authorization: Bearer <admin_jwt_token>"
```

#### User Request (View User's Own Orders):
```bash
curl -X GET http://localhost:8083/orders?user_id=1 \
-H "Authorization: Bearer <jwt_token>"
```

#### Response:
- **Success**: `200 OK`
    ```json
    {
      "orders": [
        {
          "ID": 1,
          "UserID": 1,
          "ProductID": 2,
          "Quantity": 3,
          "TotalAmount": 150.00,
          "Status": "Pending"
        },
        {
          "ID": 2,
          "UserID": 1,
          "ProductID": 3,
          "Quantity": 1,
          "TotalAmount": 100.00,
          "Status": "Shipped"
        }
      ]
    }
    ```

---

### 4. **Update Order Status** (Admin Only)

This route allows an **admin** to update the status of an order (e.g., from `Pending` to `Shipped` or `Delivered`).

#### Request:
```bash
curl -X PUT http://localhost:8083/orders/1/status \
-H "Authorization: Bearer <admin_jwt_token>" \
-H "Content-Type: application/json" \
-d '{
  "status": "Shipped"
}'
```

#### Response:
- **Success**: `200 OK`
    ```json
    {
      "message": "Order status updated successfully",
      "order": {
        "ID": 1,
        "UserID": 1,
        "ProductID": 2,
        "Quantity": 3,
        "TotalAmount": 150.00,
        "Status": "Shipped"
      }
    }
    ```

- **Error**: `403 Forbidden` (if a non-admin user tries to update the status)
    ```json
    {
      "error": "You do not have the necessary permissions to update order status"
    }
    ```

- **Error**: `404 Not Found` (if the order does not exist)
    ```json
    {
      "error": "Order not found"
    }
    ```

---

### Summary of API Endpoints:

| **Route**              | **Method** | **Authentication**  | **Role Required**   | **Description**                                               |
|------------------------|------------|---------------------|---------------------|---------------------------------------------------------------|
| `/orders`              | `POST`     | Yes                 | User/Admin          | Create an order                                                |
| `/orders/:id`          | `GET`      | Yes                 | User/Admin          | View a specific order by ID                                    |
| `/orders`              | `GET`      | Yes                 | Admin (or user can view their own) | View all orders (admin) or user's own orders                   |
| `/orders/:id/status`   | `PUT`      | Yes                 | Admin               | Update the status of an order                                  |

---

