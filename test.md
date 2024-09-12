Step 4: Test the API
You can now test your services again using localhost and the NodePorts:

User-Service:

bash
Copy code
curl -X POST http://localhost:30001/register \
-H "Content-Type: application/json" \
-d '{"username": "user11", "password": "password123", "role": "user"}'
Product-Service:

bash
Copy code
curl -X POST http://localhost:30002/products \
-H "Authorization: Bearer <jwt_token>" \
-H "Content-Type: application/json" \
-d '{"name": "Product 1", "price": 100.00, "stock": 10}'
Order-Service:

bash
Copy code
curl -X POST http://localhost:30003/orders \
-H "Authorization: Bearer <jwt_token>" \
-H "Content-Type: application/json" \
-d '{"user_id": 1, "product_id": 2, "quantity": 3}'
