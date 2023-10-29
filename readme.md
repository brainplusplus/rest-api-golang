
# Simple E-Commerce RESTApi With Golang

A simple e-Commerce RESTApi with golang for golang assestment in PandaiSuite/Cognotiv


## Prerequisite

- Install Golang SDK in your computer
- Install MySQL in your computer
- Create Database, e.g database name is simple_ecommerce
- Import Sql file named `simple_ecommerce-db.sql` in folder `requirements`

```bash
  mysql -u username -p database_name < simple_ecommerce-db.sql
```
- Copy example.config.toml to config.toml (root project folder), and change configuration with your own configuration
```bash
  [server]
  mode = development | production
  port = <your_desired_port>
  debug = true
  log_path = <your_desired_log_path>
  timezone = <your_desired_timezone>
```

if `mode = production`, the log will append in file you choose in `log_path`
else the log will append in console

```bash
  [smtp]
  username = <your_desired_username>
  password = <your_desired_password>
  port = <your_desired_port>
  host = <your_desired_host>
  from_sender = <your_desired_from_sender>
```
This app using SMTP for sending email, please change with your own configuration

```bash
  [jwt]
  timeout = <your_desired_jwt_timespan>
  signature_key = <your_desired_singnature_key_of_jwt>
  issuer = <your_desired_issuer>
```
This app using JWT as authentication

```bash
  [database]
  db_driver = "mysql"
  db_url = "username:password@(host:port)/dbName?parseTime=true"
  props_max_idle = 3
  props_max_conn = 3
  props_max_lifetime = 3
```
This App using mysql, so you can adjust the configuration

## Deployment

To deploy this project run

```bash
  go build -o simple_ecommerce.exe
  .\simple_ecommerce.exe
```

Or if you just want to run golang project (dev)

```bash
  go run main.go
```
## API Reference

i've export POSTMAN for this project, you can use and import it to your POSTMAN, the file for export named `simple_ecommerce.postman_collection.json` in folder `requirements`

### Authentication API

#### Register Customer

```http
  POST /api/auth/register
```

| Parameter | Type     | Is Required                | Description                |
| :-------- | :------- | :------------------------- | :----------- |
| `name` | `string` | **Yes**. | Name of Customer |
| `email` | `string` | **Yes**. | Email of Customer as Username |
| `password` | `string` | **Yes**. | Password of Customer |

#### Login Customer to get JWT Token

```http
  POST /api/auth/login
```

| Parameter | Type     | Is Required                | Description                |
| :-------- | :------- | :------------------------- | :----------- |
| `email` | `string` | **Yes**. | Email of Customer as Username |
| `password` | `string` | **Yes**. | Password of Customer |

#### Login Admin to get JWT Token

```http
  POST /api/auth/login_admin
```

| Parameter | Type     | Is Required                | Description                |
| :-------- | :------- | :------------------------- | :----------- |
| `email` | `string` | **Yes**. | Email of Customer as Username |
| `password` | `string` | **Yes**. | Password of Customer |

### Customer API

After Get JWT Token From `POST /api/auth/login`, you can access Customer API

```http
  Authorization:  Bearer <your_Customer_JWT_Token>
```

#### Get all products
You'll get all products can order

```http
  Authorization:  Bearer <your_Customer_JWT_Token>
  GET /api/customer/product
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `None` | `-` | No Parameter |

#### Get a Product
You'll get a product detail

```http
  Authorization:  Bearer <your_Customer_JWT_Token>
  GET /api/customer/product/${id}
```

| Parameter | Type     | Is Required                | Description                |
| :-------- | :------- | :------------------------- | :----------- |
| `id` | `string` | **Yes**. | Id of Product to fetch |

#### Create an Order
You can create an order one or more product

```http
  Authorization:  Bearer <your_Customer_JWT_Token>
  POST /api/customer/order
```

| Parameter | Type     | Is Required                | Description                |
| :-------- | :------- | :------------------------- | :----------- |
| `order_products` | `array` of order_product | **Yes**. | List Of Ordered Product |

order_product Struct

| Parameter | Type     | Is Required                | Description                |
| :-------- | :------- | :------------------------- | :----------- |
| `product_id` | `int` | **Yes**. | Id Of Product |
| `quantity` | `int` | **Yes**. | Quantity Of Ordered Product |


This is sample body request (JSON) for this API
```http
{
    "order_products" :[
        {
            "product_id" : 2,
            "quantity" : 2
        },
         {
            "product_id" : 3,
            "quantity" : 1
        }
    ]
}
```

#### Get all order
You'll get all order of current customer (based on JWT Token/customer id)

```http
  Authorization:  Bearer <your_Customer_JWT_Token>
  GET /api/customer/order
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `None` | `-` | No Parameter |

#### Get order detail by id
You'll get an order detail by id

```http
  Authorization:  Bearer <your_Customer_JWT_Token>
  GET /api/customer/order
```

| Parameter | Type     | Is Required                | Description                |
| :-------- | :------- | :------------------------- | :----------- |
| `id` | `string` | **Yes**. | Id of Order to fetch |


### Admin API

After Get JWT Token From `POST /api/auth/login_admin`, you can access Admin API

```http
  Authorization:  Bearer <your_Admin_JWT_Token>
```

#### Get all order
You'll get all order of all customer

```http
  Authorization:  Bearer <your_Admin_JWT_Token>
  GET /api/admin/order
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `None` | `-` | No Parameter |

## Features

- Validation using https://github.com/go-playground/validator
- Database access using https://github.com/jmoiron/sqlx
- Web Framework (Router, Limitter, etc) using https://github.com/gofiber/fiber
- Database using MySQL
- JWT Library using https://github.com/golang-jwt/jwt
- Cron Scheduler using https://github.com/robfig/cron
- Configuration File Reader using https://github.com/spf13/viper
- SMTP Library using https://gopkg.in/gomail.v2


## Contact

For contact, email akhmad.mib@gmail.com

