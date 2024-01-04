# Introduction
This folder is a document about the architecture chosen to make the project. 

## Folder Structure
```bash
.
├── bin
├── cmd
│   ├── migration
│   └── server
│
├── docs
│   ├── postman_collection
│   └── swagger
│
├── infrastructure
│   ├── db
│   │   └── DML
│   │
│   ├── docker
│   │   ├── go
│   │   │   └── Dockerfile
│   │   ├── postgres
│   │   │   └── Dockerfile
│   │   └── swagger
│   │       └── Dockerfile
│   │
│   ├── kubernetes
│   │   ├── config
│   │   ├── deployment
│   │   └── volume
│   │
│   └── scripts
│   
├── internal
│   ├── adapters
│   │   ├── driven
│   │   │   └── repository
│   │   │       ├── client
│   │   │       ├── order
│   │   │       ├── order_product
│   │   │       ├── product
│   │   │       └── voucher
│   │   │
│   │   └── driver
│   │       └── http
│   │           └── api-mercadoPago
│   │
│   ├── core
│   │   ├── application
│   │   ├── entity
│   │   │   └── common
│   │   ├── service
│   │   └── useCase
│   │       ├── db
│   │       ├── http
│   │       └── repository
│   │   
│   └── handler
│       └── web
│
├── pkg
│   ├── logger
│   ├── migration
│   └── postgres
│
├── Makefile
├── Jenkinsfile
├── .env
├── docker-compose.yaml
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
│   │   │   └── Dockerfile
│   │   ├── postgres
│   │   │   └── Dockerfile
│   │   └── swagger
│   │       └── Dockerfile
│   │
│   ├── kubernetes
│   │   ├── config
│   │   ├── deployment
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
├── internal
│   ├── adapters
│   │   ├── driven
│   │   │   └── repository
│   │   │       ├── client
│   │   │       ├── order
│   │   │       ├── order_product
│   │   │       ├── product
│   │   │       └── voucher
│   │   │
│   │   └── driver
│   │       └── http
│   │           └── api-mercadoPago
│   │
│   ├── core
│   │   ├── application
│   │   ├── entity
│   │   │   └── common
│   │   ├── service
│   │   └── useCase
│   │       ├── db
│   │       ├── http
│   │       └── repository
│   │   
│   └── handler
│       └── web
```
**internal** - This folder is responsible to have all resources about the server, like: external connections, business logic, entity management, etc...
- **adapter** - This folder is responsible for interfaces and structures that define the desired interface about external connections related to persistence like: SQL, NoSQL, Cache, etc...;  
    - **driven** - This folder is related to out Database;  
        - **repository** - This folder is related to interfaces definition, logic related to data persistence and concrete implementations;  
    - **driver** - This folder is responsible for interfaces and structures that define the desired interface about external connections not related to persistence like: HTTP, RPC, Message Broker;  
        - **http** - This folder is related to payment API;
- **core** - This folder is used to store the project's core or foundational code, which may include essential functionality;
    - **application** - This folder is related to the aplication, where have business logic that all layers communicate. Example: (application -> service -> adapters -> driver -> http -> api-mercadoPago -> adapters -> repository);
    - **entity** - This folder is related to store domain entity definitions;  
        - **common** - This folder is related to store definitions of structures that could be used by entities;  
    - **service** - This folder is related to services, where have the all business logic about one entity. Example: OrderService;  
    - **useCase** - This folder is related to contracts, where have the all connections that could be used by adapters or application;  
- **handler** - This folder is related to handle user requests like: http1/2, RabbitMq, RPC, etc..;
    - **web** - This folder is related to handle user http requests.

---
<br>

```
├── pkg
│   ├── logger
│   ├── migration
│   └── postgres
```
**pkg** - This folder is responsible for all external resources that our application communicates with.
- **logger** - This folder is related to logger implementation;
- **migration** - This folder is related to migration implementation;
- **postgres** - This folder is related to postgresql database implementation;


