# Album API

This is a simple RESTful API built with the [Gin framework](https://github.com/gin-gonic/gin) in Go. The API manages a list of albums, providing endpoints to retrieve, add, and search for albums.

## Directory Structure

```
/backend
  /cmd
    /api
      main.go                  # Entry point of the application
  /pkg
    /models
      album.go                 # Defines the Album data structure
    /handlers
      album_handler.go        # Contains the HTTP handlers for the application
    /data
      mock_data.go             # Mock data for the application (temporary)
  go.mod                      # Go module file
  go.sum                      # Go checksum file
```

## Setup

1. Ensure you have Go installed (version 1.15 or later recommended).
2. Clone the repository:

```
git clone git@github.com:krishanthisera/gitops-for-devs.git
cd backend
```

3. Install the required dependencies:

```
go mod tidy
```

4. Run the application:

```
go run ./cmd/api
```

The API will start and listen on `localhost:8080`.

## Endpoints

- **GET /albums**: Retrieve a list of all albums.
- **GET /albums/:id**: Retrieve a specific album by its ID.
- **POST /albums**: Add a new album.

Example request to add a new album:

```json
{
    "id": "4",
    "title": "Some Album",
    "artist": "Some Artist",
    "price": 29.99
}
```

## Generating API Documentation

This project utilizes [Gin Swag](https://github.com/swaggo/gin-swagger) to automatically generate API documentation based on annotations in the handler functions.

### Steps to Generate Documentation

1. Ensure you've annotated your handlers appropriately using Swag's comment format.

```go
// GetAlbumByID                godoc
// @Summary      Get single album by id
// @Description  Returns the album whose ISBN value matches the ID.
// @Tags         albums
// @Produce      json
// @Param        id  path      string  true  "search album by ID"
// @Success      200  {object}  models.Album
// @Router       /albums/{id} [get]
func GetAlbumByID(c *gin.Context) {
 ...
}
```

2. In the root directory of the backend application, run the following command to generate the API documentation:

```bash
swag init -g ./cmd/api/main.go -o ./cmd/api/docs/
```

3. After generating the documentation, you can view the API docs by navigating to:

```
http://127.0.0.1:[PORT]/docs/index.html#/
```

Replace `[PORT]` with the port number on which your application is running, typically `8080` for this project.
