## What is the Archive Demo
This is a small proof-of-concept of a data archive and query solution. 

## High level architecture
![architecture](/documentation/arch.png?raw=true "High-level architecture")


## Install and run

    docker-compose up --build

For now you need to create a new user account by making a direct request

If using postman

make a POST call to localhost:8101/login

    {
	    "username":"admin",
	    "password":"test",
	    "role":"admin"
    }

Make post call to localhost:8101/admin/create-account

    {
	    "username":"my-account",
	    "password":"my-acc-pwd",
	    "my_bucket":"my-bucket"
    }

Note that username and password need to be longer than 8 characters

Go to http://localhost:8181
Login with my-account my-acc-pwd

Query not implemented in UI but you can query a file with postman
Login as my-account, upload a csv or gzip file

Call localhost:8101/user/query with
    {
    "query":"SELECT * FROM name.of.file.gz LIMIT 10",
    "dataset":"name.of.file.gz",
    "record_delimiter":"\n",
    "field_delimiter":",",
    "output": "json"
    }

All your files are listed in the file explorer. You can switch from one to another by clicking a file in the list.
