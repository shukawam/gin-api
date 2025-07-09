# shukawam/me/api

This is a simple API server built with Go and the Gin framework.

## Build and Run

### Prerequisites

- Go 1.24 or higher

### Steps

1. **Build the application:**

   ```bash
   go build -o api
   ```

2. **Run the application:**

   ```bash
   ./api
   ```

Alternatively, you can run the application directly without building:

```bash
go run main.go
```

The server will start on `0.0.0.0:8080`.

## Docker

You can also build and run this application using Docker.

1. **Build the Docker image:**

   ```bash
   docker build -t shukawam/me/api .
   ```

2. **Run the Docker container:**

   ```bash
   docker run -p 8080:8080 shukawam/me/api
   ```

## API Endpoints

- `GET /ping`
  - Returns `pong` with a `200 OK` status.
- `GET /sleep`
  - Waits for 120 seconds and then returns `enough sleep?` with a `200 OK` status.
- `GET /error`
  - Causes the application to panic and returns a `500 Internal Server Error`.
