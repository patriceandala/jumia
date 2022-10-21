# JUMIA API
 
The following is a backend API by Patrice Opiyo built for the Jumia team.

## Tools used
Golang
Mysql
Swagger
Docker


## Project Structure

I've defined the following endpoints:

* /v1/product/{sku} [GET]
* /v1/product/{sku} [PATCH] -Consume a product. Checks if the product is available first
* /v1/products/  [PATCH]-allows bulk update using CSV file
 


### Local development

There are few tools that are required to run this server locally i.e: Go, docker, and docker-compose.

If you already have this installed you can run `make server` on jumia directory, this command will run docker-compose up and create the database then migrate the tables into the db
 