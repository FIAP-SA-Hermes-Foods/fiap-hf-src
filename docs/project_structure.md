# Introduction
This folder is a document about the architecture chosen to make the project. 

## Folder Structure
```bash

.
├── bin
├── cmd
│   └── migration
│
├── docs
│   ├── postman_collection
│   └── swagger
│
├── infrastructure
│   ├── db
│   │   └── DML
│   │
│   ├── docker
│   │   ├── go
│   │   ├── postgres
│   │   └── swagger
│   │
│   ├── kubernetes
│   │   ├── config
│   │   ├── deployment
│   │   ├── hpa
│   │   └── volume
│   │
│   └── scripts
│
├── migration
│
├── src
│   ├── base
│   │   ├── dto
│   │   └── interfaces
│   │
│   ├── core
│   │   ├── entity
│   │   │
│   │   └── useCase
│   │       ├── application
│   │       ├── client
│   │       ├── order
│   │       ├── order_product
│   │       ├── product
│   │       └── voucher
│   │
│   ├── external
│   │   ├── cmd
│   │   │   └── server
│   │   ├── db
│   │   │   └── postgres
│   │   └── logger
│   │
│   └── operation
│       ├── controller
│       │   └── web
│       │
│       ├── gateway
│       │   ├── http
│       │   │   └── api-mercadoPago
│       │   │
│       │   └── repository
│       │       ├── client
│       │       ├── order
│       │       ├── order_product
│       │       ├── product
│       │       └── voucher
│       │
│       └── presenter
│           ├── common
│           └── strings
│
├── Jenkinsfile
├── docker-compose.yaml
├── Makefile
└── README.md
```
## Structure explanation
```
├── cmd 
│   ├── migration
│   └── server
```

**cmd** - This folder is responsible for executables.
- **migration** - This folder is related to initialization of Database, for example to avoid starting with everything empty;  
- **server** - This folder is related to Mercado Pago Mock and Hermes Foods API;  

---
<br>

```
├── docs
│   ├── postman_collection
│   └── swagger
```
**docs** - This folder is responsible for documentation, for example: how to contribute in this project, postman collections and swagger json template;
- **postman_collection** - This folder is related to postman_collection json collection to import requests easily;  
- **swagger** - This folder is related to swagger json template;  

---
<br>

```
├── infrastructure
│   ├── db
│   │   └── DML
│   │
│   ├── docker
│   │   ├── go
│   │   ├── postgres
│   │   └── swagger
│   │
│   ├── kubernetes
│   │   ├── config
│   │   ├── deployment
│   │   ├── hpa
│   │   └── volume
│   │
│   └── scripts
```
**infrastructure** - This folder is responsible for all infrastructure configs like: docker images, manifest and helper scripts.
- **db** - This folder is related to database config;  
    - **DML** - This folder is related to .SQL files that perform our data manipulation;  
- **kubernetes** - This folder is related to app manifest;  
    - **config** - This folder is related to ConfigMaps;
    - **deployment** - This folder is related to Deployment and Service;
    - **volume** - This folder is related to Volumes kind PV (Persistent Volume) and PVC (Persistent Volume Claim);
- **docker** - This folder is related to docker images;
- **scripts** - This folder is related to scripts, for example setting up docker network;

---
<br>


```
├── src
│   ├── base
│   │   ├── dto
│   │   └── interfaces
│   │
│   ├── core
│   │   ├── entity
│   │   │
│   │   └── useCase
│   │       ├── application
│   │       ├── client
│   │       ├── order
│   │       ├── order_product
│   │       ├── product
│   │       └── voucher
│   │
│   ├── external
│   │   ├── cmd
│   │   │   └── server
│   │   │
│   │   ├── db
│   │   │   └── postgres
│   │   │
│   │   └── logger
│   │
│   └── operation
│       ├── controller
│       │   └── web
│       │
│       ├── gateway
│       │   ├── http
│       │   │   └── api-mercadoPago
│       │   │
│       │   └── repository
│       │       ├── client
│       │       ├── order
│       │       ├── order_product
│       │       ├── product
│       │       └── voucher
│       │
│       └── presenter
│           ├── common
│           └── strings
```
- **src** - This folder is responsible to have all resources about the server, like: external connections, business logic, entity management, etc...
    - **base** - This folder is responsible to contains every shared data, like interfaces and dtos;
        - **dto** - This folder is related to data transfer object, this way is possible to communicate among layers using them without need to apply some rules to entities;
        - **interfaces** - This folder is related to contracts, where have the all connections that could be used by adapters or application;

    - **core** - This folder is used to store the project's core or foundational code, which may include essential functionality;
        - **useCase** - This folder is related to services, where have the all business logic;
            - **application** - This folder is related to the aplication, where have business logic that all layers communicate. Example: (application -> useCase -> gateways -> http -> api-mercadoPago -> gateways -> repository);
            - **client** - This folder is related to the business logic in client layer;
            - **order** - This folder is related to the business logic in order layer;
            - **order_product** - This folder is related to the business logic in order_product layer;
            - **product** - This folder is related to the business logic in product layer;
            - **voucher** - This folder is related to the business logic in voucher layer;
        - **entity** - This folder is related to store domain entity definitions;

    - **external** - This layer is responsible for external resources;
        - **cmd** - This folder is responsible for executables.
            - **server** - This folder is related to Mercado Pago Mock and Hermes Foods API;  
        - **db**
            - **postgres** - This folder is related to postgresql database implementation;
        - **logger** - This folder is related to logger implementation;

        
    - **operation** - This layer is responsible for 
        - **controller** - This folder is related to handle user requests like: http1/2, RabbitMq, RPC, etc..;
            - **web** - This folder is related to handle user http requests.
        - **gateway** - This folder is responsible for interfaces and structures that define the desired interface about external connections: SQL, HTTP, RPC, Cache, etc...;  
            - **repository** - This folder is related to interfaces definition, logic related to data persistence and concrete implementations;  
            - **http** - This folder is related to interfaces definition, logic related to data http implementations;  
        - **presenter** - This folder is responsible for bring helper functions or types that could be used in entities or transformation data;  
            - **strings** - This folder is related to helper stringfy functions;  
            - **common** - This folder is related to store custom types to help entities and have some useful helper methods;
---
<br>

```
├── migration
```
**migration** - This folder is related to migration implementation;



