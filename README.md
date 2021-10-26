<!-- @format -->

# Golang Gin-Gonic Authentication (Under Development)

A Simple Token-Based User Authentication Service using JWT in Golang Gin-Gonic and MySql/Postgres.

# Installation

With this you can quickly craft a token-based user authentication system using JWT and continue your project implementation. This comes with User Registration, Email Verification, Login, Password Reset, Logout.

# Installation

### Step 1

Clone or download this repository to your machine:

- Clone the repo: `git clone https://github.com/bytesfield/golang-gin-auth.git`
- [Download from Github](https://github.com/bytesfield/golang-gin-auth/archive/refs/heads/main.zip).

### Step 2

`go mod init` to install all packages

Create your database, rename `.env.example` to `.env` then, change `DB_DATABASE` value to your database and its credentials. Change `APP_URL` to your preferred url default is `http://localhost:8000` and set up your email service, mailgun has already been set up by default just input `MAILGUN_SECRET` value `MAIL_MAILER` credentials.

On your terminal run `php artisan migrate` to create necessary tables and `php artisan jwt:secret` to create `JWT_SECRET` in your `.env`.

### Step 3

Start your development server : `go run main.go` this serves the application to default `localhost:8080`

### Step 4

Open Postman and run the Api endpoints. Documentation can be accessed below

# Documentation

The API documentation is hosted on [Postman Doc](https://documenter.getpostman.com/view/10912779/TzRUBnVB)

# Contribution

Want to suggest some improvement on the codes? Make a pull request to the `dev` branch and it will be reviewed and possibly merged.
Find me on
<a href="https://twitter.com/SaintAbrahams/">Twitter.</a>
<a href="https://www.linkedin.com/in/abraham-udele-246003130/">Linkedin.</a>

# License

Source codes is license under the MIT license.
