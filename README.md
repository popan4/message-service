# message-service
A simple Restful API built in Go lang using Fiber framework and in memory map as a db for managing messages and checking along whether they are Palindromes.

A service that allows you to create, read, update, and delete messages, as well as check if a message is a palindrome.

features:
- Create a message
- list all messages (not added pagination)
- add check if a message is a palindrome
- update a message
- delete a message
- cofigurable port
- unit tests

Technologies used:
- Go
- Fiber framework
- In-memory map as a database
- Golang testing package
- configuration using viper
- API documentation
- zap for logging

Project Structure:
```
message-service
├── cmd
│   └── main.go
├── config
│   └── config.go
├── controllers
│   └── message_controller.go
├── models
│   └── message.go
├── routes
│   └── message_routes.go
├── services
│   └── message_service.go
├── utils
│   └── palindrome_checker.go
├── go.mod
├── go.sum
└── README.md
```

Api Endpoints:
- POST /createMsg - Create a new message
- GET /getAllMsg - Get all messages
- GET /getMsg/:id - Get a message by ID
- PUT /updateMsg/:id - Update a message by ID
- DELETE /deleteMsg/:id - Delete a message by ID

# How to run the project
1. Clone the repository:
   ```bash
   git clone
   cd message-service
   ```
2. Install dependencies:
    ```bash
    go mod tidy
   ```
3. Run the application:
   ```bash
   go run cmd/main.go
   ```
   
run the tests:
   ```bash
     go test ./... or
     go test ./... -cover 
   ```

# How to build the project
1. Build the application:
   ```bash
   docker build -t message-api .
   ```
2. Run the application:
   ```bash
    docker run -p 8080:8080 message-api
    ```
   
Deployment:
You can deploy this service on any cloud provider that supports Go applications, such as AWS, Google
Cloud.