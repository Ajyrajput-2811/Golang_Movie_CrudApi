# Golang_Movie_CrudApi
This is a simple Go (Golang) application that demonstrates how to build a basic CRUD (Create, Read, Update, Delete) API for managing movie records. It uses the Gorilla Mux router for handling HTTP requests. In this example, movie data is stored in-memory, but it can easily be adapted to use a database.
# Table of Contents

- Getting Started
  - Prerequisites
  - Installation
- Usage
  - Starting the Server
  - API Endpoints
- Data Model
- Contributing
## Getting Started
### Prerequisites
Before you can run this project, make sure you have the following installed:
- Go (Golang): [Download and Install Go]
1. Clone this repository to your local machine:
```
git clone https://github.com/your-username/movie-crud-api.git
```
2. Change to the project directory:
```
cd movie-crud-api
```
3. Install project dependencies using Go modules:
```
go mod download
```
## Usage

### Starting the Server
You can start the Go server by running the following command:
```
go run main.go
```
The server will start and listen on port 8080. You can access the API at [http://localhost:8080.]
## API Endpoints
The API provides the following endpoints for managing movie records:
- GET /movies: Retrieve a list of all movies.
- GET /movies/{id}: Retrieve a specific movie by its ID.
- POST /movies: Create a new movie.
- PUT /movies/{id}: Update a movie by its ID.
- DELETE /movies/{id}: Delete a movie by its ID.
  
You can use an API client (e.g., Postman or cURL) to make HTTP requests to these endpoints.
## Data Model

In this example, movie data is stored in-memory. The data model includes the following entities:

- Movie: Represents a movie record with attributes like ID, ISBN, Title, and Director.
- Director: Represents the director of a movie with attributes like Firstname and Lastname.
In a real-world application, you would typically use a database to store and manage movie records.
## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Submit a pull request to the main repository.
   
Please ensure your code adheres to the project's coding standards and includes appropriate documentation and tests.








