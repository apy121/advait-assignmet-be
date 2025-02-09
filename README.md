# Authentication API

This is a simple REST API for user authentication.

## Setup

- Ensure Docker and Docker Compose are installed on your system.

## Running the Application

```bash
docker-compose up --build
```

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