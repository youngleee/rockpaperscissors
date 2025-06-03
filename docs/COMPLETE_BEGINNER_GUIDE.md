# Complete Beginner's Guide to Go Programming

**This guide assumes you've never programmed before and explains EVERYTHING step by step.**

## Table of Contents
1. [What is Programming?](#what-is-programming)
2. [What is Go?](#what-is-go)
3. [Understanding Our Project](#understanding-our-project)
4. [Files and Folders Explained](#files-and-folders-explained)
5. [Code Concepts Explained](#code-concepts-explained)
6. [Database Explained](#database-explained)
7. [Web Server Explained](#web-server-explained)
8. [Step-by-Step Code Walkthrough](#step-by-step-code-walkthrough)

## What is Programming?

### Think of Programming Like Cooking
- **Recipe** = Your code (instructions)
- **Ingredients** = Your data (information)
- **Kitchen** = Your computer
- **Chef** = The programming language (Go in our case)
- **Final dish** = Your running application

### What Does Our Program Do?
Our program is like a digital arcade game:
1. Players create accounts (like signing up for a game)
2. They play rock-paper-scissors against the computer
3. When they win, they earn virtual coins
4. Win streaks give bonus coins (like combo bonuses in games)
5. There's a leaderboard showing top players

## What is Go?

### Go is a Programming Language
Just like humans speak English, Spanish, or French, computers understand different programming languages:
- **Go** - What we're using (good for web servers)
- **Python** - Popular for beginners
- **JavaScript** - Used for websites
- **Java** - Used for big business applications

### Why Go?
- **Simple** - Easier to read than many languages
- **Fast** - Programs run quickly
- **Safe** - Helps prevent common mistakes
- **Modern** - Built by Google for today's internet

## Understanding Our Project

### What We're Building: A Web API
Think of our project like a restaurant:

```
Customer (Web Browser) 
    ↓ "I want to play rock-paper-scissors"
Waiter (Our API)
    ↓ Takes order to kitchen
Kitchen (Our Code)
    ↓ Prepares the response
Database (Recipe Book)
    ↓ Stores player scores
Waiter brings back result
    ↓ "You won! +20 coins"
Customer receives response
```

### The Tech Stack (Like Building Blocks)
1. **Go Language** - The foundation (like concrete)
2. **Gin Framework** - Pre-built web tools (like pre-made walls)
3. **SQLite Database** - Storage system (like a filing cabinet)
4. **HTTP Server** - Communication system (like a telephone)

## Files and Folders Explained

### Why Organize Files?
Imagine your bedroom with clothes everywhere vs. organized in drawers. Code organization works the same way!

### Our Folder Structure (Like a House)
```
rockpaperscissors/              🏠 The house (main project)
├── cmd/                        🚪 Front door (main entrance)
│   └── server/
│       └── main.go            🔑 House key (starts everything)
├── internal/                   🏠 Inside the house (private rooms)
│   ├── api/                   📞 Phone system (handles requests)
│   │   ├── handlers/          👥 Receptionists (answer specific requests)
│   │   ├── middleware/        🛡️ Security guards (check requests)
│   │   └── routes/            🗺️ Directory (maps requests to handlers)
│   ├── database/              🗄️ Filing cabinet (data storage)
│   ├── models/                📋 Forms/Templates (data structures)
│   └── services/              🍳 Kitchen (business logic)
├── web/                       🖼️ Decoration (web interface)
├── docs/                      📚 Instruction manuals
├── data/                      💾 Storage room (database files)
└── go.mod                     📦 Shipping label (project info)
```

### File Types Explained

#### `.go` Files (Go Source Code)
- These contain instructions written in Go language
- Like recipes written in English
- Computer reads these to know what to do

#### `.md` Files (Markdown Documentation)
- These are like instruction manuals
- Written in human language
- Explain what the code does

#### `.db` Files (Database Files)
- These store all our data
- Like digital filing cabinets
- Contains user accounts, game scores, etc.

## Code Concepts Explained

### 1. What is a Package?
Think of packages like departments in a store:

```go
package main        // Management office (runs the store)
package models      // Product catalog department
package handlers    // Customer service department
```

**Every Go file starts with a package declaration** - it's like putting a department label on a file.

### 2. What are Imports?
Imports are like calling other departments for help:

```go
import (
    "fmt"                    // Calling the printing department
    "net/http"              // Calling the internet department
    "github.com/gin-gonic/gin"  // Calling an external company for web tools
)
```

### 3. What are Functions?
Functions are like specific jobs or tasks:

```go
func main() {
    // This is the "start the store" function
    // Everything inside these curly braces {} happens when we start
}

func CreateUser() {
    // This is the "create new customer account" function
    // It only runs when someone wants to create an account
}
```

**Key points:**
- `func` means "function" (a job to do)
- `()` holds inputs (like ingredients for a recipe)
- `{}` holds the instructions (the actual recipe steps)

### 4. What are Variables?
Variables are like labeled boxes that hold information:

```go
var username string = "john123"      // A box labeled "username" containing "john123"
var coins int = 100                  // A box labeled "coins" containing the number 100
var isWinner bool = true             // A box labeled "isWinner" containing true/false
```

**Variable types:**
- `string` - Text (like names, messages)
- `int` - Whole numbers (like scores, ages)
- `bool` - True or false (like yes/no questions)

### 5. What are Structs?
Structs are like forms with multiple fields:

```go
type User struct {
    ID           int       // User's unique number (like a customer ID)
    Username     string    // User's chosen name
    TotalCoins   int       // How many coins they have
    GamesPlayed  int       // How many games they've played
}
```

Think of it like a membership card:
```
┌─────────────────────────┐
│   ARCADE MEMBERSHIP     │
│                         │
│ ID: 12345              │
│ Username: john123       │
│ Total Coins: 150        │
│ Games Played: 25        │
└─────────────────────────┘
```

### 6. What are Methods?
Methods are like special abilities that belong to specific types:

```go
type Choice string

// This gives the Choice type a special ability called "IsValid"
func (c Choice) IsValid() bool {
    return c == "rock" || c == "paper" || c == "scissors"
}
```

It's like teaching a type how to answer questions about itself:
- Question: "Are you a valid choice?"
- Answer: "Let me check... yes, I'm rock/paper/scissors!"

### 7. What is Error Handling?
Error handling is like having backup plans:

```go
user, err := GetUser("john123")
if err != nil {
    // Something went wrong! Handle the problem
    fmt.Println("Couldn't find user!")
    return
}
// If we get here, everything worked fine
fmt.Println("Found user:", user.Username)
```

**Why Go does this:**
- Many things can go wrong (user doesn't exist, database is down, etc.)
- Go forces you to think about what to do when things fail
- This makes programs more reliable

## Database Explained

### What is a Database?
A database is like a super-organized filing system:

**Traditional Filing Cabinet:**
```
Drawer 1: Customer Files
├── File A-C
├── File D-F
└── File G-I

Drawer 2: Game Records
├── January Games
├── February Games
└── March Games
```

**Our Database:**
```
Table: users
├── john123: {coins: 150, streak: 3}
├── alice99: {coins: 200, streak: 0}
└── bob456: {coins: 75, streak: 1}

Table: games
├── Game 1: {john123 vs computer, john123 won, +10 coins}
├── Game 2: {alice99 vs computer, alice99 lost, +0 coins}
└── Game 3: {bob456 vs computer, tie, +0 coins}
```

### Our Database Tables

#### Users Table (Like a Membership Directory)
```sql
CREATE TABLE users (
    id              INTEGER PRIMARY KEY,    -- Like a membership number
    username        TEXT UNIQUE NOT NULL,   -- Member's chosen name
    total_coins     INTEGER DEFAULT 0,      -- Their coin balance
    current_streak  INTEGER DEFAULT 0,      -- Current winning streak
    games_played    INTEGER DEFAULT 0,      -- Total games played
    games_won       INTEGER DEFAULT 0,      -- Total games won
    created_at      DATETIME,               -- When they joined
    updated_at      DATETIME                -- Last time info was updated
);
```

**Real example:**
| id | username | total_coins | current_streak | games_played | games_won |
|----|----------|-------------|----------------|--------------|-----------|
| 1  | john123  | 150         | 3              | 15           | 8         |
| 2  | alice99  | 200         | 0              | 20           | 12        |
| 3  | bob456   | 75          | 1              | 10           | 5         |

#### Games Table (Like a Game History Log)
```sql
CREATE TABLE games (
    id               INTEGER PRIMARY KEY,   -- Unique game number
    user_id          INTEGER,               -- Which player (links to users table)
    player_choice    TEXT,                  -- What player chose
    computer_choice  TEXT,                  -- What computer chose
    result           TEXT,                  -- Who won
    coins_earned     INTEGER,               -- Coins from this game
    streak_multiplier INTEGER,              -- Multiplier used
    played_at        DATETIME               -- When game happened
);
```

**Real example:**
| id | user_id | player_choice | computer_choice | result | coins_earned | streak_multiplier |
|----|---------|---------------|-----------------|--------|--------------|-------------------|
| 1  | 1       | rock          | scissors        | win    | 20           | 2                 |
| 2  | 1       | paper         | rock            | win    | 30           | 3                 |
| 3  | 2       | scissors      | rock            | lose   | 0            | 1                 |

### Why Separate Tables?
**Bad way (everything in one table):**
```
| game_id | username | user_coins | player_choice | computer_choice | result |
|---------|----------|------------|---------------|-----------------|--------|
| 1       | john123  | 150        | rock          | scissors        | win    |
| 2       | john123  | 150        | paper         | rock            | win    |
| 3       | john123  | 150        | scissors      | paper           | lose   |
```
Problems: Username and coins repeated, hard to update user info

**Good way (separate tables):**
```
Users: | id | username | coins |
       | 1  | john123  | 150   |

Games: | id | user_id | choice | result |
       | 1  | 1       | rock   | win    |
       | 2  | 1       | paper  | win    |
       | 3  | 1       | scissors| lose  |
```
Benefits: No repetition, easy to update user info, more organized

## Web Server Explained

### What is a Web Server?
A web server is like a restaurant:

1. **Customer** (web browser) walks in with a request
2. **Host** (router) determines which table (endpoint) they need
3. **Waiter** (handler) takes their order and serves the response
4. **Kitchen** (business logic) prepares what they asked for
5. **Storage** (database) provides ingredients

### HTTP Requests (How Browsers Talk to Servers)

#### GET Request (Asking for Information)
```
Browser: "GET /api/users/john123"
Translation: "Please give me information about user john123"

Server: "200 OK" + user data
Translation: "Here's john123's information: coins=150, streak=3"
```

#### POST Request (Sending Information)
```
Browser: "POST /api/play" + {username: "john123", choice: "rock"}
Translation: "john123 wants to play rock-paper-scissors with rock"

Server: "200 OK" + game result
Translation: "Computer played scissors, john123 wins! +20 coins"
```

### Our API Endpoints (Like Menu Items)
| HTTP Method | URL | What it does | Like ordering... |
|-------------|-----|--------------|------------------|
| GET | /health | Check if server is working | "Are you open?" |
| POST | /api/users | Create new user account | "I'd like to sign up" |
| GET | /api/users/john123 | Get user info | "What's my account status?" |
| POST | /api/play | Play a game | "I'd like to play a game" |
| GET | /api/leaderboard | Get top players | "Who are the best players?" |

### Middleware (Like Security Checks)
Middleware runs before your request gets to the handler:

```go
func SecurityMiddleware() {
    // Like a bouncer checking IDs at a club
    // Runs for EVERY request before anything else
}

func JSONMiddleware() {
    // Like a translator ensuring everyone speaks the same language
    // Makes sure all responses are in JSON format
}
```

**Request flow:**
1. Request comes in
2. Security middleware checks it
3. JSON middleware processes it
4. Finally reaches the actual handler

## Step-by-Step Code Walkthrough

### 1. Starting the Application (`cmd/server/main.go`)

```go
package main  // This file is the "manager" - it starts everything
```

**What this means:** This file is in charge of the whole program, like a store manager who opens the store each day.

```go
import (
    "log"                                    // For printing messages
    "net/http"                              // For web server stuff
    "rockpaperscissors/internal/api/routes" // Our custom route definitions
    "rockpaperscissors/internal/database"   // Our custom database code
    "github.com/gin-gonic/gin"              // External web framework
)
```

**What this means:** We're calling in help from different departments:
- `log` - The announcement system (prints messages)
- `net/http` - The internet communication department
- Our custom packages - Our own employees
- `gin` - An external consultant company

```go
func main() {
    // Initialize database
    db, err := database.InitDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()
```

**What this means:**
1. "Hey database department, set up the filing system"
2. If something goes wrong, shut down everything and explain why
3. `defer db.Close()` means "when this program ends, properly close the database connection"

**Why `defer`?** It's like telling someone "when you leave the office, please turn off the lights" - it ensures cleanup happens even if something goes wrong.

```go
    // Set Gin mode
    gin.SetMode(gin.ReleaseMode)
```

**What this means:** Tell Gin to run in "production mode" (quiet, efficient) instead of "development mode" (chatty, shows debug info).

```go
    // Initialize router
    router := gin.Default()
```

**What this means:** Create the main traffic director that will decide which handler gets which request.

```go
    // Add CORS middleware for development
    router.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        
        c.Next()
    })
```

**What this means:** Add a security guard that:
1. Tells browsers "yes, you can make requests from any website"
2. Lists which types of requests are allowed
3. Handles special "preflight" requests browsers send
4. Then lets the request continue to the actual handler

**Why CORS?** Browsers have security rules that block websites from talking to servers on different domains. This tells the browser "it's okay, we allow this."

```go
    // Setup routes
    routes.SetupRoutes(router, db)
```

**What this means:** "Hey routes department, please set up all our endpoints and connect them to the database."

```go
    // Start server
    log.Println("Server starting on :8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
```

**What this means:**
1. Print a message saying we're starting
2. Start listening for requests on port 8080
3. If the server can't start, shut down and explain why

### 2. Database Setup (`internal/database/sqlite.go`)

```go
func InitDB() (*sql.DB, error) {
    // Ensure data directory exists
    dataDir := "data"
    if err := os.MkdirAll(dataDir, 0755); err != nil {
        return nil, fmt.Errorf("failed to create data directory: %v", err)
    }
```

**What this means:**
1. We want to put our database file in a folder called "data"
2. `os.MkdirAll` means "create this folder if it doesn't exist"
3. `0755` are permission settings (who can read/write the folder)
4. If folder creation fails, return an error explaining what went wrong

```go
    // Database file path
    dbPath := filepath.Join(dataDir, "rockpaperscissors.db")
```

**What this means:** Create the full path to our database file. `filepath.Join` is smart about combining folder names regardless of operating system (Windows uses `\`, Mac/Linux use `/`).

```go
    // Open database connection
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }
```

**What this means:**
1. "Please open a connection to a SQLite database at this file path"
2. If it fails, return an error
3. `sql.Open` doesn't actually connect yet - it just prepares to connect

```go
    // Test connection
    if err := db.Ping(); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }
```

**What this means:**
1. `db.Ping()` actually tries to connect to the database
2. If it fails, properly close the connection and return an error
3. This is like testing if a phone line works before making important calls

```go
    // Enable foreign key constraints
    if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
        db.Close()
        return nil, fmt.Errorf("failed to enable foreign keys: %v", err)
    }
```

**What this means:**
1. Send a command to SQLite saying "please enforce foreign key rules"
2. Foreign keys ensure data consistency (you can't reference a user that doesn't exist)
3. If this setting fails, close everything and return an error

### 3. Understanding Struct Tags

```go
type User struct {
    ID           int       `json:"id" db:"id"`
    Username     string    `json:"username" db:"username"`
    TotalCoins   int       `json:"total_coins" db:"total_coins"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
```

**What those weird backtick things mean:**

The `json:"id"` parts are called "tags" - they're like labels that tell other code how to handle this field.

**JSON tags (`json:"id"`):**
- When we send this struct as a web response, use this name
- Go field name: `ID` → JSON field name: `id`
- Why different? JSON convention uses lowercase, Go uses uppercase

**Database tags (`db:"id"`):**
- When we save/load from database, use this column name
- Go field name: `ID` → Database column: `id`

**Example:**
```go
user := User{
    ID: 123,
    Username: "john123",
    TotalCoins: 150,
}

// When sent as JSON response:
{
    "id": 123,
    "username": "john123", 
    "total_coins": 150
}

// When saved to database:
INSERT INTO users (id, username, total_coins) VALUES (123, "john123", 150)
```

### 4. Understanding Methods on Types

```go
func (c Choice) IsValid() bool {
    return c == Rock || c == Paper || c == Scissors
}
```

**Breaking this down:**
- `func` - This is a function
- `(c Choice)` - This function belongs to the `Choice` type
- `c` - This is the name we use inside the function to refer to the choice
- `IsValid()` - The function name
- `bool` - This function returns true or false

**How to use it:**
```go
playerChoice := Choice("rock")
if playerChoice.IsValid() {
    fmt.Println("Valid choice!")
} else {
    fmt.Println("Invalid choice!")
}
```

**Why this is useful:**
Instead of writing validation code everywhere, we attach the ability to self-validate to the Choice type itself.

### 5. Understanding Error Handling Pattern

```go
user, err := GetUser(username)
if err != nil {
    // Handle the error
    return err
}
// Use the user
fmt.Println(user.Username)
```

**Why Go does this:**
- Forces you to think about what can go wrong
- Makes errors visible in the code
- Prevents programs from crashing unexpectedly

**Common error patterns:**
```go
// Pattern 1: Return early on error
if err != nil {
    return err  // Stop here, let caller handle it
}

// Pattern 2: Log and return error
if err != nil {
    log.Printf("Database error: %v", err)
    return fmt.Errorf("failed to get user: %v", err)
}

// Pattern 3: Provide default value
user, err := GetUser(username)
if err != nil {
    user = &User{Username: "guest"}  // Use default
}
```

## Why Our Code is Organized This Way

### Separation of Concerns
Each part has one job:

```
Handlers (internal/api/handlers/)
├── Parse HTTP requests
├── Validate input
├── Call business logic
└── Format HTTP responses

Services (internal/services/) - Coming next!
├── Game rules
├── User management
├── Coin calculations
└── Business logic

Database (internal/database/)
├── Connect to SQLite
├── Run migrations
├── Execute queries
└── Handle database errors

Models (internal/models/)
├── Define data structures
├── Validation methods
├── Type safety
└── JSON/Database mapping
```

### Why This Helps
1. **Easy to find things** - Need to change game rules? Look in services
2. **Easy to test** - Can test each part separately
3. **Easy to change** - Want to switch from SQLite to PostgreSQL? Only change database layer
4. **Easy to understand** - Each file has a clear purpose

## Next Steps for Learning

### What We'll Implement Next
1. **Services Layer** - The brain of our application
2. **Complete Handlers** - Connect everything together
3. **Testing** - Make sure everything works
4. **Frontend** - Pretty interface for users

### Key Concepts You've Learned
- **Project structure** - How to organize Go code
- **Packages and imports** - How Go code reuses other code
- **Structs and methods** - How to define and use custom types
- **Error handling** - How Go deals with things going wrong
- **Database basics** - How to store and retrieve data
- **Web servers** - How browsers talk to backend code

### What Makes You Ready for Next Steps
You now understand:
- ✅ Why files are organized the way they are
- ✅ What each piece of code does
- ✅ How data flows through the application
- ✅ How errors are handled
- ✅ How the database stores information
- ✅ How web requests work

**You're ready to see how we implement the actual game logic!** 