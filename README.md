# 🪨📄✂️ Rock Paper Scissors

[![Live Demo](https://img.shields.io/badge/🚀_Live_Demo-Available-brightgreen)](https://rockpaperscissorsgo-ei6j.onrender.com)
[![Docker](https://img.shields.io/badge/🐳_Docker-Ready-blue)](https://hub.docker.com)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8)](https://golang.org)

A full-stack Rock Paper Scissors game built with **Go** and modern web technologies. Features a beautiful interactive frontend, comprehensive REST API, streak multiplier system, leaderboard, and containerized deployment.

<div align="center">

## 🎮 **[► PLAY THE GAME NOW ◄](https://rockpaperscissorsgo-ei6j.onrender.com)** 🎮

[![Play Now](https://img.shields.io/badge/🎮_PLAY_NOW-LIVE_DEMO-ff6b6b?style=for-the-badge&logoColor=white&logo=gamepad2)](https://rockpaperscissorsgo-ei6j.onrender.com)

**👆 Click above to play the live game! 👆**

</div>

## ✨ Features

### 🎮 **Game Features**
- **Classic Gameplay**: Rock beats Scissors, Scissors beats Paper, Paper beats Rock
- **Streak System**: Win consecutive games for multiplier bonuses (2x → 3x → 4x)
- **Coin Rewards**: Earn coins for wins with streak multipliers
- **Real-time Stats**: Live tracking of wins, losses, coins, and streaks
- **Leaderboard**: Global rankings updated in real-time

### 🖥️ **Frontend Features**
- **Beautiful UI**: Modern gradient design with smooth animations
- **Responsive Design**: Works perfectly on desktop, tablet, and mobile
- **Interactive Elements**: Hover effects, loading spinners, result animations
- **Real-time Updates**: Live leaderboard and stats without page refresh
- **User-Friendly**: Intuitive interface with emoji-based game choices

### 🔧 **Technical Features**
- **REST API**: Comprehensive backend with full CRUD operations
- **Database**: SQLite with proper migrations and indexing
- **Docker Ready**: Multi-stage Dockerfile for production deployment
- **Test Coverage**: Comprehensive unit and integration tests
- **Error Handling**: Robust error handling with user-friendly messages
- **CORS Support**: Properly configured for frontend-backend communication

## 🛠️ Tech Stack

- **Backend**: Go 1.21+ with Gin web framework
- **Database**: SQLite3 with foreign key constraints
- **Frontend**: Vanilla HTML5, CSS3, JavaScript (ES6+)
- **Containerization**: Docker with multi-stage builds
- **Deployment**: Render, Railway, or any Docker-compatible platform
- **Testing**: Go testing framework with comprehensive coverage

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                          🌐 Frontend                            │
│  Interactive UI with real-time updates and responsive design    │
└─────────────────────┬───────────────────────────────────────────┘
                      │ HTTP/JSON API
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                     🏢 Go REST API                              │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌─────────────┐│
│  │   Routes    │ │  Handlers   │ │  Services   │ │ Middleware  ││
│  │   (Gin)     │ │ (HTTP Layer)│ │(Business)   │ │ (CORS/etc)  ││
│  └─────────────┘ └─────────────┘ └─────────────┘ └─────────────┘│
└─────────────────────┬───────────────────────────────────────────┘
                      │ SQL Queries
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                    🗄️ SQLite Database                           │
│  Users Table: ID, Username, Coins, Streak, Games               │
│  Games Table: UserID, Choices, Result, Coins, Timestamp        │
└─────────────────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
rockpaperscissors/
├── 🐳 Dockerfile                    # Multi-stage production build
├── 🐳 docker-compose.yml           # Local development setup
├── ⚙️ railway.json                 # Railway deployment config
├── ⚙️ render.yaml                  # Render deployment config
├── 📝 .dockerignore                # Docker build optimization
│
├── cmd/server/
│   └── main.go                     # 🚀 Application entry point
│
├── internal/
│   ├── api/
│   │   ├── handlers/               # 📞 HTTP request handlers
│   │   │   ├── game.go            # Game play endpoints
│   │   │   └── user.go            # User management endpoints
│   │   ├── middleware/            # 🛡️ CORS, error handling
│   │   └── routes/                # 🗺️ API route definitions
│   │
│   ├── database/
│   │   └── sqlite.go              # 🗄️ Database connection & migrations
│   │
│   ├── models/                    # 📋 Data structures
│   │   ├── user.go               # User model
│   │   └── game.go               # Game model
│   │
│   └── services/                  # 🧠 Business logic
│       ├── user_service.go       # User operations
│       ├── game_service.go       # Game operations
│       └── game_logic.go         # Game rules & calculations
│
├── web/
│   ├── templates/
│   │   └── index.html            # 🎨 Interactive frontend
│   └── static/
│       └── favicon.ico           # Website icon
│
├── tests/                        # 🧪 Comprehensive test suite
└── docs/
    └── VISUAL_GUIDE.md          # 📚 Complete architecture guide
```

## 🚀 Quick Start

### Local Development

```bash
# Clone the repository
git clone https://github.com/yourusername/rockpaperscissors.git
cd rockpaperscissors

# Install dependencies
go mod tidy

# Run the server
go run cmd/server/main.go

# Open in browser
open http://localhost:8080
```

### With Docker

```bash
# Build and run with Docker Compose
docker-compose up --build

# Or build and run manually
docker build -t rockpaperscissors .
docker run -p 8080:8080 rockpaperscissors
```

## 🎮 How to Play

1. **Enter Username**: Create an account or sign in with existing username
2. **Choose Your Weapon**: Click 🪨 Rock, 📄 Paper, or ✂️ Scissors
3. **Battle Computer**: See the results instantly with battle animations
4. **Build Streaks**: Win consecutive games for coin multipliers
5. **Climb Leaderboard**: Compete with other players for the top spot

### 🏆 Scoring System

| Streak Level | Multiplier | Coins per Win |
|:------------:|:----------:|:-------------:|
| 0 wins       | 1x         | 10 coins      |
| 1+ wins      | 2x         | 20 coins      |
| 2+ wins      | 3x         | 30 coins      |
| 3+ wins      | 4x         | 40 coins      |

## 📡 API Reference

### Game Endpoints
```http
POST /api/play
Content-Type: application/json

{
  "username": "player123",
  "player_choice": "rock"
}
```

### User Management
```http
# Create new user
POST /api/users
Content-Type: application/json

{
  "username": "newplayer"
}

# Get user info
GET /api/users/:username

# Get user statistics
GET /api/stats/:username
```

### Leaderboard
```http
# Get top players
GET /api/leaderboard

# Get user's game history
GET /api/users/:username/games
```

## 🐳 Deployment

### Deploy to Render (Free)

1. **Fork this repository**
2. **Connect to [Render](https://render.com)**
3. **Create Web Service** from your GitHub repo
4. **Environment**: Docker
5. **Deploy!** 🚀

### Deploy to Railway

1. **Connect to [Railway](https://railway.app)**
2. **Deploy from GitHub**
3. **Automatic Docker build** with `railway.json`

### Deploy Anywhere with Docker

```bash
# Build production image
docker build -t rockpaperscissors .

# Run in production
docker run -p 8080:8080 \
  -e GIN_MODE=release \
  rockpaperscissors
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests in verbose mode
go test -v ./...

# Test specific service
go test ./internal/services/...
```

## 🔧 Development

### Prerequisites
- Go 1.21 or higher
- SQLite3
- Docker (optional)

### Environment Variables
```bash
# Optional environment variables
GIN_MODE=release        # Set to 'release' for production
PORT=8080              # Server port (default: 8080)
```

### Database Schema
```sql
-- Users table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    total_coins INTEGER DEFAULT 0,
    current_streak INTEGER DEFAULT 0,
    games_played INTEGER DEFAULT 0,
    games_won INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Games table
CREATE TABLE games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    player_choice TEXT NOT NULL,
    computer_choice TEXT NOT NULL,
    result TEXT NOT NULL,
    coins_earned INTEGER DEFAULT 0,
    streak_multiplier INTEGER DEFAULT 1,
    played_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## 🎯 Features Roadmap

- [ ] **Multiplayer Mode**: Real-time player vs player games
- [ ] **Tournaments**: Bracket-style competitions
- [ ] **Achievement System**: Unlock badges and rewards
- [ ] **Daily Challenges**: Special game modes with bonus rewards
- [ ] **Social Features**: Friend lists and private matches
- [ ] **Analytics Dashboard**: Detailed player statistics
- [ ] **Mobile App**: React Native or Flutter client

## 🤝 Contributing

1. **Fork the Project**
2. **Create Feature Branch** (`git checkout -b feature/amazing-feature`)
3. **Commit Changes** (`git commit -m 'Add amazing feature'`)
4. **Push to Branch** (`git push origin feature/amazing-feature`)
5. **Open Pull Request**

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Gin Framework** - Fast HTTP web framework for Go
- **SQLite** - Lightweight, serverless database
- **Docker** - Containerization platform
- **Render/Railway** - Modern deployment platforms

---

<div align="center">
  <strong>🎮 Happy Gaming! 🎮</strong>
  <br><br>
  Made with ❤️ and Go
  <br>
  <a href="https://rockpaperscissorsgo-ei6j.onrender.com">🚀 Play Live Demo</a>
</div> 