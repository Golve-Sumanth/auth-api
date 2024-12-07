# REST API for Authentication

This project provides a REST API for user authentication, including functionalities for signing up, signing in, token revocation, and token refresh.

## Prerequisites

Ensure that **Docker** and **Docker Compose** are installed on your machine. You can follow the installation guide here: [Docker Installation](https://docs.docker.com/get-docker/).

## Starting the API Service

To start the REST API and its dependencies, run the following command:

```bash
docker-compose up --build
```

## API Endpoints

Here are the endpoints you can use to interact with the REST API.

1. Sign Up

Registers a new user.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password123"}' http://localhost:8080/signup
```

2. Sign In

Authenticates the user and returns a JWT token.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password123"}' http://localhost:8080/signin
```

3. Access a Protected Route

Use the JWT token to access a protected resource.

```bash
curl -X GET -H "Authorization: Bearer <your_token>" http://localhost:8080/protected
```

4. Revoke a Token

Revokes a token, making it invalid for future use.

```bash
curl -X POST -H "Authorization: Bearer <your_token>" http://localhost:8080/revoke
```

5. Refresh a Token

Generates a new token using a valid (non-revoked) token.

```bash
curl -X POST -H "Authorization: Bearer <your_token>" http://localhost:8080/refresh
```


