# Go API Learning Project

This project is a step-by-step guide to building a basic RESTful API in Go.

## Project Goal

To implement a simple API for managing a collection of "items" with basic CRUD (Create, Read, Update, Delete) operations.

## Project Structure

- `main.go`: Sets up the HTTP server and defines routes using `gorilla/mux`.
- `models/`: Contains data structure definitions (e.g., `item.go` for the `Item` struct).
- `database/`: Handles database connection and initialization (e.g., `db.go`).
- `handlers/`: Contains the logic for each API endpoint (e.g., `get_items.go`, `create_item.go`, etc.).

## Current Status

- Initial Go module created.
- Project refactored into `models`, `database`, and `handlers` packages for better organization.
- `main.go` file set up with an HTTP server using `gorilla/mux` for routing.
- `/` endpoint returns "Welcome to the Go API!".
- `/hello` endpoint returns "Hello World!" for GET requests.
- `Item` struct defined for data modeling, now using a `string` `ID` as the primary key (no `gorm.Model`).
- Integrated SQLite database using GORM.
- Database connection established and `Item` schema migrated automatically.
- UUIDs are now used for item IDs, generated for initial data and new creations.
- `/items` endpoint implemented to handle GET requests (retrieve all items) and POST requests (create a new item).
- `/items/{id}` endpoint implemented to handle GET requests (retrieve a single item), PUT requests (update an item), and DELETE requests (delete an item) from the database.

## How to Run

1. Ensure you are in the project root directory (`go-api-test`).
2. Run the server: `go run main.go`

## Continuous Integration (CI)

This project uses GitHub Actions for continuous integration. The workflow is configured to automatically build and test the Go application on every push to the `main` branch and on every pull request targeting the `main` branch.

The CI workflow performs the following steps:

- Checks out the repository code.
- Sets up the Go environment (version 1.22).
- Caches Go modules to speed up builds.
- Runs `go mod tidy` to ensure module dependencies are clean.
- Builds the Go application.
- Runs all Go tests.

You can view the status of the CI builds in the "Actions" tab of the GitHub repository.

## How to Test Endpoints

- **GET /:** `curl http://localhost:8080/`
- **GET /hello:** `curl http://localhost:8080/hello`
- **GET /items:** `curl http://localhost:8080/items`
- **POST /items:** `curl -X POST -H "Content-Type: application/json" -d '{"name":"New Item","description":"A description for the new item"}' http://localhost:8080/items`
- **GET /items/{id}:** `curl http://localhost:8080/items/YOUR_ITEM_UUID` (replace YOUR_ITEM_UUID with an actual item's UUID)
- **PUT /items/{id}:** `curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Item Name","description":"Updated description"}' http://localhost:8080/items/YOUR_ITEM_UUID` (replace YOUR_ITEM_UUID with an actual item's UUID)
- **DELETE /items/{id}:** `curl -X DELETE http://localhost:8080/items/YOUR_ITEM_UUID` (replace YOUR_ITEM_UUID with an actual item's UUID)

## Application Summary

The application is a basic RESTful API in Go for managing "items".

**Key features implemented:**

- **HTTP Server:** Uses `gorilla/mux` for robust routing.
- **Database:** Integrates with SQLite using the GORM ORM.
- **Data Model:** `Item` struct with `ID` (string, primary key), `Name`, and `Description`.
- **ID Generation:** Uses UUIDs for item IDs, generated for both initial data and new items.
- **Endpoints:**
      - `/`: Returns "Welcome to the Go API!".
      - `/hello`: Returns "Hello World!" for GET requests.
      - `/items`: Handles GET (retrieve all items) and POST (create new item) requests.
      - `/items/{id}`: Handles GET (retrieve single item), PUT (update item), and DELETE (delete item) requests.
