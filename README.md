# Board CRUD API using Golang Gin
* This is a web service project built with Golang Gin framework that allows users to register and view their portfolios.
* Ref : https://github.com/MongLong0214/BoardCRUD

## Prerequisites
* Golang 1.2 or higher
* Gin framework
* MySQL

## Getting Started
1. Clone this repository to your local machine.
2. Install the necessary dependencies by running the following command:

```bash
go mod download
```

3. Set up the environment variables by creating a .env file in the root directory of the project and adding the following variables:


```bash
MYSQL_URI=<MySQL uri>
JWT_SECRET=<your jwt secret key>
```

4. Run the project by running the following command:
```bash
go run main.go
```

Open your web browser and navigate to http://localhost:5001 to access the web service.