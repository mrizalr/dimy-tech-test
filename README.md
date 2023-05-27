# dimy-tech-test
How to run this project

Open project `cd dimy-tech-test`\n
Change .env files with yours configuration
Run program `go run main.go`

BASE URL `host.com:3001/api/v1`

Customer
GET `/customers`
GET `/customers/:id`
POST `/customers`
PATCH `/customers/:id`
DELETE `/customers/:id`

Address
GET `/addresses?customer_id=id`
GET `/addresses/:id`
POST `/addresses?customer_id=id`
PATCH `/addresses/:id`
DELETE `/addresses/:id`

Product
GET `/products`
GET `/products/:id`
POST `/products`
PATCH `/products/:id`
DELETE `/products/:id`

Payment Method
GET `/payment_methods`
GET `/payment_methods/:id`
POST `/payment_methods`
PATCH `/payment_methods/:id`
DELETE `/payment_methods/:id`

Order
GET `/orders`
GET `/orders/:id`
POST `/orders?customer_id=id`
PATCH `/orders/:id`
DELETE `/orders/:id`
