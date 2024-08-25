# Basic Go web server

A simple TCP server implementation that handles basic HTTP GET requests. This server listens on port 8080 and responds with `index.html` file for requests to the root path (`/`). It returns a `404.html` file for all other paths.

## Features

- Listens for TCP connections on `localhost:8080`.
- Handles basic HTTP GET requests (`/`) & other not supported paths.


### Prerequisites

- Go (version 1.22.1 or higher)

### Running the Server
```bash
go run main.go
```

### Test manually using curl

- Homepage
    ```bash
    curl http://localhost:8080/
    ```
- Other paths
  ```bash
  curl http://localhost:8080/example
  ```    
