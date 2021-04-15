# Lana Demo App

- Create a new checkout basket
- Add a product to a basket
- Get the total amount in a basket
- Remove the basket
- Get basket information

#### Create a new checkout basket
```
curl --request POST 'localhost:8085/api/basket/'
```

#### Add a product to a basket
```
curl --request PATCH 'localhost:8085/api/basket/1' \
--header 'Content-Type: application/json' \
--data-raw '{
"code": "{product-code}",
"quantity": {quantity}    
}'
```
  
    - Code (string) is mandatory and it should be one of the valid items: PEN, TSHIRT or MUG
    - Quantity (int) is optional and default value is 1 

#### Get the total amount in a basket
```
curl --request GET 'localhost:8085/api/basket/total/{basket-id}'
```

#### Remove the basket
```
curl --request DELETE 'localhost:8085/api/basket/{basket-id}'
```

#### Get basket information
```
curl --request GET 'localhost:8085/api/basket/{basket-id}'
```

### Technology stack
- Java 8
- Spring Boot 2.4.4
- JUnit 5

### How to run

##### Please note:
* The storage is in memory so, no database will be needed, but it could be added easily*
* Default port is set to 8085 (e.g: http://localhost:8085)
* It's needed to have installed `maven` and `java 8` for being able to run it

1. Execute the following command in the console:
   ```
   mvn spring-boot:run
   ```

2. Wait until the following message is shown:
    ```
   Started RobertomApplication...
   ```

3. Now you're ready to send requests :)

##### If you want to run in a docker env, then you'll need to exec:
1. Execute the following command in the console:
   ```
   mvn clean package
   ```
2. Create docker image and run the container
   ```
   docker build -t robertom-lana-challenge . && docker run -it  -p 8085:8085 robertom-lana-challenge
   ```

### Testing
There are some unit and integrations test under the project, that you can run them executing the following command: ```mvn clean test```

### Lana Backend Challenge Description

Lana has come to conclusion that users are very likely to buy awesome Lana merchandising from a physical store that
sells the following 3 products:

```
Code         | Name              |  Price
-----------------------------------------------
PEN          | Lana Pen          |   5.00€
TSHIRT       | Lana T-Shirt      |  20.00€
MUG          | Lana Coffee Mug   |   7.50€
```

Various departments have insisted on the following discounts:

* The sales department thinks a buy 2 get 1 free promotion will work best (buy two of the same product, get one free),
  and would like this to only apply to `PEN` items.

* The CFO insists that the best way to increase sales is with discounts on bulk purchases (buying x or more of a
  product, the price of that product is reduced), and requests that if you buy 3 or more `TSHIRT` items, the price per
  unit should be reduced by 25%.

Your task is to implement a simple checkout server and client that communicate over the network.

We'd expect the server to expose the following independent operations:

- Create a new checkout basket
- Add a product to a basket
- Get the total amount in a basket
- Remove the basket

The server must support concurrent invocations of those operations: any of them may be invoked at any time, while other
operations are still being performed, even for the same basket.

At this stage, the service shouldn't use any external databases of any kind, but it should be possible to add one easily
in the future.

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

**The solution should:**

- Build and execute in a Unix operating system.
- Focus on solving the business problem (less boilerplate!)
- Have a clear structure.
- Be easy to grow with new functionality.
- Don't include binaries, and use a dependency management tool.

**Bonus Points For:**

- Be written in Go (let us know if this is your first time!)
- Unit/Functional tests
- Dealing with money as integers
- Formatting money output
- Useful comments
- Documentation
- Docker images / CI
- Commit messages (include .git in zip)
- Thread-safety
- Clear scalability
