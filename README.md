## **SIMPLE E-COMMERCE WITH MONGODB AND GIN FRAMEWORK**

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

[![Build Status](https://travis-ci.org/gin-gonic/gin.svg)](https://travis-ci.org/gin-gonic/gin)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Sourcegraph](https://sourcegraph.com/github.com/gin-gonic/gin/-/badge.svg)](https://sourcegraph.com/github.com/gin-gonic/gin?badge)
[![Open Source Helpers](https://www.codetriage.com/gin-gonic/gin/badges/users.svg)](https://www.codetriage.com/gin-gonic/gin)
[![Release](https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square)](https://github.com/gin-gonic/gin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin)](https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin)

# project run

```bash
# Start the mongodb container for local development
install laragon server on local system
go run main.go
```

#### ` Using Mongodb for small scale ecommerce industry is not a good idea instead use redis or mysql`

## API FUNCTIONALITY CURRENTLY ADDED?

- Signup 🔒
- Login 🔒
- Product listing General View 👀
- Adding the products to DB
- Sorting the products from DB using regex 👀
- Adding the products to cart 🛒
- Removing the Product from cart🛒
- Viewing the items in cart with total price🛒💰
- Adding the Home Address 🏠
- Adding the Work Address 🏢
- Editing the Address ✂️
- Deleting the Adress 🗑️
- Checkout the Items from Cart
- Buy Now products💰


## Configuration Database(database.go)

```go
// code example of url and port in databasetup.go
client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&3t.uriVersion=3&3t.connection.name=hadidev+-+imported+on+Aug+11%2C+2022&3t.alwaysShowAuthDB=true&3t.alwaysShowDBFromUserRole=true"))
// code example of defining collection name
var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
```

- **SIGNUP FUNCTION API CALL (POST REQUEST)**

http://localhost:8000/users/signup

```json
{
    "first_name": "Muhammad",
    "last_name": "Nur Hadi",
    "email": "hadi@gmail.com",
    "password": "hadi123",
    "phone": "+62882005424842"
}
```

Response :"Successfully Signed Up!!"

- **LOGIN FUNCTION API CALL (POST REQUEST)**

  http://localhost:8000/users/login


```json
{
  "email": "hadi@gmail.com",
  "password": "hadi123"
}
```

response will be like this

```json
{
    "_id": "63059f7cd51c4779126aa3a8",
    "first_name": "Muhammad",
    "last_name": "Nur Hadi",
    "password": "$2a$14$3M8zbRBEB58mKwvwfwSbT.aGRO3CsRYEaNQ6o.nFqdXrTWg/MCBfa",
    "email": "hadi@gmail.com",
    "phone": "+62882005424842",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImhhZGlAZ21haWwuY29tIiwiRmlyc3RfTmFtZSI6Ik11aGFtbWFkIiwiTGFzdF9OYW1lIjoiTnVyIEhhZGkiLCJVaWQiOiI2MzA1OWY3Y2Q1MWM0Nzc5MTI2YWEzYTgiLCJleHAiOjE2NjEzOTk1Njd9.z0gU6-Sq-uT-RqQ5f4AHYIVKTNZdMUHR7IjWaviMaQM",
    "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE2NjE5MTc5Njd9.6b9XRNNB-WV6t0m8aC7Yw746H48I8nc1_dCy0KkgUIQ",
    "created_at": "2022-08-24T03:48:12Z",
    "updtaed_at": "2022-08-24T03:48:12Z",
    "user_id": "63059f7cd51c4779126aa3a8",
    "usercart": [],
    "address": [],
    "orders": []
}
```

- **CART FUNCTION API CALL (GET REQUEST)**

  http://localhost:8000/list/cart?id=USERID

response will be like this

```json
3000
[
    {
        "Product_ID": "6305b2506df98551bd8d81a5",
        "product_name": "Tas Selempang",
        "price": 3000,
        "rating": 8,
        "image": "2.jpg",
        "stock": 30,
        "review": "Barang sangat istimewa dan pas dikantong",
        "detail": [
            {
                "color": "hitam",
                "corak": "Ninja Kagura",
                "motif": "Kagura memakai payung"
            },
            {
                "color": "Biru Dongker",
                "corak": "Batik Bunga",
                "motif": "Naga lagi tidur"
            }
        ]
    }
]
```
