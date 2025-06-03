# Go Beginner's Guide: Rock Paper Scissors Project

This document explains the Go concepts and project structure we've implemented. Perfect for beginners learning Go!

## Table of Contents
1. [Go Project Structure](#go-project-structure)
2. [Go Modules](#go-modules)
3. [Packages and Imports](#packages-and-imports)
4. [Database Concepts](#database-concepts)
5. [HTTP Server and API](#http-server-and-api)
6. [Project Architecture](#project-architecture)
7. [Key Go Concepts Used](#key-go-concepts-used)

## Go Project Structure

### Why This Structure?
```
rockpaperscissors/
├── cmd/                    # Main applications
├── internal/               # Private application code
├── web/                    # Web assets
├── data/                   # Database files
├── docs/                   # Documentation
└── tests/                  # Test files
```

**Key Principles:**
- `cmd/` contains the main entry points (executables)
- `internal/` contains private code that can't be imported by other projects
- Clear separation of concerns (database, API, business logic)

### Directory Breakdown

#### `cmd/server/main.go`
This is the **entry point** of our application. In Go:
- `package main` indicates this is an executable program
- `func main()` is where execution starts
- We import other packages to build our application

#### `internal/` Directory
The `internal` directory is special in Go:
- Code here is **private** to this project
- Other projects cannot import from `internal/`
- Great for organizing internal business logic

**Subdirectories:**
- `api/` - HTTP-related code (handlers, routes, middleware)
- `database/` - Database connection and schema management
- `models/` - Data structures that represent our domain
- `services/` - Business logic (will be implemented next)

## Go Modules

### What is `go.mod`?
```go
module rockpaperscissors

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/mattn/go-sqlite3 v1.14.17
)
```

**Explanation:**
- `module rockpaperscissors` - This is our module name
- `go 1.21` - Minimum Go version required
- `require` - External dependencies our project needs

### Dependencies We're Using
1. **Gin** (`github.com/gin-gonic/gin`) - HTTP web framework
   - Makes building REST APIs easy
   - Handles routing, middleware, JSON responses

2. **SQLite Driver** (`github.com/mattn/go-sqlite3`) - Database driver
   - Allows Go to talk to SQLite databases
   - Uses CGO (C bindings) for performance

## Packages and Imports

### Package Declaration
Every Go file starts with `package packagename`:
```go
package main        // Executable program
package models      // Package for data models
package handlers    // Package for HTTP handlers
```

### Import Statements
```go
import (
    "log"                                    // Standard library
    "net/http"                              // Standard library
    "rockpaperscissors/internal/database"   // Our internal package
    "github.com/gin-gonic/gin"              // External dependency
)
```

**Types of Imports:**
- **Standard library**: Built into Go (`log`, `net/http`, `database/sql`)
- **Internal packages**: Our own code (`rockpaperscissors/internal/...`)
- **External packages**: Third-party libraries (`github.com/...`)

## Database Concepts

### SQLite Setup (`internal/database/sqlite.go`)

#### Database Connection
```go
func InitDB() (*sql.DB, error) {
    // Create data directory if it doesn't exist
    dataDir := "data"
    if err := os.MkdirAll(dataDir, 0755); err != nil {
        return nil, fmt.Errorf("failed to create data directory: %v", err)
    }

    // Open database connection
    db, err := sql.Open("sqlite3", dbPath)
    // ... error handling
}
```

**What's happening:**
1. Create a `data/` directory for our database file
2. Open a connection to SQLite database
3. Return the connection or an error

#### Database Migrations
```go
func RunMigrations(db *sql.DB) error {
    usersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        total_coins INTEGER DEFAULT 0,
        current_streak INTEGER DEFAULT 0,
        // ... more fields
    );`
}
```

**Migrations** create and update database structure:
- `CREATE TABLE IF NOT EXISTS` - Only create if table doesn't exist
- `PRIMARY KEY AUTOINCREMENT` - Auto-incrementing ID field
- `UNIQUE NOT NULL` - Username must be unique and not empty
- `DEFAULT 0` - Default value for new records

### Database Schema Design

#### Users Table
```sql
users (
    id              INTEGER PRIMARY KEY,    -- Unique identifier
    username        TEXT UNIQUE NOT NULL,   -- Player name
    total_coins     INTEGER DEFAULT 0,      -- Total coins earned
    current_streak  INTEGER DEFAULT 0,      -- Current win streak
    games_played    INTEGER DEFAULT 0,      -- Total games played
    games_won       INTEGER DEFAULT 0,      -- Total games won
    created_at      DATETIME,               -- When user was created
    updated_at      DATETIME                -- Last update time
)
```

#### Games Table
```sql
games (
    id               INTEGER PRIMARY KEY,   -- Unique game ID
    user_id          INTEGER,               -- Links to users table
    player_choice    TEXT,                  -- rock/paper/scissors
    computer_choice  TEXT,                  -- computer's choice
    result           TEXT,                  -- win/lose/tie
    coins_earned     INTEGER,               -- Coins from this game
    streak_multiplier INTEGER,              -- Multiplier applied
    played_at        DATETIME,              -- When game was played
    FOREIGN KEY (user_id) REFERENCES users(id)  -- Relationship constraint
)
```

## HTTP Server and API

### Gin Web Framework

#### Server Setup (`cmd/server/main.go`)
```go
// Initialize router
router := gin.Default()

// Add CORS middleware
router.Use(func(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    // ... more headers
})

// Setup routes
routes.SetupRoutes(router, db)

// Start server
router.Run(":8080")
```

**What's happening:**
1. Create a new Gin router
2. Add middleware (functions that run before/after requests)
3. Setup our API routes
4. Start server on port 8080

#### Middleware (`internal/api/middleware/`)
Middleware functions run for every request:

```go
func JSONMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Content-Type", "application/json")
        c.Next()  // Continue to next handler
    }
}
```

**Common middleware:**
- **CORS**: Allow cross-origin requests
- **JSON**: Set response content type
- **Error handling**: Consistent error responses

#### Routes (`internal/api/routes/routes.go`)
```go
func SetupRoutes(router *gin.Engine, db *sql.DB) {
    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // API group
    api := router.Group("/api")
    api.POST("/users", userHandler.CreateUser)
    api.POST("/play", gameHandler.PlayGame)
}
```

**Route patterns:**
- `GET /health` - Health check endpoint
- `POST /api/users` - Create new user
- `POST /api/play` - Play a game
- `GET /api/stats/:username` - Get user stats (`:username` is a parameter)

## Project Architecture

### Layered Architecture
```
┌─────────────────┐
│   HTTP Layer    │  <- Handles web requests (handlers)
├─────────────────┤
│  Business Layer │  <- Game logic, user management (services)
├─────────────────┤
│   Data Layer    │  <- Database operations (models)
└─────────────────┘
```

### Data Flow
1. **HTTP Request** comes in (e.g., POST /api/play)
2. **Router** matches the URL to a handler
3. **Middleware** processes the request (CORS, JSON, etc.)
4. **Handler** extracts data from request
5. **Service** contains business logic (game rules, streak calculation)
6. **Database** stores/retrieves data
7. **Response** sent back as JSON

## Key Go Concepts Used

### 1. Structs (Data Types)
```go
type User struct {
    ID           int       `json:"id" db:"id"`
    Username     string    `json:"username" db:"username"`
    TotalCoins   int       `json:"total_coins" db:"total_coins"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
```

**Struct tags:**
- `json:"id"` - How this field appears in JSON
- `db:"id"` - How this field maps to database column

### 2. Methods on Types
```go
func (c Choice) IsValid() bool {
    return c == Rock || c == Paper || c == Scissors
}

func (c Choice) Beats(other Choice) bool {
    switch c {
    case Rock:
        return other == Scissors
    // ...
    }
}
```

**Method syntax:**
- `(c Choice)` - This method belongs to the `Choice` type
- `c` is the receiver (like `this` in other languages)

### 3. Constants and Enums
```go
type Choice string

const (
    Rock     Choice = "rock"
    Paper    Choice = "paper"
    Scissors Choice = "scissors"
)
```

**Custom types** with constants create type-safe "enums"

### 4. Error Handling
```go
db, err := database.InitDB()
if err != nil {
    log.Fatalf("Failed to initialize database: %v", err)
}
```

**Go error handling:**
- Functions return `(result, error)`
- Always check if `err != nil`
- Handle errors explicitly (no exceptions)

### 5. Interfaces (Implicit)
```go
// http.Handler interface
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// Gin handlers implement this implicitly
func MyHandler(c *gin.Context) {
    // Handler code
}
```

### 6. Function Types
```go
// gin.HandlerFunc is a function type
type HandlerFunc func(*Context)

// We can create variables of function types
var handler gin.HandlerFunc = MyHandler
```

## Next Steps

Now that you understand the foundation, we'll implement:

1. **Services Layer** - Business logic for game rules and user management
2. **Complete Handlers** - Full implementation of API endpoints
3. **Testing** - Unit and integration tests
4. **Advanced Features** - Leaderboards, game history, etc.

## Common Go Commands

```bash
go run cmd/server/main.go          # Run the application
go build cmd/server/main.go        # Build executable
go mod tidy                        # Clean up dependencies
go test ./...                      # Run all tests
go fmt ./...                       # Format all code
```

## Helpful Resources

- [Go Tour](https://tour.golang.org/) - Interactive Go tutorial
- [Go by Example](https://gobyexample.com/) - Practical examples
- [Effective Go](https://golang.org/doc/effective_go.html) - Best practices
- [Gin Documentation](https://gin-gonic.com/docs/) - Web framework docs 