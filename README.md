# Quasar-Fire

## How to run
run the following command to start the server
```bash
export PORT="8080"
go run cmd/main.go
```

## Project Structure
### Packages
- healthcheck: contains the healthcheck handler to check if the server is up
- router: contains the router struct and the routes group, it also contains the gin router instance
- swagger: contains the swagger documentation
- topsecret: contains the topsecret domain to get the location and message from the satellites
### Technologies
- Gin: web framework
- Swagger: documentation
- SQLite: database
- Gorm: ORM
- Logrus: logger
### Architecture
Clean Architecture is used to separate the business logic from the framework and the database. The project is divided in 3 layers:
- Domain: contains the business logic about the position calculation and the message
- Infrastructure: contains the framework and the database, it is built with the following files:
  - routes file: contains the routes and the handlers
  - dto file: contains the data transfer objects
  - repository file: contains the database queries
- Service: contains the use case to get the location and message from the satellites


## API Heroku
- https://quasar-fire-1243502fea29.herokuapp.com/api/quasar-fire/v1/swagger/index.html#/