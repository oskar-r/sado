## What is the Archive Demo
This is a small proof-of-concept of a data archive and query solution. 

## High level architecture
![architecture](/documentation/arch.png?raw=true "High-level architecture")


## Install and run

    docker-compose up --build

Default the web-client will runt on port 8181 and the backend service on port 8101


For now you need to create a new user account by making a direct request

If using postman

To get credentials needed for accoutn creation

### POST /login
```javascript
{
    "username":"admin",
    "password":"test",
    "role":"admin"
}
```

Create user account

### POST /admin/create-account
```javascript
{
    "username":"my-account",
    "password":"my-acc-pwd",
    "my_bucket":"my-bucket"
}
```

Note that username and password need to be longer than 8 characters. 

Go to http://localhost:8181
Login with my-account my-acc-pwd

N.B Query not implemented in UI but you can query a file with postman
Login as my-account, upload a csv or gzip file

You can then call the query api to get a sample. Remeber to use right credentials in the call

POST /user/query 
```javascript
{
    "query":"SELECT * FROM name.of.file.gz LIMIT 10",
    "dataset":"name.of.file.gz",
    "record_delimiter":"\n",
    "field_delimiter":",",
    "output": "json"
}
```

All your files are listed in the file explorer. You can switch from one to another by clicking a file in the list.
