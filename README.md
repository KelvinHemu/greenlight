# 🎬 Greenlight API

[![Go Version](https://img.shields.io/badge/Go-1.25.1-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![API Version](https://img.shields.io/badge/API-v1.0.0-green.svg)](https://github.com/KelvinHemu/greenlight)
[![Database](https://img.shields.io/badge/PostgreSQL-13+-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)

A production-ready RESTful API service for managing movie information, built with Go. This project demonstrates enterprise-level Go development practices, clean architecture principles, database integration, security best practices, and scalable API design patterns.

## 🌟 Features

### Core API Functionality
- **Full CRUD Operations** - Complete Create, Read, Update, Delete operations for movies
- **RESTful API Design** - Clean, intuitive endpoints following REST principles
- **JSON API** - Full JSON request/response handling with comprehensive error management
- **Custom JSON Marshaling** - Elegant custom types with tailored JSON serialization
- **Partial Updates** - Support for PATCH requests with partial field updates
- **Advanced Filtering** - Search and filter movies by title, genres with full-text search
- **Pagination & Sorting** - Efficient data retrieval with customizable pagination and multi-field sorting

### Database & Persistence
- **PostgreSQL Integration** - Production-grade database with connection pooling
- **Database Migrations** - Version-controlled schema migrations for easy deployment
- **Optimistic Locking** - Version-based concurrency control to prevent race conditions
- **Full-Text Search** - PostgreSQL's native full-text search capabilities
- **Database Indexes** - Optimized queries with proper indexing strategies
- **Context-Aware Queries** - Timeout handling for all database operations

### Security & Performance
- **Rate Limiting** - IP-based rate limiting to prevent abuse
- **Password Hashing** - Bcrypt-based secure password storage
- **User Authentication** - Secure user registration system
- **Panic Recovery** - Graceful panic recovery middleware
- **Connection Pooling** - Efficient database connection management
- **Request Validation** - Comprehensive input validation with detailed error messages

### Production Readiness
- **Structured JSON Logging** - Production-grade structured logging system
- **Graceful Shutdown** - Proper server shutdown handling
- **Email Integration** - SMTP-based email notifications for user events
- **Error Handling** - Centralized error handling with appropriate HTTP status codes
- **Configuration Management** - Environment-based configuration with command-line flags
- **TLS Support** - HTTPS ready with TLS certificates included
- **Healthcheck Endpoint** - Service health monitoring for production deployments

## 🏗️ Architecture

The project follows a clean, modular architecture that separates concerns and promotes maintainability:

```
greenlight/
├── cmd/api/                    # Application entry point and HTTP handlers
│   ├── main.go                # Server setup, configuration, and database connection
│   ├── server.go              # HTTP server with graceful shutdown
│   ├── routes.go              # Route definitions and middleware chain
│   ├── movies.go              # Movie CRUD handlers
│   ├── users.go               # User registration and management
│   ├── middleware.go          # Rate limiting and panic recovery
│   ├── healthcheck.go         # Health monitoring endpoint
│   ├── helpers.go             # JSON encoding/decoding utilities
│   └── errors.go              # Centralized error responses
├── internal/
│   ├── data/                  # Data models and database operations
│   │   ├── models.go          # Database model aggregator
│   │   ├── movies.go          # Movie model with full CRUD + search
│   │   ├── users.go           # User model with authentication
│   │   ├── filters.go         # Pagination, sorting, and filtering logic
│   │   └── runtime.go         # Custom runtime type with JSON marshaling
│   ├── jsonlog/               # Structured JSON logging
│   │   └── jsonlog.go         # Logger implementation
│   ├── mailer/                # Email sending functionality
│   │   ├── mailer.go          # SMTP mail sender
│   │   └── templates/         # Email templates
│   ├── models/                # Domain models
│   └── validator/             # Input validation framework
│       └── validator.go       # Validation utilities and rules
├── migrations/                # Database migration files
│   ├── 000001_create_movies_table.up.sql
│   ├── 000002_add_movies_check_constraints.up.sql
│   ├── 000003_add_movies_indexes.up.sql
│   └── 000004_create_users_table.up.sql
└── tls/                       # TLS certificates for HTTPS
```

### Design Patterns & Best Practices

- **Repository Pattern** - Database operations abstracted through model interfaces
- **Dependency Injection** - Application dependencies injected through struct composition
- **Middleware Chain** - Composable middleware for cross-cutting concerns (rate limiting, panic recovery)
- **Interface-Based Design** - Decoupled components for easier testing and maintenance
- **Custom Types** - Type-safe custom types (e.g., `Runtime`) with specialized behavior
- **Envelope Pattern** - Consistent JSON response wrapping
- **Error Wrapping** - Descriptive error messages with context
- **Validation Layer** - Reusable validation logic separate from handlers
- **Database Migrations** - Version-controlled schema changes
- **Optimistic Concurrency Control** - Version field prevents lost updates
- **Context Propagation** - Request context for timeouts and cancellation
- **Connection Pooling** - Efficient resource management with configurable pool settings

## 🚀 Getting Started

### Prerequisites

- Go 1.25.1 or higher
- PostgreSQL 13 or higher
- Git
- migrate CLI (for database migrations) - [Installation Guide](https://github.com/golang-migrate/migrate)

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

3. **Set up PostgreSQL database**
   ```bash
   # Create database
   createdb greenlight
   
   # Run migrations
   migrate -path=./migrations -database="postgres://username:password@localhost/greenlight?sslmode=disable" up
   ```

4. **Set environment variables**
   ```bash
   export GREENLIGHT_DB_DSN="postgres://username:password@localhost/greenlight?sslmode=disable"
   ```

5. **Run the server**
   ```bash
   go run ./cmd/api
   ```

   The server will start on `http://localhost:4000` by default.

### Configuration Options

The API supports extensive configuration via command-line flags:

```bash
go run ./cmd/api \
  -port=3000 \
  -env=production \
  -db-dsn="postgres://user:pass@localhost/greenlight" \
  -db-max-open-conns=25 \
  -db-max-idle-conns=25 \
  -limiter-rps=2 \
  -limiter-burst=4 \
  -limiter-enabled=true
```

**Available flags:**

**Server Configuration:**
- `-port` - Server port (default: 4000)
- `-env` - Environment mode: development|production (default: development)

**Database Configuration:**
- `-db-dsn` - PostgreSQL DSN connection string
- `-db-max-open-conns` - Maximum open connections (default: 25)
- `-db-max-idle-conns` - Maximum idle connections (default: 25)
- `-db-max-idle-time` - Maximum connection idle time (default: "15m")

**Rate Limiter:**
- `-limiter-rps` - Requests per second per IP (default: 2)
- `-limiter-burst` - Maximum burst size (default: 4)
- `-limiter-enabled` - Enable/disable rate limiting (default: true)

**Email Configuration:**
- `-smtp-host` - SMTP server hostname
- `-smtp-port` - SMTP server port (default: 25)
- `-smtp-username` - SMTP authentication username
- `-smtp-password` - SMTP authentication password
- `-smtp-sender` - Email sender address

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

---

### Movies

#### Create Movie
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

**Response:** `201 Created`
```json
{
    "movie": {
        "id": 1,
        "title": "The Shawshank Redemption",
        "year": 1994,
        "runtime": "142 mins",
        "genres": ["drama", "crime"],
        "version": 1
    }
}
```

#### Get Movie
```http
GET /v1/movies/:id
```

**Response:** `200 OK`
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

#### Update Movie
```http
PATCH /v1/movies/:id
Content-Type: application/json
```

**Request Body** (partial updates supported):
```json
{
    "title": "Casablanca (Updated)",
    "year": 1942
}
```

**Response:** `200 OK`

#### Delete Movie
```http
DELETE /v1/movies/:id
```

**Response:** `200 OK`
```json
{
    "message": "movie deleted"
}
```

#### List Movies
```http
GET /v1/movies?title=godfather&genres=crime,drama&page=1&page_size=20&sort=-year
```

**Query Parameters:**
- `title` - Filter by movie title (supports partial matching)
- `genres` - Filter by genres (comma-separated)
- `page` - Page number (default: 1)
- `page_size` - Items per page (default: 20, max: 100)
- `sort` - Sort field: `id`, `title`, `year`, `runtime` (prefix with `-` for descending)

**Response:** `200 OK`
```json
{
    "movies": [
        {
            "id": 1,
            "title": "The Godfather",
            "year": 1972,
            "runtime": "175 mins",
            "genres": ["crime", "drama"],
            "version": 1
        }
    ],
    "metadata": {
        "current_page": 1,
        "page_size": 20,
        "first_page": 1,
        "last_page": 5,
        "total_records": 92
    }
}
```

---

### Users

#### Register User
```http
POST /v1/users
Content-Type: application/json
```

**Request Body:**
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword123"
}
```

**Validation Rules:**
- Name: Required, max 500 characters
- Email: Required, valid email format, unique
- Password: Required, 8-72 characters

**Response:** `201 Created`
```json
{
    "user": {
        "id": 1,
        "created_at": "2025-11-08T10:30:00Z",
        "name": "John Doe",
        "email": "john@example.com",
        "activated": false
    }
}
```

**Note:** A welcome email is automatically sent upon successful registration.

## 🔧 Technical Highlights

### Database Integration

Production-grade PostgreSQL integration with advanced features:

```go
// Connection pool configuration
db.SetMaxOpenConns(cfg.db.maxOpenConns)
db.SetMaxIdleConns(cfg.db.maxIdleConns)
db.SetConnMaxIdleTime(duration)

// Context-aware queries with timeouts
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
```

**Key Features:**
- Connection pooling for efficient resource management
- Context-based query timeouts
- Prepared statement optimization
- Transaction support
- Full-text search with PostgreSQL's `tsvector` and `tsquery`

### Custom JSON Types

The project implements custom JSON marshaling/unmarshaling for the `Runtime` type, demonstrating advanced Go interface implementation:

```go
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
    jsonValue := fmt.Sprintf("%d mins", r)
    return []byte(strconv.Quote(jsonValue)), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
    // Custom parsing logic for "142 mins" format
}
```

### Validation Framework

A flexible, reusable validation system that separates validation logic from business logic:

```go
v := validator.New()
v.Check(movie.Title != "", "title", "must be provided")
v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
if !v.Valid() {
    // Returns structured error response
}
```

### Middleware Architecture

Composable middleware chain for cross-cutting concerns:

```go
// Panic recovery middleware
func (app *application) recoverPanic(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                w.Header().Set("Connection", "close")
                app.serverErrorResponse(w, r, fmt.Errorf("%v", err))
            }
        }()
        next.ServeHTTP(w, r)
    })
}

// Rate limiting with token bucket algorithm
func (app *application) rateLimit(next http.Handler) http.Handler {
    // IP-based rate limiting with golang.org/x/time/rate
}
```

### Rate Limiting Implementation

IP-based rate limiting using the token bucket algorithm:

- Configurable requests per second (RPS) and burst size
- Per-client rate limiters with automatic cleanup
- Background goroutine removes inactive clients
- Thread-safe with mutex synchronization

### Optimistic Locking

Prevents lost updates in concurrent environments:

```go
// Version field incremented on each update
UPDATE movies 
SET title = $1, version = version + 1
WHERE id = $2 AND version = $3
RETURNING version

// Returns ErrEditConflict if version mismatch detected
```

### Security Features

**Password Security:**
```go
// Bcrypt hashing with cost factor 12
hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

// Secure password comparison
err := bcrypt.CompareHashAndPassword(hash, []byte(plaintext))
```

**Input Validation:**
- SQL injection prevention through prepared statements
- JSON parsing with unknown field detection
- Request size limits (1MB default)
- Type-safe parameter binding

### Error Handling

Centralized error handling with specific error types for different scenarios:

- `badRequestResponse` - 400 Bad Request
- `notFoundResponse` - 404 Not Found
- `methodNotAllowedResponse` - 405 Method Not Allowed
- `editConflictResponse` - 409 Conflict (version mismatch)
- `rateLimitExceededResponse` - 429 Too Many Requests
- `failedValidationResponse` - 422 Unprocessable Entity
- `serverErrorResponse` - 500 Internal Server Error

### JSON Request Handling

Robust JSON decoding with comprehensive error detection:

- Syntax errors with character position
- Type mismatches with field names
- Unknown fields detection (strict parsing)
- Request size limits
- Multiple JSON value detection
- Empty body detection

### Structured Logging

Production-ready JSON logging for easy parsing and analysis:

```go
logger.PrintInfo("database connection pool established", map[string]string{
    "db": cfg.db.dsn,
})

logger.PrintError(err, map[string]string{
    "request_method": r.Method,
    "request_url": r.URL.String(),
})
```

### Pagination & Filtering

Efficient data retrieval with advanced filtering:

```go
// SQL query with full-text search and pagination
SELECT count(*) OVER(), id, title, year, runtime, genres, version
FROM movies
WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
AND (genres @> $2 OR $2 = '{}')
ORDER BY %s %s, id ASC
LIMIT $3 OFFSET $4
```

**Features:**
- Full-text search on movie titles
- Array containment for genre filtering
- Window functions for total count
- Configurable sorting and pagination
- Metadata with page information

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

- **[httprouter](https://github.com/julienschmidt/httprouter)** - High-performance HTTP request router
- **[pq](https://github.com/lib/pq)** - Pure Go PostgreSQL driver for database/sql
- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - Password hashing implementation
- **[rate](https://pkg.go.dev/golang.org/x/time/rate)** - Token bucket rate limiting
- **[mail](https://github.com/go-mail/mail)** - Email sending library with SMTP support

## 🗄️ Database Schema

### Movies Table
```sql
CREATE TABLE movies (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    year integer NOT NULL,
    runtime integer NOT NULL,
    genres text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
);

-- Indexes for performance
CREATE INDEX IF NOT EXISTS movies_title_idx ON movies USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS movies_genres_idx ON movies USING GIN (genres);

-- Check constraints
ALTER TABLE movies ADD CONSTRAINT movies_runtime_check CHECK (runtime >= 1);
ALTER TABLE movies ADD CONSTRAINT movies_year_check CHECK (year BETWEEN 1888 AND date_part('year', now()));
ALTER TABLE movies ADD CONSTRAINT genres_length_check CHECK (array_length(genres, 1) BETWEEN 1 AND 5);
```

### Users Table
```sql
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1
);
```

## 🛣️ Roadmap

Completed features:
- [x] RESTful API with full CRUD operations
- [x] PostgreSQL database integration
- [x] Database migrations
- [x] User registration system
- [x] Password hashing with bcrypt
- [x] Rate limiting
- [x] Pagination and sorting
- [x] Advanced filtering and search
- [x] Optimistic concurrency control
- [x] Email notifications
- [x] Structured JSON logging
- [x] Graceful shutdown
- [x] Panic recovery middleware

Future enhancements planned:
- [ ] JWT-based authentication & authorization
- [ ] User activation via email tokens
- [ ] Refresh token mechanism
- [ ] Role-based access control (RBAC)
- [ ] API metrics and monitoring (Prometheus)
- [ ] Distributed tracing (OpenTelemetry)
- [ ] Comprehensive test coverage (unit, integration, e2e)
- [ ] Docker containerization
- [ ] Kubernetes deployment manifests
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] API documentation (Swagger/OpenAPI)
- [ ] GraphQL endpoint
- [ ] Caching layer (Redis)
- [ ] Background job processing
- [ ] WebSocket support for real-time updates

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**Kelvin Hemu**

- GitHub: [@KelvinHemu](https://github.com/KelvinHemu)

## 💡 Key Takeaways

This project demonstrates:

- ✅ **Enterprise-Level Go Development** - Production-ready code with proper architecture
- ✅ **Database Expertise** - PostgreSQL integration, migrations, indexing, full-text search
- ✅ **Security Best Practices** - Password hashing, input validation, rate limiting, SQL injection prevention
- ✅ **RESTful API Design** - Complete CRUD operations with proper HTTP methods and status codes
- ✅ **Concurrency & Performance** - Context usage, goroutines, connection pooling, optimistic locking
- ✅ **Middleware Architecture** - Composable middleware for panic recovery and rate limiting
- ✅ **Advanced Go Features** - Custom types, interfaces, error handling, JSON marshaling
- ✅ **Production Readiness** - Structured logging, graceful shutdown, email notifications, monitoring endpoints
- ✅ **Code Quality** - Clean architecture, comprehensive validation, proper error handling
- ✅ **Professional Documentation** - Clear, detailed, and well-organized

## 🙏 Acknowledgments

- Inspired by best practices in Go API development
- Built with clean architecture principles
- Follows RESTful API design standards
- Based on modern Go web service patterns

---

⭐ If you find this project useful, please consider giving it a star!
