## What is the Archive Demo
This is a small proof-of-concept of a data archive and query solution. 

## High level architecture
![architecture](/documentation/arch.png?raw=true "High-level architecture")

## Credentials and variables
There is an example .env file with the variables that needs to be set.

Note that ADMIN_PWD and NATS_PWD are bcrypt encoded
Generate these with e.g. 
```shell
htpasswd -n -B -C 11 admin
````

Replace the $2y in the begining with $2a. Nats password also need to be provided in clear text for minio and the backend


## Install and run

    docker-compose up --build

## Try
You can now access the front end client trough
```shell
front.app.localhost
```
Login with admin and the ADMIN_PWD
You can now switch user to admin and then add a user account 
