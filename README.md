# URL Shortener Service

## Description

This is a URL shortener service built using the following technologies:

- Gin for the web framework
- Gorm for ORM
- PostgreSQL for the database
- Redis for caching
- JWT for authentication
- Rate limiting for request control

## Folder Structure
```
url-shortener/
├── config/
│ └── config.go
├── controllers/
│ ├── auth_controller_test.go
│ ├── auth_controller.go
│ ├── url_controller_test.go
│ ├── url_controller.go
├── middlewares/
│ ├── auth_middleware.go
│ ├── cors.go
│ └── rate_limiter.go
├── models/
│ ├── url.go
│ └── user.go
├── routes/
│ ├── auth_routes.go
│ ├── url_routes.go
│ └── router.go
├── utils/
│ ├── token.go
│ ├── password.go
│ └── redis.go
| └── validation.go
├── main.go
├── .env
├── go.mod
└── go.sum
```


## Setup

1. **Clone the repository:**
    ```sh
    git clone https://github.com/clim-bot/url-shortener.git
    cd url-shortener
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    brew install redis
    docker run --name redis -d -p 6379:6379 redis
    ```

3. **Setup environment variables:**
    Create a `.env` file in the root directory and add the following variables:
    ```
    DB_DSN=postgres://username:password@localhost:5432/yourdb
    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=
    ```

4. **Run the application:**
    ```sh
    go run main.go
    ```

## Endpoints

### Register a User

**URL**: `/auth/register`  
**Method**: `POST`  
**Description**: Register a new user

**Request Body**:
```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Curl Command**:
```sh
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{
  "username": "your_username",
  "password": "your_password"
}'
```

### Login a User

**URL**: `/auth/login`  
**Method**: `POST`  
**Description**: Login an existing user

**Request Body**:
```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Curl Command**:
```sh
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{
  "username": "your_username",
  "password": "your_password"
}'
```

### Logout a User

**URL**: `/auth/logout`  
**Method**: `POST`  
**Description**: Logout the current user

**Request Headers**:
- `Authorization`: `Bearer <JWT_TOKEN>`

**Curl Command**:
```sh
curl -X POST http://localhost:8080/auth/logout \
-H "Authorization: Bearer <JWT_TOKEN>"
```

### Create a Short URL

**URL**: `/url/shorten`  
**Method**: `POST`  
**Description**: Create a shortened URL
**Request Headers**:
- `Authorization`: `Bearer <JWT_TOKEN>`

**Request Body**:
```json
{
  "original_url": "http://example.com"
}
```

**Curl Command**:
```sh
curl -X POST http://localhost:8080/url/shorten \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <JWT_TOKEN>" \
-d '{
  "original_url": "http://example.com"
}'
```

### Get Original URL

**URL**: `/url/:shorten`  
**Method**: `GET`  
**Description**: Retrieve the original URL using the short code
**Request Headers**:
- `Authorization`: `Bearer <JWT_TOKEN>`

**Curl Command**:
```sh
curl -X GET http://localhost:8080/url/<SHORT_CODE>
```
Replace `<JWT_TOKEN>` with the JWT token you receive from the login endpoint, and `<SHORT_CODE>` with the short code you receive from the URL shortening endpoint.

### Unit Tests
To run the unit tests, use the following command:
```sh
go test ./...
```

### License
This project is licensed under the MIT License.
