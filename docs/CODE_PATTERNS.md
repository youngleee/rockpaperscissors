# Go Code Patterns Reference

Quick reference for the specific code patterns used in our Rock Paper Scissors project.

## 1. Package Structure Pattern

```go
// Every file starts with package declaration
package main           // Executable
package models         // Data types
package handlers       // HTTP handlers
package services       // Business logic
```

## 2. Import Organization Pattern

```go
import (
    // Standard library first
    "database/sql"
    "fmt"
    "log"
    
    // Internal packages second
    "rockpaperscissors/internal/models"
    "rockpaperscissors/internal/database"
    
    // External packages last
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"  // Underscore for side-effects only
)
```

## 3. Struct Definition Pattern

```go
// User model with JSON and database tags
type User struct {
    ID           int       `json:"id" db:"id"`
    Username     string    `json:"username" db:"username"`
    TotalCoins   int       `json:"total_coins" db:"total_coins"`
    CurrentStreak int      `json:"current_streak" db:"current_streak"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
```

**Tag explanations:**
- `json:"field_name"` - How field appears in JSON responses
- `db:"column_name"` - How field maps to database columns
- `binding:"required"` - Gin validation requirements

## 4. Constants and Custom Types Pattern

```go
// Define custom type for type safety
type Choice string
type GameResult string

// Group related constants
const (
    Rock     Choice = "rock"
    Paper    Choice = "paper"
    Scissors Choice = "scissors"
)

const (
    Win  GameResult = "win"
    Lose GameResult = "lose"
    Tie  GameResult = "tie"
)
```

## 5. Method on Types Pattern

```go
// Method with receiver
func (c Choice) IsValid() bool {
    return c == Rock || c == Paper || c == Scissors
}

// Method that takes parameter
func (c Choice) Beats(other Choice) bool {
    switch c {
    case Rock:
        return other == Scissors
    case Paper:
        return other == Rock
    case Scissors:
        return other == Paper
    default:
        return false
    }
}
```

**Key points:**
- `(c Choice)` is the receiver - the type this method belongs to
- Methods can access the receiver's data
- Methods can be called like `myChoice.IsValid()`

## 6. Error Handling Pattern

```go
// Function that returns value and error
func InitDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }
    
    if err := db.Ping(); err != nil {
        db.Close()  // Clean up on error
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }
    
    return db, nil  // Success: return result and nil error
}

// Calling the function
db, err := database.InitDB()
if err != nil {
    log.Fatalf("Database initialization failed: %v", err)
}
defer db.Close()  // Ensure cleanup happens
```

## 7. HTTP Handler Pattern

```go
// Handler function signature for Gin
func (h *UserHandler) CreateUser(c *gin.Context) {
    // 1. Parse request
    var req models.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request format",
        })
        return
    }
    
    // 2. Validate input
    if req.Username == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Username is required",
        })
        return
    }
    
    // 3. Business logic (would call service)
    // user, err := h.userService.CreateUser(req.Username)
    
    // 4. Return response
    c.JSON(http.StatusCreated, gin.H{
        "message": "User created successfully",
        "user":    user,
    })
}
```

## 8. Database Query Pattern

```go
// Query with parameters (prevents SQL injection)
func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
    query := `
        SELECT id, username, total_coins, current_streak, 
               games_played, games_won, created_at, updated_at
        FROM users 
        WHERE username = ?`
    
    var user models.User
    err := db.QueryRow(query, username).Scan(
        &user.ID,
        &user.Username,
        &user.TotalCoins,
        &user.CurrentStreak,
        &user.GamesPlayed,
        &user.GamesWon,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("database error: %v", err)
    }
    
    return &user, nil
}
```

## 9. Middleware Pattern

```go
// Middleware function returns a handler function
func JSONMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Do something before request
        c.Header("Content-Type", "application/json")
        
        // Continue to next handler
        c.Next()
        
        // Do something after request (optional)
    }
}

// Error handling middleware
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()  // Execute handlers first
        
        // Then check for errors
        if len(c.Errors) > 0 {
            err := c.Errors.Last()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
        }
    }
}
```

## 10. Dependency Injection Pattern

```go
// Handler struct holds dependencies
type UserHandler struct {
    db          *sql.DB
    userService *services.UserService  // Will add this
}

// Constructor function
func NewUserHandler(db *sql.DB) *UserHandler {
    return &UserHandler{
        db: db,
    }
}

// Handlers can access dependencies
func (h *UserHandler) CreateUser(c *gin.Context) {
    // Use h.db for database operations
    // Use h.userService for business logic
}
```

## 11. Route Setup Pattern

```go
func SetupRoutes(router *gin.Engine, db *sql.DB) {
    // Create handlers with dependencies
    gameHandler := handlers.NewGameHandler(db)
    userHandler := handlers.NewUserHandler(db)
    
    // Group related routes
    api := router.Group("/api")
    {
        // Apply middleware to group
        api.Use(middleware.JSONMiddleware())
        api.Use(middleware.ErrorHandler())
        
        // Define routes
        api.POST("/users", userHandler.CreateUser)
        api.GET("/users/:username", userHandler.GetUser)
        api.POST("/play", gameHandler.PlayGame)
    }
}
```

## 12. Configuration Pattern

```go
// Environment variables or config struct
type Config struct {
    Port       string
    DBPath     string
    LogLevel   string
}

func LoadConfig() *Config {
    return &Config{
        Port:     getEnv("PORT", "8080"),
        DBPath:   getEnv("DB_PATH", "./data/game.db"),
        LogLevel: getEnv("LOG_LEVEL", "info"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

## Common Go Idioms Used

### 1. Early Returns
```go
func ValidateUser(user *User) error {
    if user == nil {
        return errors.New("user cannot be nil")
    }
    if user.Username == "" {
        return errors.New("username is required")
    }
    return nil  // Success case at the end
}
```

### 2. Zero Values
```go
var user User                    // All fields get zero values
var count int                    // 0
var name string                  // ""
var active bool                  // false
var createdAt time.Time         // Zero time
```

### 3. Multiple Return Values
```go
// Common pattern: (result, error)
user, err := GetUser(id)
count, err := CountUsers()
success, err := DeleteUser(id)

// Can also return multiple results
min, max, err := GetRange()
```

### 4. Defer for Cleanup
```go
func ProcessFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Will always run when function exits
    
    // Do work with file
    return nil
}
```

These patterns provide a solid foundation for Go development and are commonly used across Go projects! 