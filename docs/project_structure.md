# Introduction
This folder is a document about the architecture chosen to make the project. 

## Folder Structure
```bash
.
├── bin
│
├── cmd
│   └── migration
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
├── src
│   ├── base
│   │   ├── dto
│   │   ├── interfaces
│   │   ├── logger
│   │   └── mocks
│   │
│   ├── core
│   │   ├── entity
│   │   └── useCase
│   │
│   ├── external
│   │   ├── db
│   │   │   └── postgres
│   │   │
│   │   ├── http
│   │   │
│   │   └── rest
│   │       └── cmd
│   │
│   └── operation
│       ├── controller
│       │   └── web
│       │
│       ├── gateway
│       │   ├── db
│       │   └── http
│       │
│       └── presenter
│           ├── common
│           └── strings
│
├── migration
│
├── docker-compose.yaml
├── Jenkinsfile
├── Makefile
└── README.md

```
## Structure explanation
```
├── cmd 
│   └── migration
```

**cmd** - This folder is responsible for executables.
- **migration** - This folder is related to initialization of Database, for example to avoid starting with everything empty;  

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
│   │   ├── interfaces
│   │   ├── logger
│   │   └── mocks
│   │
│   ├── core
│   │   ├── entity
│   │   └── useCase
│   │
│   ├── external
│   │   ├── db
│   │   │   └── postgres
│   │   │
│   │   ├── http
│   │   │
│   │   └── rest
│   │       └── cmd
│   │
│   └── operation
│       ├── controller
│       │   └── web
│       │
│       ├── gateway
│       │   ├── db
│       │   └── http
│       │
│       └── presenter
│           ├── common
│           └── strings
│
```
- **src** - This folder is responsible to have all resources about the server, like: external connections, business logic, entity management, etc...
    - **base** - This folder is responsible to contains every shared data through the application, like interfaces and dtos, logger, etc...;
        - **dto** - This folder is related to data transfer object, this way is possible to communicate among layers using them without need to apply some rules to entities;
        - **interfaces** - This folder is related to contracts, where have the all connections that could be used by adapters or application;
        - **logger** - This folder is related to logger implementation;
        - **mocks** - This folder is related to mocks implementation to help in unity tests;

    - **core** - This folder is used to store the project's core or foundational code, which may include essential functionality;
        - **useCase** - This folder is related to services, where have the all business logic;
        - **entity** - This folder is related to store domain entity definitions;

    - **external** - This layer is responsible for external resources;
        - **db** - This folder contains every domain database implemetation, in our case: client, order, product, etc...;
            - **postgres** - This folder is related to postgresql database implementation;
        - **rest** - This layer contains every handler to call each controller;
            - **cmd** - This folder contains the executables;
        
    - **operation** - This layer is responsible for 
        - **controller** - This folder is related to handle user requests like: http1/2, RabbitMq, RPC, etc..;
            - **web** - This contains every domain http1/2 controller implementation;
        - **gateway** - This folder is responsible for interfaces and structures that define the desired interface about external connections: SQL, HTTP, RPC, Cache, etc...;  
            - **db** - This folder contains the gateway to access database implementation externally;  
            - **http** - This folder contains the gateway to access http implementation externally;  
        - **presenter** - This folder is responsible for bring helper functions or types that could be used in entities or transformation data;  
            - **strings** - This folder is related to helper stringfy functions;
            - **common** - This folder is related to store custom types to help entities and have some useful helper methods;
---
<br>

```
├── migration
```
**migration** - This folder is related to migration implementation;

