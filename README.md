# Go-CRM-Backend  
- This is a simple CRM backend server written in Go for Udacity Go Language (Golang) Nanodegree Program final project.
- Launch the server: `go run main.go` we can see the server is running on port 3000 and we can open browser to http://localhost:3000.
- run `go test` in terminal we can see the test result all PASS.
## Create RESTful server endpoints for CRUD operations
HTTP Method | URL Path        | Description
----------- | --------------- | ----------------------------
GET         | /customers/{id} | Getting a single customer
GET         | /customers      | Getting all customers
POST        | /customers      | Creating a customer
PUT         | /customers/{id} | Updating a customer
DELETE      | /customers/{id} | Deleting a customer 

