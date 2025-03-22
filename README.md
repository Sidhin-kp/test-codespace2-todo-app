TO-DO App

A simple ToDo app built with HTML, JavaScript, and Go as the web server, using PostgreSQL as the database.
The app is containerized using Docker and orchestrated with Docker Compose for deployment

Commands

inside backend/

go mod init todoapp

go mod tidy


then in main directory


docker-compose up --build

------------------------------------------------------------------------------
Then frontend will be listen on localhost:3000
and go in localhost:8008/tasks/
