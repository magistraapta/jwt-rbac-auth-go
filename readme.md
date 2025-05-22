# RBAC with JWT Authentication in Go

This project implements a Role-Based Access Control (RBAC) system with JWT authentication using Go and Gin framework. It demonstrates how to implement secure authentication and authorization in a RESTful API.

## What is RBAC and JWT?

### Role-Based Access Control (RBAC)
RBAC is a method of regulating access to resources based on the roles of individual users. In this system:
- Users are assigned specific roles (e.g., admin, user)
- Each role has specific permissions
- Access to resources is controlled based on these roles

### JSON Web Tokens (JWT)
JWT is a compact, URL-safe means of representing claims between two parties. In this project:
- JWTs are used for authentication
- Tokens contain user information and role
- Tokens are signed to ensure they haven't been tampered with

## Features

- User authentication with JWT
- Role-based access control
- User registration (admin and regular users)
- Product management with role-based permissions
- Secure password hashing
- Middleware for role verification

## API Endpoints

### Authentication
- `POST /login` - User login
- `POST /admin/register` - Register new admin user
- `POST /user/register` - Register new regular user
- `GET /admin/dashboard` - Admin dashboard (admin only)
- `GET /user/dashboard` - User dashboard (user only)

### Products
- `POST /products` - Create new product (admin only)
- `GET /products` - Get all products (public)

## Prerequisites

- Go 1.16 or higher
- PostgreSQL database
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/rbac-go.git
cd rbac-go
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

4. Run the application:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`

## Project Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── auth/
│   │   ├── controllers/
│   │   ├── dto/
│   │   └── model/
│   └── product/
│       ├── controllers/
│       └── model/
├── middleware/
├── utils/
├── db/
├── go.mod
└── go.sum
```

## Security Features

- Password hashing using bcrypt
- JWT token-based authentication
- Role-based middleware for route protection
- Secure cookie handling
- Environment variable configuration

## Testing

Run the tests using:
```bash
go test ./...
```

## License

MIT License
