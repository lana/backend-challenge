Lana has come to conclusion that users are very likely to buy awesome Lana merchandising from a physical store that sells the following 3 products:

```
Code         | Name              |  Price
-----------------------------------------------
PEN          | Lana Pen          |   5.00€
TSHIRT       | Lana T-Shirt      |  20.00€
MUG          | Lana Coffee Mug   |   7.50€
```

Various departments have insisted on the following discounts:

 * The sales department thinks a buy 2 get 1 free promotion will work best (buy two of the same product, get one free), and would like this to only apply to `PEN` items.

 * The CFO insists that the best way to increase sales is with discounts on bulk purchases (buying x or more of a product, the price of that product is reduced), and requests that if you buy 3 or more `TSHIRT` items, the price per unit should be reduced by 25%.

Your task is to implement a simple checkout server and client that communicate over the network.

We'd expect the server to expose the following independent operations:

- Create a new checkout basket
- Add a product to a basket
- Get the total amount in a basket
- Remove the basket

The server must support concurrent invocations of those operations: any of them may be invoked at any time, while other operations are still being performed, even for the same basket.

At this stage, the service shouldn't use any external databases of any kind, but it should be possible to add one easily in the future.

Implement a checkout service and its client that fulfills these requirements.

Examples:

    Items: PEN, TSHIRT, MUG
    Total: 32.50€

    Items: PEN, TSHIRT, PEN
    Total: 25.00€

    Items: TSHIRT, TSHIRT, TSHIRT, PEN, TSHIRT
    Total: 65.00€

    Items: PEN, TSHIRT, PEN, PEN, MUG, TSHIRT, TSHIRT
    Total: 62.50€

**Setup:**
- Run follow command:
~~~bash
make setup
~~~

## Creation a Docker

~~~bash
make build-docker
~~~

## Testing

~~~bash
make test
~~~

## Run app

~~~bash
docker-compose up
~~~

## Stop app

~~~bash
docker-compose down
~~~

## Endpoints

name                                   method          description
- /health                              GET             Check status of app is (live/died)

- /baskets                             POST            Create a new basket
- /baskets/:id                         GET             Get a basket
- /baskets/:id                         DELETE          delete a basket

- /baskets/:id/products/:code          POST            return basket with a new product 

- /baskets/:id/products/:code          DELETE          Return basket without this product

- /baskets/:id/checkout                POST            will close the basket and calculate the discount
                                                       Return basket without this product


## Client

To run client
~~~bash
go run cmd/client/cli.go

Usage:
  app [command]

Examples:
you can us the follow commands: create/add/remove/checkout

Available Commands:
  basket      call different operations
Flags:
  -h, --help   help for app


~~~
EXAMPLES:

* create a new basket
~~~bash
go run cmd/client/cli.go basket create

output:

Basket created
{"Code":"f855f846-5057-11ec-b55b-1e003b1e5256","Items":{},"Total":0,"Close":false}
~~~

* remove basket
~~~bash
go run cmd/client/cli.go basket remove fa4ae6e8-5057-11ec-b55b-1e003b1e5256

output:

basket ID deleted
~~~

* add a new product to a basket
~~~bash
go run cmd/client/cli.go basket add f855f846-5057-11ec-b55b-1e003b1e5256 Pen

output:

product added
~~~

* close basket and get total amount
~~~bash
go run cmd/client/cli.go basket checkout f855f846-5057-11ec-b55b-1e003b1e5256

output:

Basket ID: f855f846-5057-11ec-b55b-1e003b1e5256
Items:
      Item: Mug
      Quantity: 1      Unit price: 7.5
      Total With Discount:         7.5
      Item: Pen
      Quantity: 4      Unit price: 5
      Total With Discount:         15
      Item: Tshirt
      Quantity: 4      Unit price: 20
      Total With Discount:         60
----------------------------------------
Amount Total: 82.5

~~~

