# Rock Paper Scissors Game

A mini rock-paper-scissors game built in Go as a portfolio project for adjoe (mobile ads company). The game simulates a rewarded ad experience where players earn coins for winning games with a streak multiplier system.

## Features

- 🎮 Classic rock-paper-scissors gameplay
- 🏆 Streak multiplier system (2x, 3x, 4x coins for consecutive wins)
- 💰 Coin reward system
- 📊 Player statistics and leaderboard
- 🔄 REST API backend
- 💾 SQLite database for persistence

## Tech Stack

- **Backend**: Go 1.21+ with Gin framework
- **Database**: SQLite3
- **Frontend**: HTML/CSS/JavaScript (or terminal interface)
- **Testing**: Go testing with testify

## Project Structure

```
rockpaperscissors/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/        # HTTP handlers
│   │   ├── middleware/      # API middleware
│   │   └── routes/          # Route definitions
│   ├── database/
│   │   ├── migrations/      # Database migrations
│   │   └── sqlite.go        # Database connection and setup
│   ├── models/              # Data models
│   ├── services/            # Business logic
│   └── utils/               # Utility functions
├── web/
│   ├── static/              # CSS, JS files
│   └── templates/           # HTML templates
├── tests/                   # Integration tests
├── go.mod
├── go.sum
└── README.md
```

## API Endpoints

- `POST /api/play` - Play a game round
- `GET /api/stats/:username` - Get player statistics
- `GET /api/leaderboard` - Get top players
- `POST /api/users` - Create new user

## Getting Started

1. Clone the repository
2. Install Go 1.21 or higher
3. Run `go mod tidy` to install dependencies
4. Run `go run cmd/server/main.go` to start the server
5. Open `http://localhost:8080` in your browser

## Game Rules

- Rock beats Scissors
- Scissors beats Paper
- Paper beats Rock
- Win streaks multiply coin rewards: 1x → 2x → 3x → 4x (max)
- Base coin reward: 10 coins per win

## Development

Run tests:
```bash
go test ./...
```

Run with hot reload (requires air):
```bash
air
```

## License

This project is for portfolio demonstration purposes. 