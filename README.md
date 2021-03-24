<p align="center"><a href="https://hoopspots.io" target="_blank"><img src="https://hoopspots.s3.amazonaws.com/hoopspots-orange-text.png" alt="logo" width="300"></a></p>
Welcome to the GraphQL API for HoopSpots. 

## Setup Docker
1. Install [Docker](https://docs.docker.com/docker-for-mac/install/)
2. Build Docker containers `docker-compose build`
3. Run `docker-compose up`

### GraphQL Playground
Go to http://localhost:5000/

## About
- Install Docker and setup local
- Mailhog Access
- Database Access
- How to use the cli
- Seed your database
- Create a test command
- Run the tests
- Lint your code
- Git-cz before you commit

<p>Welcome to the backend codebase for HoopSpots: Pick-Up Basketball App</p>

## Getting Started
Make sure you have Golang 1.16
1. Run `go get`
2. Run `go run github.com/99designs/gqlgen generate`
3. Run `go run server.go`


### Timestamps
We use Carbon

### Database Seeding
How to run the database seeder
How do I add a seeder

### Go Report Card
https://github.com/gojp/goreportcard

### Concerns
- How do I add tests?
- Add a CLI tool similar to Laravel artisan command
- Add Queued Jobs (SQS, Redis, RabbitMQ)