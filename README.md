# Golang REST API Boilerplate / Template

## About

This boilerplate/template is using not using Framework, just Golang Standard Library.
This too using `GROUP BY FUNCTION` folder structure, similar with `MVC (Model View Controller)`, which contains:

- Handlers/Controller: Handling user request for each function
- Middlewares: List middlewares
- Model: Handling datatype for each function
- Repositories: Handling data to database
- Routes: Setup routes endpoint for each function
- Storage: Setup storage or connection to database
- Utils: Reusable function or helper function

## Libraries

- [https://pkg.go.dev/net/http@go1.22.1](https://pkg.go.dev/net/http@go1.22.1): Golang Standard Library for Handling HTTP Request, Response and Other
- [https://github.com/joho/godotenv](https://github.com/joho/godotenv): Read from .env file
- [https://github.com/lib/pq](https://github.com/lib/pq): Database Driver for PostgreSQL
- [https://github.com/gorilla/mux](https://github.com/gorilla/mux): Handling routing
