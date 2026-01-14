# Event Management System

A RESTful API service built with Go for managing events, user authentication, and event registrations.

## Features

- **Event Management**: Create, read, update, and delete events with ownership-based authorization
- **User Authentication**: Secure user signup and login with JWT tokens
- **Event Registration**: Register and cancel registrations for events
- **Password Security**: Bcrypt password hashing for secure user credentials
- **Database Integration**: PostgreSQL database with GORM ORM and automatic migrations
- **RESTful API**: Clean and intuitive API endpoints with proper error handling

## Authentication

All protected endpoints require a JWT token in the `Authorization` header:
```
Authorization: <your-jwt-token>
```

The token is obtained from the `/login` endpoint and is valid for 2 hours.

## API Endpoints

### Events

#### Get All Events
Retrieve a list of all available events.

```http
GET /events
```

**Response**: List of all events

---

#### Get Event by ID
Retrieve a specific event by its ID.

```http
GET /events/:id
```

**Parameters**:
- `id` (path parameter): The unique identifier of the event

**Response**: Event details

---

#### Create Event
Create a new event. Requires authentication.

```http
POST /events
```

**Headers**:
- `Authorization` (required): JWT token

**Request Body**: Event details (JSON)
- `name` (string, required): Event name
- `description` (string): Event description
- `location` (string): Event location
- `datetime` (string, required): Event date and time (ISO 8601 format)

**Response**: Created event object

---

#### Update Event
Update an existing event. Only the event owner can update their events.

```http
PUT /events/:id
```

**Headers**:
- `Authorization` (required): JWT token

**Parameters**:
- `id` (path parameter): The unique identifier of the event

**Request Body**: Updated event details (JSON)
- `name` (string): Event name
- `description` (string): Event description
- `location` (string): Event location
- `datetime` (string): Event date and time (ISO 8601 format)

**Response**: Updated event object

---

#### Delete Event
Delete an event by its ID. Only the event owner can delete their events.

```http
DELETE /events/:id
```

**Headers**:
- `Authorization` (required): JWT token

**Parameters**:
- `id` (path parameter): The unique identifier of the event to delete

**Response**: Success message

---

### Authentication

#### User Signup
Create a new user account.

```http
POST /signup
```

**Request Body**: User registration details (JSON)
- `username` (string, required): Username for the account (must be unique)
- `email` (string, required): User email address (must be unique and valid email format)
- `password` (string, required): User password

**Response**: User object with id, username, and email

---

#### User Login
Authenticate an existing user.

```http
POST /login
```

**Request Body**: Login credentials (JSON)
- `email` (string, required): User email address
- `password` (string, required): User password

**Response**: Authentication token and user details
- `token` (string): JWT authentication token (valid for 2 hours)
- `user` (object): User information (id, username, email)

---

### Event Registration

#### Register for Event
Register a user for a specific event. Requires authentication.

```http
POST /events/:id/register
```

**Headers**:
- `Authorization` (required): JWT token

**Parameters**:
- `id` (path parameter): The unique identifier of the event

**Response**: Registration confirmation

---

#### Cancel Registration
Cancel a user's registration for an event. Requires authentication.

```http
DELETE /events/:id/delete-registration
```

**Headers**:
- `Authorization` (required): JWT token

**Parameters**:
- `id` (path parameter): The unique identifier of the event

**Response**: Cancellation confirmation

---

## Response Format

All API responses follow a consistent format with `status`, `message`, and `data` fields.

### HTTP Status Codes

- `200 OK`: Successful request
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request payload or parameters
- `401 Unauthorized`: Missing or invalid authentication token
- `403 Forbidden`: User not authorized to perform the action
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

---

## Database Models

### User
- `ID` (uint): Primary key
- `Username` (string): Unique username
- `Email` (string): Unique email address
- `Password` (string): Hashed password (bcrypt)
- `CreatedAt` (time.Time): Account creation timestamp
- `UpdatedAt` (time.Time): Last update timestamp

### Event
- `ID` (uint): Primary key
- `Name` (string): Event name
- `Description` (string): Event description
- `Location` (string): Event location
- `DateTime` (time.Time): Event date and time
- `UserID` (*int64): Foreign key to User (event creator)
- `CreatedAt` (time.Time): Creation timestamp
- `UpdatedAt` (time.Time): Last update timestamp

### Registration
- `ID` (uint): Primary key
- `UserID` (int64): Foreign key to User
- `EventID` (uint): Foreign key to Event
- `CreatedAt` (time.Time): Registration timestamp
- `UpdatedAt` (time.Time): Last update timestamp

---

## Getting Started

### Prerequisites

- Go 1.25.5 or higher
- PostgreSQL database
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/SaroarShahan/event-management.git
cd event-management
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
Create a `.env` file in the root directory with the following variables:
```env
VERSION=1.0.0
SERVICE_NAME=event-management
HOST=localhost
SECRET_KEY=your-secret-key-here
PORT=8080
DB_NAME=event_management
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=5432
```

4. Run the application:
```bash
go run main.go
```

The API server will start on the port specified in your `.env` file (default: `:8080`). The application will automatically:
- Connect to the PostgreSQL database
- Run database migrations to create necessary tables (Users, Events, Registrations)

## Project Structure

```
event-management-by-go/
├── main.go                    # Application entry point
├── go.mod                     # Go module dependencies
├── go.sum                     # Dependency checksums
├── README.md                  # Project documentation
├── api/
│   ├── handlers/              # Business logic handlers
│   │   ├── event.go          # Event CRUD operations
│   │   ├── registration.go   # Registration management
│   │   └── user.go           # User management and authentication
│   ├── middlewares/          # HTTP middlewares
│   │   └── authenticate.go   # JWT authentication middleware
│   ├── responses/            # HTTP response handlers
│   │   ├── event.go         # Event API responses
│   │   └── user.go          # User API responses
│   └── routes/               # Route definitions
│       ├── eventRoutes.go    # Event-related routes
│       ├── loginRoutes.go    # Authentication routes
│       └── routes.go        # Route registration
├── cmd/
│   ├── migration.go          # Database migration logic
│   └── server.go            # Server setup and initialization
├── config/
│   └── config.go            # Configuration management
├── infra/
│   └── database/
│       └── db.go            # Database connection setup
└── internals/
    ├── hash.go              # Password hashing utilities
    └── jwt.go               # JWT token generation and verification
```

## Technologies

- **Go 1.25.5**: Programming language
- **Gin**: Web framework for HTTP routing and middleware
- **GORM**: ORM library for database operations
- **PostgreSQL**: Relational database
- **JWT (golang-jwt/jwt/v5)**: JSON Web Token for authentication
- **Bcrypt (golang.org/x/crypto)**: Password hashing
- **godotenv**: Environment variable management

## License

This project is open source and available under the MIT License.

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

## Author

**Saroar Shahan**
