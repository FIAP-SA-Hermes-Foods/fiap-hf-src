# Introduction

This file is about the database model chosen based on business logic provided by postech.

<img src='./img/hf-db-model.png" width="600px" height="500px">

## Database SQL Amazon RDS with Postgres

We decided to migrate our Postgres database that was running in a Container, used to persist data from our APIs (Customer, Product, Order and Voucher) to Amazon RDS, for the following advantages:

1. Scalability and performance: With Amazon RDS we can scale our database vertically and horizontally as needed, making it easier to manage the performance of our application in the future as our user/product base grows.

2. Data availability and durability: With Amazon RDS it is possible to configure our database to guarantee the high availability of our data, through automatic backups, replication in multiple availability zones and recover our data in cases of disaster.

3. Simplified management: Being a platform as a service (PaaS), we don't need to worry about patch and operating system updates, and simplified access via the AWS console and CLI to configure, monitor and adjust our database.

4. Security: Amazon RDS offers robust security features such as data encryption both securely and in transit, as well as integration with other AWS services for access control and compliance.

5. Costs: We decided that if our application were to grow, managing it through Amazon RDS would generate savings in the long term, rather than working within an infrastructure as a service (IaaS) environment, as it is a elastic and scalable environment.


## Database noSQL - Amazon Dynamo DB

We decided to create a table in the noSQL Dynamo DB database to work together with cognito in our authentication system, for the following advantages:

1. Scalability and performance: Dynamo DB is a highly scalable noSQL database and can handle workloads in more varied and unpredictable ways, works with a large number of users or traffic spikes and has the ability to scale automatically.

2. Simplified management: It is a service fully managed by AWS and for this reason we do not need to worry about the underlying infrastructure, AWS takes care of the configuration, monitoring and maintenance of the database.

3. Integration with Amazon Cognito: One of the reasons for our choice is the easy integration between Dynamo DB, Cognito authentication service and lambda functions. We used Cognito to manage users and authenticate them, while Dynamo DB was used to store authentication related data.


