# ShopWithGolang

ShopWithGolang is a mini e-commerce platform built for learning Go web development.

## Features

- 📦 Product management with CRUD operations
- 🔐 User authentication and registration
- 🛒 Shopping cart support
- 💳 Order processing and order history

## Tech Stack

- **Language:** Go
- **Routing:** Gorilla Mux
- **Database:** PostgreSQL / MySQL (optional)

## Prerequisites

- Go 1.20 or newer

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/shopwithgolang.git
   cd shopwithgolang
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Copy `.env` file from example and update values:
   ```bash
   cp env.example .env
   ```

## Run

```bash
go run main.go
```

## Notes

- The server uses port `9000` by default.
- Ensure port `9000` is free before running the app.
