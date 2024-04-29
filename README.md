# url-shortener

A Go and Gin-based URL shortener application designed to provide users with a convenient way to shorten long URLs into shorter, more manageable links. The application follows the principles of hexagonal architecture pattern to ensure modularity, flexibility, and maintainability.

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

2. **Build the Project**:
   ```bash
   go build -o url_shortener ./cmd/url_shortener
   ```

3. **Run Docker-compose**:
   ```bash
   docker-compose up -d
   ```

4. **Run the Application**:
   ```bash
   ./url_shortener
   ```

## TODO

- Enhancing inmemory caching
- Adding statistics for links
- Provide functionaity for monitoring and logging the application
