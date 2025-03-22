TO-DO App

Simple todo app using HTML, JS, and Go as a web server and Postgres DB.
Containerized it using Docker and deployed it using docker-compose.

Commands

inside backend/

go mod init todoapp
go mod tidy


docker-compose up --build

------------------------------------------------------------------------------
Then fromtend will listen on localhost:3000
and go in localhost:8008/tasks/
