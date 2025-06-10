# Visual Guide: How Everything Works Together

**This guide uses pictures and diagrams to show how our Rock Paper Scissors game works.**

## 🏗️ The Big Picture: Our Application Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                          🌐 THE INTERNET                        │
│                                                                 │
│  👤 User's Browser                                              │
│  ├── Types: http://localhost:8080/api/play                     │
│  ├── Sends: {"username": "john123", "choice": "rock"}          │
│  └── Gets back: {"result": "win", "coins_earned": 20}          │
└─────────────────────────┬───────────────────────────────────────┘
                          │ HTTP Request
                          ▼
┌─────────────────────────────────────────────────────────────────┐
│                    🏢 OUR GO SERVER                            │
│                                                                 │
│  🚪 Router (Traffic Director)                                  │
│  ├── Sees: POST /api/play                                       │
│  ├── Thinks: "This goes to the game handler"                    │ 
│  └── Routes to: GameHandler.PlayGame()                          │
│                          │                                      │
│                          ▼                                      │
│  👔 Middleware (Security Guards)                                │
│  ├── CORS: "Are you allowed to make this request?" ✅           │
│  ├── JSON: "Let's make sure responses are JSON format"          │
│  └── Error: "If anything goes wrong, handle it nicely"          │
│                          │                                      │
│                          ▼                                     │
│  🎯 Handler (Receptionist)                                     │
│  ├── Parses: "john123 wants to play rock"                       │
│  ├── Validates: "Is 'rock' a valid choice?" ✅                 │
│  ├── Calls: GameService.PlayGame(john123, rock)                 │
│  └── Returns: JSON response                                     │
│                          │                                      │
│                          ▼                                      │
│  🧠 Service (The Brain)                                         │
│  ├── Gets user: UserService.GetUser("john123")                  │
│  ├── Computer picks: "scissors"                                 │
│  ├── Determines: "rock beats scissors = win!"                   │
│  ├── Calculates: "streak=2, so 10 coins × 2 = 20 coins"         │
│  ├── Updates user: +20 coins, streak=3                          │
│  └── Records game in history                                    │
│                          │                                      │
│                          ▼                                      │
│  🗄️ Database (Filing Cabinet)                                   │
│  ├── Updates users table: john123 now has 170 coins             │
│  ├── Inserts into games table: new game record                  │
│  └── Returns: success                                           │
└─────────────────────────────────────────────────────────────────┘
```

## 📁 File Organization Visual

```
🏠 rockpaperscissors/
│
├── 🚪 cmd/server/main.go          ← "The key that starts everything"
│   │
│   └── 🔧 What it does:
│       ├── Opens the database
│       ├── Sets up the web server
│       ├── Connects all the pieces
│       └── Starts listening for requests
│
├── 🏠 internal/ (The private rooms of our house)
│   │
│   ├── 📞 api/ (Communication Department)
│   │   ├── 👥 handlers/ (Receptionists)
│   │   │   ├── user.go     ← "Handles user account requests"
│   │   │   └── game.go     ← "Handles game playing requests"
│   │   │
│   │   ├── 🛡️ middleware/ (Security Guards)
│   │   │   └── middleware.go ← "Checks requests before they proceed"
│   │   │
│   │   └── 🗺️ routes/ (Directory)
│   │       └── routes.go   ← "Maps URLs to handlers"
│   │
│   ├── 🗄️ database/ (Filing Cabinet)
│   │   └── sqlite.go       ← "Manages data storage and retrieval"
│   │
│   ├── 📋 models/ (Forms and Templates)
│   │   ├── user.go         ← "Defines what a User looks like"
│   │   └── game.go         ← "Defines what a Game looks like"
│   │
│   └── 🍳 services/ (Kitchen - where the magic happens)
│       └── [Coming next!]  ← "Business logic and game rules"
│
├── 🖼️ web/ (Decoration)
│   └── templates/index.html ← "The pretty face users see"
│
├── 💾 data/ (Storage Room)
│   └── rockpaperscissors.db ← "Where all the data lives"
│
└── 📚 docs/ (Instruction Manuals)
    ├── This file you're reading!
    └── Other helpful guides
```

## 🔄 Data Flow: From Click to Response

### Step 1: User Makes a Request
```
👤 User in browser:
"I want to play rock-paper-scissors with 'rock'"

🖥️ Browser sends:
POST http://localhost:8080/api/play
Content-Type: application/json
{
  "username": "john123",
  "player_choice": "rock"
}
```

### Step 2: Router Receives Request
```
🚪 Router thinks:
"POST /api/play... let me check my routes..."
"Ah! This goes to gameHandler.PlayGame()"

🗺️ Routes to: internal/api/handlers/game.go
```

### Step 3: Middleware Processing
```
🛡️ CORS Middleware:
"Is this request from an allowed origin? ✅ Yes"

🛡️ JSON Middleware:
"Setting response type to JSON..."

🛡️ Error Middleware:
"I'll catch any errors that happen..."
```

### Step 4: Handler Processes Request
```
👔 Game Handler:
1. "Let me parse this JSON..."
   ├── username: "john123" ✅
   └── player_choice: "rock" ✅

2. "Is 'rock' valid?"
   ├── Calls: playerChoice.IsValid()
   └── Returns: true ✅

3. "Now let me call the game service..."
   └── Calls: gameService.PlayGame("john123", "rock")
```

### Step 5: Service (Business Logic)
```
🧠 Game Service:
1. "Get user john123 from database..."
   ├── Calls: userService.GetUser("john123")
   └── Gets: {id: 1, username: "john123", coins: 150, streak: 2}

2. "Computer makes random choice..."
   └── Computer chooses: "scissors"

3. "Who wins? rock vs scissors..."
   ├── Calls: rock.Beats(scissors)
   └── Returns: true (rock beats scissors!)

4. "Calculate coins earned..."
   ├── Base coins: 10
   ├── Current streak: 2
   ├── New streak: 3 (win increases streak)
   └── Coins earned: 10 × 3 = 30 coins

5. "Update user's stats..."
   ├── New total coins: 150 + 30 = 180
   ├── New streak: 3
   ├── Games played: +1
   └── Games won: +1

6. "Save everything to database..."
   ├── Update users table
   └── Insert new game record
```

### Step 6: Database Operations
```
🗄️ Database:
1. "Updating users table..."
   UPDATE users 
   SET total_coins = 180, current_streak = 3, 
       games_played = games_played + 1,
       games_won = games_won + 1
   WHERE username = 'john123'

2. "Recording game in games table..."
   INSERT INTO games 
   (user_id, player_choice, computer_choice, result, 
    coins_earned, streak_multiplier)
   VALUES (1, 'rock', 'scissors', 'win', 30, 3)

3. "All done! ✅"
```

### Step 7: Response Sent Back
```
🧠 Service returns to handler:
{
  player_choice: "rock",
  computer_choice: "scissors",
  result: "win",
  coins_earned: 30,
  new_streak: 3,
  total_coins: 180,
  message: "You won! Rock beats scissors!"
}

👔 Handler sends to browser:
HTTP 200 OK
Content-Type: application/json
{
  "player_choice": "rock",
  "computer_choice": "scissors",
  "result": "win",
  "coins_earned": 30,
  "new_streak": 3,
  "total_coins": 180,
  "message": "You won! Rock beats scissors!"
}

👤 User sees in browser:
"You won! Rock beats scissors! +30 coins"
```

## 🎯 How Different Request Types Work

### Creating a New User
```
👤 User: "I want to create account 'alice99'"

POST /api/users
{"username": "alice99"}
      ↓
👔 UserHandler.CreateUser()
      ↓
🧠 UserService.CreateUser("alice99")
  ├── Check if username exists
  ├── Create new user record
  └── Save to database
      ↓
🗄️ Database: INSERT INTO users...
      ↓
👤 Response: {"message": "User created!", "user": {...}}
```

### Getting User Stats
```
👤 User: "What are alice99's stats?"

GET /api/stats/alice99
      ↓
👔 UserHandler.GetUserStats()
      ↓
🧠 UserService.GetUserStats("alice99")
  ├── Get user from database
  ├── Calculate win rate
  └── Format response
      ↓
🗄️ Database: SELECT * FROM users WHERE username = 'alice99'
      ↓
👤 Response: {"username": "alice99", "coins": 150, "win_rate": 0.75}
```

### Getting Leaderboard
```
👤 User: "Who are the top players?"

GET /api/leaderboard
      ↓
👔 UserHandler.GetLeaderboard()
      ↓
🧠 UserService.GetLeaderboard()
  ├── Get top 10 users by coins
  ├── Calculate win rates
  └── Add ranking numbers
      ↓
🗄️ Database: SELECT * FROM users ORDER BY total_coins DESC LIMIT 10
      ↓
👤 Response: [
  {"rank": 1, "username": "alice99", "coins": 500},
  {"rank": 2, "username": "john123", "coins": 300},
  ...
]
```

## 🗃️ Database Tables Visual

### Users Table
```
┌─────────────────────────────────────────────────────────────────┐
│                          👥 USERS TABLE                         │
├─────┬─────────┬─────────────┬───────────────┬──────────────┬─────────┤
│ id  │username │ total_coins │current_streak │ games_played │games_won│
├─────┼─────────┼─────────────┼───────────────┼──────────────┼─────────┤
│  1  │john123  │     180     │       3       │      25      │   15    │
│  2  │alice99  │     500     │       0       │      40      │   30    │
│  3  │bob456   │     120     │       1       │      15      │    8    │
└─────┴─────────┴─────────────┴───────────────┴──────────────┴─────────┘
```

### Games Table
```
┌──────────────────────────────────────────────────────────────────────────┐
│                              🎮 GAMES TABLE                              │
├────┬────────┬──────────────┬──────────────────┬────────┬─────────────┬─────┤
│ id │user_id │player_choice │ computer_choice  │ result │coins_earned│mult │
├────┼────────┼──────────────┼──────────────────┼────────┼─────────────┼─────┤
│ 1  │   1    │     rock     │    scissors      │  win   │     20      │  2  │
│ 2  │   1    │    paper     │      rock        │  win   │     30      │  3  │
│ 3  │   2    │   scissors   │      rock        │  lose  │      0      │  1  │
│ 4  │   1    │     rock     │     paper        │  lose  │      0      │  1  │
└────┴────────┴──────────────┴──────────────────┴────────┴─────────────┴─────┘
```

### How Tables Connect
```
👤 john123 (user_id = 1) has played these games:
   ├── Game 1: won with rock vs scissors (+20 coins)
   ├── Game 2: won with paper vs rock (+30 coins)  
   └── Game 4: lost with rock vs paper (0 coins)

👤 alice99 (user_id = 2) has played these games:
   └── Game 3: lost with scissors vs rock (0 coins)
```

## 🎯 Game Logic Visual

### Rock Paper Scissors Rules
```
    🗿 ROCK
    ├── ✅ Beats: ✂️ Scissors
    └── ❌ Loses to: 📄 Paper

    📄 PAPER  
    ├── ✅ Beats: 🗿 Rock
    └── ❌ Loses to: ✂️ Scissors

    ✂️ SCISSORS
    ├── ✅ Beats: 📄 Paper
    └── ❌ Loses to: 🗿 Rock
```

### Streak Multiplier System
```
🏆 STREAK MULTIPLIER SYSTEM

Streak 0: 🥉 1x multiplier  (Base: 10 coins)
Streak 1: 🥈 2x multiplier  (Earn: 20 coins) 
Streak 2: 🥇 3x multiplier  (Earn: 30 coins)
Streak 3+: 💎 4x multiplier (Earn: 40 coins - MAX!)

Examples:
├── Win game 1: 0 → 1 streak, earn 10 × 2 = 20 coins
├── Win game 2: 1 → 2 streak, earn 10 × 3 = 30 coins  
├── Win game 3: 2 → 3 streak, earn 10 × 4 = 40 coins
├── Win game 4: 3 → 4 streak, earn 10 × 4 = 40 coins (capped)
└── Lose any game: streak resets to 0
```

### Complete Game Example
```
🎮 COMPLETE GAME FLOW

Initial State:
👤 john123: {coins: 100, streak: 1}

Game Request:
🎯 john123 plays "rock"

Computer Choice:
🤖 Computer randomly picks "scissors"

Battle:
🗿 rock vs ✂️ scissors
Result: 🗿 ROCK WINS! ✅

Calculations:
├── Streak before: 1
├── Streak after: 2 (win increases streak)  
├── Multiplier: 3x (for streak of 2)
├── Base coins: 10
└── Coins earned: 10 × 3 = 30

Updates:
├── john123 coins: 100 + 30 = 130
├── john123 streak: 2
├── games_played: +1
└── games_won: +1

Database Changes:
🗄️ UPDATE users SET total_coins=130, current_streak=2...
🗄️ INSERT INTO games (user_id=1, player_choice='rock'...)

Response:
🎉 "You won! Rock beats scissors! +30 coins"
📊 "New total: 130 coins, streak: 2"
```

## 🔧 Error Handling Visual

### What Happens When Things Go Wrong
```
❌ ERROR SCENARIOS

1. User doesn't exist:
   Request: POST /api/play {"username": "nobody", "choice": "rock"}
   ↓
   🧠 Service: "Looking for user 'nobody'..."
   🗄️ Database: "No user found!"
   ↓
   👔 Handler: Returns 404 Not Found
   Response: {"error": "User not found"}

2. Invalid choice:
   Request: POST /api/play {"username": "john123", "choice": "dynamite"}
   ↓
   👔 Handler: "Is 'dynamite' valid?"
   🎯 Choice.IsValid(): false
   ↓
   Response: {"error": "Invalid choice. Use rock, paper, or scissors"}

3. Database connection fails:
   Request: POST /api/play {"username": "john123", "choice": "rock"}
   ↓
   🗄️ Database: "Connection lost!"
   🧠 Service: "Can't get user data!"
   ↓
   🛡️ Error Middleware: Catches error
   Response: {"error": "Internal server error"}
```

## 🚀 Next Steps Preview

### What We'll Build Next
```
🔮 COMING SOON: Services Layer

🧠 UserService
├── CreateUser(username) 
├── GetUser(username)
├── UpdateUserStats(userID, coins, streak)
└── GetLeaderboard()

🎮 GameService  
├── PlayGame(username, choice)
├── DetermineWinner(playerChoice, computerChoice)
├── CalculateCoins(streak, isWin)
└── GetGameHistory(username)

🎯 Computer Opponent
├── GenerateRandomChoice()
├── Maybe: Add difficulty levels
└── Maybe: Add patterns
```

**You now understand the complete picture! Every piece of our Rock Paper Scissors game and how they all work together. Ready to implement the business logic? 🚀** 