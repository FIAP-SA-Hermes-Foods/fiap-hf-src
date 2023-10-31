# Introduction
This is a document about the architecture chosen to make the project. 

## Folder Structure
```bash
├── cmd
│   ├── migration
│   └── server
├── docs
├── infrastructure
│   ├── db
│   │   ├── DML
│   │   ├── migration
│   │   └── postgres
│   ├── scripts
│   └── swagger_doc
├── internal
│   ├── adapters
│   │   ├── driven
│   │   │   └── repository
│   │   │       ├── client
│   │   │       ├── order
│   │   │       └── product
│   │   └── driver
│   │       └── http
│   └── core
│       ├── application
│       ├── domain
│       │   ├── entity
│       │   ├── http
│       │   ├── repository
│       │   └── valueObject
│       ├── service
│       └── ui

```
## Structure explanation
```
├── cmd 
│   ├── migration
│   └── server
```

**CMD** - This folder is responsible for executables;  
**Migration** - This is realated to initialization of Database, for example to avoid starting with everything empty;  
**Server** - This is realated to main and Mercado Pago API;  

---
<br>

```
├── infrastructure
│   ├── db
│   │   ├── DML
│   │   ├── migration
│   │   └── postgres
│   ├── scripts
│   └── swagger_doc
```
**Infrastructure** - This folder is responsible for all external resources that our application communicates with;  
**DB** - This is realated to database external resources;  
**DML** - This is realated to .SQL files that perform our data manipulation;  
**Scripts** - This is realated to scripts, for example setting up docker network;  
**Swagger_doc** - This is realated to swagger documentation;  

---
<br>

```
├── docs
```
**Docs** - This folder is responsible for documentation, for example: how to contribute in this project;

---
<br>


```
├─ internal
│   ├── adapters
│   │   ├── driven
│   │   │   └── repository
│   │   │       ├── client
│   │   │       ├── order
│   │   │       └── product
│   │   └── driver
│   │       └── http
│   └── core
│       ├── application
│       ├── domain
│       │   ├── entity
│       │   ├── http
│       │   ├── repository
│       │   └── valueObject
│       ├── service
│       └── ui
```
**Adapter** - This folder is responsible for interfaces and structures that define the desired interface;  
**Driven** - This is realated to out Database;  
**Repository** - This is realated to interfaces definition, logic related to data persistence and concrete implementations;  
**Driver/http** - This is realated to payment API;  
**Core** - This folder is used to store the project's core or foundational code, which may include essential functionality;  
**Application** - This is realated to the aplication, for example business logic;  
**Domain** - This is realated to project domain, for example the product struct;  
**Entity** - This is realated to store domain entity definitions;  
**Value Object** - This is realated to store definitions of types that represent value objects;  
**Service** - This is realated to services, for example OrderService;  
**UI** - This is realated to User Interface;  