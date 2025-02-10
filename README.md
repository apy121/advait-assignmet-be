# Authentication API

This is a simple REST API for user authentication.

## Setup

- Ensure Docker and Docker Compose are installed on your system.

# Awesome Project

### Prerequisites
Ensure you have the following installed:
- Go 1.23.5
- Docker and Docker Compose
---
### 🔧 Clone the Repository
```bash
git clone https://github.com/apy121/advait-assignmet-be.git
cd advait-assignmet-be
git checkout master
go mod tidy
go run cmd/main/main.go
```

## Run with Docker

Pull and run the backend application:
```bash
docker pull apy121/awesomeproject
docker run -p 8080:8080 apy121/awesomeproject
````

## POST /signup: Create a new user account

```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"user@example.com", "password":"password123"}' http://localhost:8080/signup
```

## POST /signin: Authenticate and get a JWT token

```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"user@example.com", "password":"password123"}' http://localhost:8080/signin
```

## GET /protected: Test token authorization (requires a valid token in Authorization header)

```bash
curl -X GET -H "Authorization: Bearer <token>" http://localhost:8080/protected
```

## POST /revoke: Revoke a token

```bash
curl -X POST -H "Authorization: Bearer <token>" http://localhost:8080/revoke
```

## POST /refresh: Refresh token

```bash
curl -X POST -H "Authorization: Bearer <token>" http://localhost:8080/refresh
```
