Here are all the API routes for your **user-service** with role-based authentication to test:

### 1. **Register a New User**

This route allows you to create a new user (either with the role of `user` or `admin`).

#### Request:
```bash
curl -X POST http://localhost:8081/register \
-H "Content-Type: application/json" \
-d '{"username": "user1", "password": "password123", "role": "user"}'
```

#### Response:
- **Success**: `200 OK`
    ```json
    {
      "message": "User registered successfully",
      "user": {
        "username": "user1",
        "password": "$2a$10$hashed_password",
        "role": "user"
      }
    }
    ```

- **Error**: `400 Bad Request` (if username already exists)
    ```json
    {
      "error": "Username already exists"
    }
    ```

---

### 2. **Login User**

This route allows a user (with either `user` or `admin` role) to log in and receive a JWT token.

#### Request:
```bash
curl -X POST http://localhost:8081/login \
-H "Content-Type: application/json" \
-d '{"username": "user1", "password": "password123"}'
```

#### Response:
- **Success**: `200 OK`
    ```json
    {
      "token": "<jwt_token>"
    }
    ```

- **Error**: `401 Unauthorized` (if credentials are invalid)
    ```json
    {
      "error": "Invalid credentials"
    }
    ```

---

### 3. **View User Profile**

This route allows an authenticated user (either `admin` or `user`) to view their own profile.

#### Request:
```bash
curl -X GET http://localhost:8081/profile \
-H "Authorization: Bearer <jwt_token>"
```

#### Response:
- **Success**: `200 OK`
    ```json
    {
      "user": {
        "ID": 1,
        "username": "user1",
        "password": "$2a$10$hashed_password",
        "role": "user",
        "CreatedAt": "2024-09-12T00:00:00Z",
        "UpdatedAt": "2024-09-12T00:00:00Z"
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

### 4. **Protected Admin-Only Route (Example)**

If you have any admin-only routes (e.g., accessing user management), you can restrict them based on the `admin` role.

#### Example Request:
```bash
curl -X GET http://localhost:8081/admin \
-H "Authorization: Bearer <jwt_token>"
```

#### Response:
- **Success** (Admin Access): `200 OK`
    ```json
    {
      "message": "Admin-only access"
    }
    ```

- **Error** (Non-admin Access): `403 Forbidden`
    ```json
    {
      "error": "You do not have the necessary permissions to access this route"
    }
    ```

---

### Summary of API Endpoints:

| **Route**         | **Method** | **Authentication** | **Description**                                        |
|-------------------|------------|--------------------|--------------------------------------------------------|
| `/register`       | `POST`     | No                 | Register a new user (with either `admin` or `user` role) |
| `/login`          | `POST`     | No                 | Login to get a JWT token                                |
| `/profile`        | `GET`      | Yes (user/admin)   | View the profile of the currently logged-in user        |
| `/admin`          | `GET`      | Yes (admin)        | Example admin-only route (replace with actual route)    |

---

You can test these routes using the provided `curl` commands. Let me know if you need additional routes or features!