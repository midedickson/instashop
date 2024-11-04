# E-commerce API

A robust RESTful API built with Go for managing an e-commerce platform. This API handles user authentication, product management, and order processing with proper authentication and authorization mechanisms.

## 🚀 Features

- User authentication using JWT
- Product management (CRUD operations)
- Order management with status tracking
- Role-based access control
- Standardized error handling
- Payload validation
- Docker support
- PostgreSQL database
- API documentation

## 🛠 Tech Stack

- **Language:** Go (standard libraries)
- **Router:** Gorilla Mux
- **Database:** PostgreSQL
- **Authentication:** JWT (golang-jwt)
- **Password Hashing:** bcrypt
- **Testing:** Go testing with Testify
- **Documentation:** Swagger
- **Containerization:** Docker
- **Architecture:** Clean Architecture

## 📁 Project Structure

```
├── config/             # Configuration management
├── constants/          # Application constants
├── database/          # Database connection and setup
├── dto/               # Data Transfer Objects
│   ├── payloads.go    # Request/Response structures
│   └── responses.go   # Response formatting
├── entity/            # Database entities
│   ├── order.go
│   ├── product.go
│   └── user.go
├── internal/          # Internal application code
│   ├── http/         # HTTP handlers and middleware
│   ├── repository/   # Data access layer
│   └── services/     # Business logic layer
├── tests/            # Test suites
│   ├── mocks/        # Mock objects
│   └── unit/         # Unit tests
└── docker-compose.yml # Docker configuration
```

## 🔧 Prerequisites

- Docker and Docker Compose
- Go 1.19 or higher (for development)
- PostgreSQL (handled by Docker)

## 🚀 Getting Started

1. Clone the repository:

```bash
git clone [https://github.com/midedickson/instashop]
cd [instashop]
```

2. Copy the example environment file:

```bash
cp .env.sample .env
```

3. Start the application using Docker:

```bash
docker compose up -d --build
```

The API will be available at `http://localhost:8080/api/v1`
pgAdmin interface will be accessible at `http://localhost:5050`

## Video Walkthrough

Watch a complete demonstration of the API functionality.

The videos covers:

- Project setup and installation
- Database configuration
- API endpoint demonstrations
- Authentication flow [https://www.loom.com/share/142f449a5b754c74818b4366fe51b4f8?sid=c05118e7-c8a3-4d06-8b52-81c42ccf5b6d]
- Error handling examples [https://www.loom.com/share/cee04c7113e7483bb9b2d11e18ece68e?sid=a3220be6-1722-4bc4-910e-9018245d14c3]
- Testing procedures

## 📄 API Documentation

Full API documentation is available here: [https://documenter.getpostman.com/view/26825676/2sAY4xBMhq]

The Postman documentation provides detailed information about:

- All available endpoints
- Request/response schemas
- Authentication requirements
- Example requests
- Status codes and error responses

## 🔒 Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-token>
```

## 🧪 Testing

The project includes unit tests with approximately 50% coverage. Run the tests using:

```bash
go test ./... -v
```

Tests use the Testify framework for assertions and mocking.

## 🔧 Error Handling

The API implements standardized error responses using custom error types and a consistent response writer pattern. All responses follow the format:

```json
{
    "success": boolean,
    "message": "string",
    "data": object|null
}
```

## 🐳 Docker Configuration

The application runs in Docker containers with the following services:

- API server (port 8080)
- PostgreSQL database
- pgAdmin (port 5050)

## 📝 Development Notes

- Uses clean architecture principles for better separation of concerns
- Implements middleware for authentication, permissions, and payload validation
- Custom error handling and response writing patterns
- DTOs for request/response handling
- Repository pattern for data access
