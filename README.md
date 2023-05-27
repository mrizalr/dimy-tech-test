# dimy-tech-test
How to run this project<br/>

Open project `cd dimy-tech-test` <br/><br/>
Change .env files with yours configuration <br/><br/>
Run program `go run main.go` <br/><br/><br/>

BASE URL `host.com:3001/api/v1`<br/><br/>

Customer<br/>
GET `/customers`<br/>
GET `/customers/:id`<br/>
POST `/customers`<br/>
PATCH `/customers/:id`<br/>
DELETE `/customers/:id`<br/><br/>

Address<br/>
GET `/addresses?customer_id=id`<br/>
GET `/addresses/:id`<br/>
POST `/addresses?customer_id=id`<br/>
PATCH `/addresses/:id`<br/>
DELETE `/addresses/:id`<br/><br/>

Product<br/>
GET `/products`<br/>
GET `/products/:id`<br/>
POST `/products`<br/>
PATCH `/products/:id`<br/>
DELETE `/products/:id`<br/><br/>

Payment Method<br/>
GET `/payment_methods`<br/>
GET `/payment_methods/:id`<br/>
POST `/payment_methods`<br/>
PATCH `/payment_methods/:id`<br/>
DELETE `/payment_methods/:id`<br/><br/>

Order<br/>
GET `/orders`<br/>
GET `/orders/:id`<br/>
POST `/orders?customer_id=id`<br/>
PATCH `/orders/:id`<br/>
DELETE `/orders/:id`<br/><br/>
