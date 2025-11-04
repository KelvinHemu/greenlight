# 🎬 Greenlight API

[![Go Version](https://img.shields.io/badge/Go-1.25.1-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![API Version](https://img.shields.io/badge/API-v1.0.0-green.svg)](https://github.com/KelvinHemu/greenlight)

A production-ready RESTful API service for managing movie information, built with Go. This project demonstrates modern Go development practices, clean architecture principles, and robust API design patterns.

## 🌟 Features

- **RESTful API Design** - Clean, intuitive endpoints following REST principles
- **JSON API** - Full JSON request/response handling with comprehensive error management
- **Custom JSON Marshaling** - Elegant custom types with tailored JSON serialization
- **Request Validation** - Robust input validation system with detailed error messages
- **Error Handling** - Centralized error handling with appropriate HTTP status codes
- **Structured Logging** - Built-in logging for debugging and monitoring
- **Configuration Management** - Environment-based configuration with command-line flags
- **HTTP Router** - Fast routing using `httprouter` for efficient request handling
- **TLS Support** - HTTPS ready with TLS certificates included
- **Healthcheck Endpoint** - Service health monitoring for production deployments

## 🏗️ Architecture

The project follows a clean, modular architecture that separates concerns and promotes maintainability:

```
greenlight/
├── cmd/api/              # Application entry point and HTTP handlers
│   ├── main.go          # Server setup and configuration
│   ├── routes.go        # Route definitions
│   ├── movies.go        # Movie-related handlers
│   ├── healthcheck.go   # Health monitoring endpoint
│   ├── helpers.go       # Shared helper functions
│   └── errors.go        # Error handling utilities
├── internal/
│   ├── data/            # Data models and business logic
│   │   ├── movies.go    # Movie model and validation
│   │   └── runtime.go   # Custom runtime type with JSON marshaling
│   ├── models/          # Additional data models
│   └── validator/       # Input validation framework
│       └── validator.go # Validation utilities
└── tls/                 # TLS certificates for HTTPS
```

### Design Patterns & Best Practices

- **Dependency Injection** - Application dependencies injected through struct composition
- **Interface-Based Design** - Decoupled components for easier testing and maintenance
- **Custom Types** - Type-safe custom types (e.g., `Runtime`) with specialized behavior
- **Envelope Pattern** - Consistent JSON response wrapping
- **Error Wrapping** - Descriptive error messages with context
- **Validation Layer** - Reusable validation logic separate from handlers

## 🚀 Getting Started

### Prerequisites

- Go 1.25.1 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/KelvinHemu/greenlight.git
   cd greenlight
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the server**
   ```bash
   go run ./cmd/api
   ```

   The server will start on `http://localhost:4000` by default.

### Configuration Options

The API supports configuration via command-line flags:

```bash
go run ./cmd/api -port=3000 -env=production
```

**Available flags:**
- `-port` - Server port (default: 4000)
- `-env` - Environment mode: development|production (default: development)

## 📡 API Endpoints

### Health Check
```http
GET /v1/healthcheck
```
Returns the current status and version of the API.

**Response:**
```json
{
    "status": "available",
    "system_info": {
        "environment": "development",
        "version": "1.0.0"
    }
}
```

### Create Movie
```http
POST /v1/movies
Content-Type: application/json
```

**Request Body:**
```json
{
    "title": "The Shawshank Redemption",
    "year": 1994,
    "runtime": "142 mins",
    "genres": ["drama", "crime"]
}
```

**Validation Rules:**
- Title: Required, max 500 characters
- Year: Required, between 1888 and current year
- Runtime: Required, positive integer
- Genres: Required, 1-5 unique genres

**Response:** `200 OK`

### Get Movie
```http
GET /v1/movies/:id
```

**Response:**
```json
{
    "movie": {
        "id": 1,
        "title": "Casablanca",
        "year": 1942,
        "runtime": "102 mins",
        "genres": ["drama", "romance", "war"],
        "version": 1
    }
}
```

## 🔧 Technical Highlights

### Custom JSON Types

The project implements custom JSON marshaling/unmarshaling for the `Runtime` type, demonstrating advanced Go interface implementation:

```go
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
    jsonValue := fmt.Sprintf("%d mins", r)
    return []byte(strconv.Quote(jsonValue)), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
    // Custom parsing logic
}
```

### Validation Framework

A flexible, reusable validation system that separates validation logic from business logic:

```go
v := validator.New()
v.Check(movie.Title != "", "title", "must be provided")
v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
if !v.Valid() {
    // Handle validation errors
}
```

### Error Handling

Centralized error handling with specific error types for different scenarios:

- `badRequestResponse` - 400 Bad Request
- `notFoundResponse` - 404 Not Found
- `methodNotAllowedResponse` - 405 Method Not Allowed
- `failedValidationResponse` - 422 Unprocessable Entity
- `serverErrorResponse` - 500 Internal Server Error

### JSON Request Handling

Robust JSON decoding with comprehensive error detection:

- Syntax errors with character position
- Type mismatches with field names
- Unknown fields detection
- Request size limits
- Multiple JSON value detection

## 🧪 Testing

Run tests with:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

## 📦 Dependencies

- [httprouter](https://github.com/julienschmidt/httprouter) - High-performance HTTP request router

## 🛣️ Roadmap

Future enhancements planned:
- [ ] Database integration (PostgreSQL)
- [ ] Authentication & Authorization (JWT)
- [ ] Rate limiting
- [ ] Pagination support
- [ ] Full CRUD operations
- [ ] Filtering and sorting
- [ ] API versioning
- [ ] Comprehensive test coverage
- [ ] Docker containerization
- [ ] CI/CD pipeline
- [ ] API documentation (Swagger/OpenAPI)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



## 👤 Author

**Kelvin Hemu**

- GitHub: [@KelvinHemu](https://github.com/KelvinHemu)

##  Acknowledgments

- Inspired by best practices in Go API development
- Built with clean architecture principles
- Follows RESTful API design standards

---

⭐ If you find this project useful, please consider giving it a star!
