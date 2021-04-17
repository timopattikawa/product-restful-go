# API Spec

## Create Product

Request :
- Method : POST
- Endpoint : `/api/products`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json 
{
    "name" : "long",
    "price" : "string",
    "quantity" : "integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
         "name" : "long",
         "price" : "string",
         "quantity" : "integer",
     }
}
```

## Get Product

Request :
- Method : GET
- Endpoint : `/api/products/{id_product}`
- Header :
  - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
         "id" : "long, unique",
         "name" : "string",
         "price" : "string",
         "quantity" : "integer",
     }
}
```

## Update Product

Request :
- Method : PUT
- Endpoint : `/api/products/{id_product}`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json 
{
    "name" : "string",
    "price" : "long",
    "quantity" : "integer"
}
```

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : {
         "id" : "long, unique",
         "name" : "string",
         "price" : "string",
         "quantity" : "integer",
     }
}
```

## List Product

Request :
- Method : GET
- Endpoint : `/api/products`
- Header :
  - Accept: application/json
- Query Param :
  - size : number,
  - page : number

Response :

```json 
{
    "code" : "number",
    "status" : "string",
    "data" : [
        {
             "id" : "long, unique",
             "name" : "string",
             "price" : "string",
             "quantity" : "integer",
        },
        {
             "id" : "long, unique",
             "name" : "string",
             "price" : "string",
             "quantity" : "integer",
         }
    ]
}
```

## Delete Product

Request :
- Method : DELETE
- Endpoint : `/api/products/{id_product}`
- Header :
  - Accept: application/json

Response :

```json 
{
    "code" : "number",
    "status" : "string"
}
```
