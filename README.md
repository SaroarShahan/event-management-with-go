# Event Management System

A RESTful API service built with Go for managing events, user authentication, and event registrations.

## Features

- **Event Management**: Create, read, update, and delete events
- **User Authentication**: User signup and login functionality
- **Event Registration**: Register and cancel registrations for events
- **RESTful API**: Clean and intuitive API endpoints

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
GET /events/<id>
```

**Parameters**:
- `id` (path parameter): The unique identifier of the event

**Response**: Event details

---

#### Create Event
Create a new event.

```http
POST /events
```

**Request Body**: Event details (JSON)

**Response**: Created event object

---

#### Update Event
Update an existing event.

```http
PUT /events
```

**Request Body**: Updated event details (JSON)

**Response**: Updated event object

---

#### Delete Event
Delete an event by its ID.

```http
DELETE /events/<id>
```

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
- `username` (string): Username for the account
- `email` (string): User email address
- `password` (string): User password

**Response**: User object or authentication token

---

#### User Login
Authenticate an existing user.

```http
POST /login
```

**Request Body**: Login credentials (JSON)
- `email` (string): User email address
- `password` (string): User password

**Response**: Authentication token

---

### Event Registration

#### Register for Event
Register a user for a specific event.

```http
POST /events/<id>/register
```

**Parameters**:
- `id` (path parameter): The unique identifier of the event

**Request Body**: Registration details (JSON)

**Response**: Registration confirmation

---

#### Cancel Registration
Cancel a user's registration for an event.

```http
DELETE /events/<id>/register
```

**Parameters**:
- `id` (path parameter): The unique identifier of the event

**Response**: Cancellation confirmation

---

## Getting Started

### Prerequisites

- Go 1.25.5 or higher
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

3. Run the application:
```bash
go run main.go
```

The API server will start on the default port (typically `:8080` or as configured).

## Project Structure

```
event-management-by-go/
├── main.go          # Main application entry point
├── go.mod           # Go module dependencies
└── README.md        # Project documentation
```

## Technologies

- **Go**: Programming language
- **Standard Library**: HTTP server and routing

## License

This project is open source and available under the MIT License.

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

## Author

**Saroar Shahan**
