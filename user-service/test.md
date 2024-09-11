Great to hear that it's running now! Let's test your user service to ensure that everything is working as expected. Here are some steps to test the API endpoints you have implemented:

### 1. **Register a New User**
We will first test the registration of a new user. Use the following `curl` command or Postman to make the request.

```bash
curl -X POST http://localhost:8081/register \
     -H "Content-Type: application/json" \
     -d '{"username": "user1", "password": "password123"}'
```

You should expect a response like:

```json
{
  "message": "User registered successfully"
}
```

### 2. **Login and Get JWT Token**
Next, we will test the login functionality, which should return a JWT token.

```bash
curl -X POST http://localhost:8081/login \
     -H "Content-Type: application/json" \
     -d '{"username": "user1", "password": "password123"}'
```

Expected response (this will contain the JWT token):

```json
{
  "token": "your_jwt_token_here"
}
```

Copy the `token` from the response as you'll need it for the next step.

### 3. **Access Protected Route (Profile)**
Now, let's test accessing a protected route that requires JWT authentication. Use the token from the previous step and pass it in the `Authorization` header.

```bash
curl -X GET http://localhost:8081/profile \
     -H "Authorization: Bearer your_jwt_token_here"
```

Expected response:

```json
{
  "message": "Welcome user1"
}
```

### 4. **Handling Error Cases**
Let’s also test for some common error cases to ensure the service behaves as expected:

#### Register with an existing username:
```bash
curl -X POST http://localhost:8081/register \
     -H "Content-Type: application/json" \
     -d '{"username": "user1", "password": "password123"}'
```

Expected response (user already exists):

```json
{
  "error": "user already exists"
}
```

#### Invalid Login:
```bash
curl -X POST http://localhost:8081/login \
     -H "Content-Type: application/json" \
     -d '{"username": "user1", "password": "wrongpassword"}'
```

Expected response:

```json
{
  "error": "Invalid credentials"
}
```

#### Access protected route with invalid/expired token:
```bash
curl -X GET http://localhost:8081/profile \
     -H "Authorization: Bearer invalid_token"
```

Expected response:

```json
{
  "error": "Invalid or expired token"
}
```

### If Everything Works

If all these tests pass, that means your user service is functioning correctly with registration, login, and JWT authentication.

Let me know the results of the testing!