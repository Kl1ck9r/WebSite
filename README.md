# Web Site to store personal notes



# Usage

Endpoint | Usage
------------ | -------------
/page/signup | Use this route to sign up a new user
/page/login   | Authenticate user to access
/page/reset/password | Password recovery in case you forget it
/page/main           | The Home page
/page/error        | Page in case of an error
/page/show/notes    |  Page displaying user record saving
/page/delete/note  | Page to delete note 


# Requirements

Go - from v1.19

PostgreSQL - from v15


# Setup
- Before setup, you must have PostgreSQL on your machine. Then clone this repository

    > https://github.com/Kl1ck9r/WebSite.git

- cd into the project directory

   > cd Web-Site-Go

- Set environment variables

  - A sample is presented in the repository.
  
  > cp .env

- Run application

  > go run main.go
