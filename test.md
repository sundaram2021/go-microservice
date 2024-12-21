
User-Service:

```
curl -X POST http://localhost:30001/register \
-H "Content-Type: application/json" \
-d '{"username": "user11", "password": "password123", "role": "user"}'

```
Product-Service:

```
curl -X POST http://localhost:30002/products \
-H "Authorization: Bearer <jwt_token>" \
-H "Content-Type: application/json" \
-d '{"name": "Product 1", "price": 100.00, "stock": 10}'

```
Order-Service:

```
curl -X POST http://localhost:30003/orders \
-H "Authorization: Bearer <jwt_token>" \
-H "Content-Type: application/json" \
-d '{"user_id": 1, "product_id": 2, "quantity": 3}'
```
