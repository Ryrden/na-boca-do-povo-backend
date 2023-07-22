# Na Boca do povo - Backend

## Anotações de estudo sobre o backend do projetoThe main. go file just have the main function and the routes

- The `api folder` is responsible for the routes and the request/response
- The `controller folder` is responsible for controlling the flow of requests and responses
- The `service folder` is responsible for the business logic
- The `repository folder` is responsible for the database operations
- The `model folder` is responsible for the database models
- the `dto folder` is responsible for the data transfer objects, current there is one for the response
- The `database folder` is responsible for storing database script (.sql)

## How to run

- Install [Go](https://golang.org/doc/install)
- Install [Docker](https://docs.docker.com/get-docker/) with [Docker Compose](https://docs.docker.com/compose/install/)

Execute the following commands after cloning the repository:

```bash
# Start the database
docker-compose up -d

# Run the application
go run main.go
```

To stop the database, run: `docker-compose down`
