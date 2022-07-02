## Manu CMS Backend (WIP)

# This is the backend for the Manu CMS, a simplistic CMS using golang and mongodb

# Getting started:

Create a .env file.
The .env file should contain 5 fields, 
`DB_USER,
DB_PASS,
DB_LINK (in the format cluster0@asdf.mongodb.net),
SECRET_KEY
FRONTEND_LINK (ex http://localhost:3000)`

Then, use `go run main.go` to run the app.

The codebase is kind of a mess, in my defense it is my first project with Go.
