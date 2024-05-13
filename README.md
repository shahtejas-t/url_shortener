# url-shortener

A Go and Gin-based URL shortener application designed to provide users with a convenient way to shorten long URLs into shorter, more manageable links. The application follows the principles of **hexagonal architecture** pattern to ensure modularity, flexibility, and maintainability.

## Functional Requirements

- **Shortening URLs**: Users should be able to input a URL and receive a shortened version.
- **URL Redirection**: When accessing a shortened URL, users should be redirected to the original URL.
- **API Access**: Offer API endpoints for creating, retrieving, and managing shortened URLs.

## Features

- **URL Generation**: Create shortened URLs efficiently.
- **Redirection**: Redirect users to original URLs via shortened links.
- **Deletion**: Safely remove shortened URLs and their associated data.

## Prerequisites

- Go (Golang) installed on your system.
- Docker and Docker-compose

## Technologies Used

- **Go (Golang)**
- **Gin Framework**
- **Redis**: An in-memory data structure store.
- **MongoDB**: A NoSQL database service used for storing and retrieving data efficiently.
- **Docker**: A containerization tool.
- **Docker-compose**: A basic container orchestration tool.

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/shahtejas-t/url_shortener.git
   cd url_shortener
   ```

2. **Run Docker-compose**: Docker compose will build and start golang app, MongoDB, MongoDB Express and Redis
   ```bash
   docker-compose up -d
   ```

3. **Build the Project (Without Docker)**:
   ```bash
   go build -o url_shortener ./cmd/url_shortener
   ```

4. **Run the Application (Without Docker)**:
   ```bash
   ./url_shortener
   ```

## Usage :

### Accessing MongoDB Express :
Access this url through web browser
```
http://localhost:8081/
```
For password check docker-compose.yml file.

Default credentials :
- username : myuser
- password : mypassword

### Accessing the APIs :
1. **Generate Short url :**
```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"long": "https://www.google.com/"}' \
  http://localhost:8080/shorten
```

2. **Get Original URL from Shorten URL(5X0fXztP) :**
```
curl http://localhost:8080/5X0fXztP
```

3. **Delete the Shorten URL :**
```
curl -X DELETE http://localhost:8080/shorten?id=5X0fXztP
``` 

4. **Get top N records for most shorten urls (limit specifies top N records) :**
```
curl -X GET \
  -H "Content-Type: application/json" \
  -d '{"limit": 3}' \
  http://localhost:8080/toprecords

```

5. **Get total times a specific shorten url has been created :**
```
curl -X GET \
  -H "Content-Type: application/json" \
  -d '{"short_link": "5X0fXztP"}' \
  http://localhost:8080/linkcount
```

## TODO

- Enhancing inmemory caching
- Adding statistics for links
- Provide functionaity for monitoring and logging the application
