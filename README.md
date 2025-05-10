# Workout Tracker API

A robust RESTful API built with Go for tracking and managing workout sessions. This application allows users to create, read, update, and delete workout records with detailed exercise entries.

## Project Overview

This project is a Go-based backend service that provides a complete workout tracking solution with user authentication. The application follows clean architecture principles with clear separation of concerns between handlers, business logic, and data storage.

### Key Features

- **User Authentication**: Secure registration and token-based authentication
- **Workout Management**: Create, read, update, and delete workout sessions
- **Exercise Tracking**: Add detailed exercise entries to workouts including sets, reps, duration, and weight
- **Authorization**: Users can only modify their own workouts

## Technology Stack

- **Language**: Go 1.24.1
- **Web Framework**: Chi Router
- **Database**: PostgreSQL
- **Authentication**: JWT-based token authentication
- **Containerization**: Docker and Docker Compose for development environment

## Project Structure

```
go-workout-tracker-api/
├── database/               # Database data directories
├── docs/                   # Detailed documentation
├── docker-compose.yml      # Docker configuration for PostgreSQL
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── internal/               # Application internal packages
│   ├── api/                # API handlers
│   │   ├── tokens_handler.go
│   │   ├── user_handler.go
│   │   └── workout_handler.go
│   ├── app/                # Application setup and configuration
│   │   └── app.go
│   ├── middleware/         # HTTP middleware
│   ├── routes/             # Route definitions
│   │   └── routes.go
│   ├── store/              # Data access layer
│   │   ├── database.go
│   │   ├── tokens.go
│   │   ├── user_store.go
│   │   └── workout_store.go
│   ├── tokens/             # Token management
│   └── utils/              # Utility functions
├── main.go                 # Application entry point
└── migrations/             # Database migrations
└── .env                    # Environment variables
```

## API Endpoints

### Authentication

- `POST /users` - Register a new user
- `POST /tokens/authentication` - Create authentication token (login)

### Workouts (Authenticated Routes)

- `GET /workouts/{id}` - Get a specific workout by ID
- `POST /workouts` - Create a new workout
- `PUT /workouts/{id}` - Update an existing workout
- `DELETE /workouts/{id}` - Delete a workout

### Health Check

- `GET /health` - Check API health status

## Data Models

### Workout

```go
type Workout struct {
    ID              int            `json:"id"`
    UserID          int            `json:"user_id"`
    Title           string         `json:"title"`
    Description     string         `json:"description"`
    DurationMinutes int            `json:"duration_minutes"`
    CaloriesBurned  int            `json:"calories_burned"`
    Entries         []WorkoutEntry `json:"entries"`
}
```

### Workout Entry

```go
type WorkoutEntry struct {
    ID              int      `json:"id"`
    ExerciseName    string   `json:"exercise_name"`
    Sets            int      `json:"sets"`
    Reps            *int     `json:"reps"`
    DurationSeconds *int     `json:"duration_seconds"`
    Weight          *float64 `json:"weight"`
    Notes           string   `json:"notes"`
    OrderIndex      int      `json:"order_index"`
}
```

### User

```go
type User struct {
    ID           int       `json:"id"`
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    PasswordHash password  `json:"-"` // Not exposed in JSON
    Bio          string    `json:"bio"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
```

### Token

```go
type Token struct {
    Plaintext string    `json:"token"`
    Hash      []byte    `json:"-"` // Not exposed in JSON
    UserID    int       `json:"-"` // Not exposed in JSON
    Expiry    time.Time `json:"expiry"`
    Scope     string    `json:"-"` // Not exposed in JSON
}
```

## Getting Started

### Prerequisites

- Go 1.24.1 or higher
- Docker and Docker Compose (for local development)

### Setup and Installation

1. Clone the repository:

   ```
   git clone https://github.com/vinayakvispute/go-lang.git
   cd go-lang
   ```

2. Create a `.env` file in the root directory with this env variable and add your own database credentials:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_SSLMODE=your_jwt_secret

   ```

3. Start the PostgreSQL database:

   ```
   docker-compose up
   ```

4. Run the application:

   ```
   go run main.go
   ```

   By default, the server runs on port 8080. You can specify a different port using the `-port` flag:

   ```
   go run main.go -port 3000
   ```

## Development

### Database Migrations

The application uses the `goose` migration tool to manage database schema changes. Migrations are automatically applied when the application starts.

### Testing

To run tests:

```
go test ./...
```

A separate test database is configured in the Docker Compose file to ensure tests don't interfere with development data.

## Security Considerations

- Authentication is implemented using JWT tokens
- Passwords are securely hashed before storage
- API endpoints are protected with middleware to ensure proper authorization
- Users can only access and modify their own workout data

## Future Enhancements

- Add support for workout categories and tags
- Implement workout statistics and reporting
- Add social features like sharing workouts
- Create mobile app integration

## Documentation

Detailed documentation for the API, including endpoints, data models, and implementation details, can be found in the [docs/](./docs/) directory.

### Code Improvement's To-Do List

- [ ] Create detailed documentation in docs/ directory covering API endpoints, data models, and implementation details
- [ ] Implement query timeouts for database operations
- [ ] Add proper database locking mechanisms
- [ ] Improve concurrency handling
- [ ] Add more comprehensive error handling

## License

This project is licensed under the MIT License - see the LICENSE file for details.
