## GCC compiller is required to run this project. To install, follow this guide:
- https://code.visualstudio.com/docs/cpp/config-mingw

## To run this project execute following commands while in project directory

- `cd path/to/project/directory/ebinzes_go`
- `go mod init (repository name)`, in this case `go mod init github.com/RobertOchmanek/ebiznes_go`
- `go get github.com/labstack/echo/v4`
- `go get github.com/mattn/go-sqlite3`
- `go get -u github.com/jinzhu/gorm`
- `go build main.go`
- `./main`

## CRUD operations:

### Categories:
- `curl -v -X GET http://localhost:8080/categories` get all categories (without products)
- `curl -v -X GET http://localhost:8080/categories/:id` get category by id (with products)
- `curl -v -X POST http://localhost:8080/categories -H 'Content-Type: application/json' -d '{"Name":""}'` add new category
- `curl -v -X PUT http://localhost:8080/categories/:id -H 'Content-Type: application/json' -d '{"Name":""}'` update existing category by id
- `curl -v -X DELETE http://localhost:8080/categories/:id` delete category by id (executed only if category doesn't contain any products)

### Orders:
- `curl -v -X GET http://localhost:8080/orders` get all orders
- `curl -v -X GET http://localhost:8080/orders/:id` get order by id
- `curl -v -X POST http://localhost:8080/orders -H 'Content-Type: application/json' -d '{"UserId":0,"Payment":{"Accepted":true,"PaymentType":0},"OrderItems":[{"ProductId":0,"Quantity":0}]}'` add new order (payment type must be one of 'paymentType` enumeration values)

### Products:
- `curl -v -X GET http://localhost:8080/products` get all products
- `curl -v -X GET http://localhost:8080/products/:id` get product by id
- `curl -v -X POST http://localhost:8080/products -H 'Content-Type: application/json' -d '{"CategoryId":0,"Name":"","Price":0}'` add new product
- `curl -v -X PUT http://localhost:8080/products/:id -H 'Content-Type: application/json' -d '{"CategoryId":0,"Name":"","Price":0}'` update existing product by id

### Users:
- `curl -v -X GET http://localhost:8080/users` get all users (without cart and orders)
- `curl -v -X GET http://localhost:8080/users/:id` get user by id (with cart and orders)
- `curl -v -X POST http://localhost:8080/users -H 'Content-Type: application/json' -d '{"Username":"","Email":""}'` add new user (with new cart)
- `curl -v -X PUT http://localhost:8080/users/:id -H 'Content-Type: application/json' -d '{"Username":"","Email":""}'` update existing user by id

### Cart:
- `curl -v -X GET http://localhost:8080/cartItems/:userId` get cart items by user id
- `curl -v -X PUT http://localhost:8080/cart -H 'Content-Type: application/json' -d '{"UserId":0,"CartItems":[{"ID":0,"CategoryId":0,"Name":"","Price":0,"Image":"","Quantity":0,"CreatedAt":"","UpdatedAt":"","DelatedAt":""}]}'` update existing cart (cart items not included in updated cart but existing in DB are removed, other items are saved or updated if they exist)