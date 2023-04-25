# Backend Template
This is a backend template project written in Golang, 
which utilizes MongoDB as a database, Redis for caching, Minio for object storage, 
SwagGo for a swagger document and Docker for deployment.

## Table of Contents
- [Getting Started](#getting-started) 
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Deployment](#deployment)
- [Testing](#testing)
- [Linting](#linting)
- [Generate Documents](#generate-documents)
- [CI/CD](#cicd)
- [Built With](#built-with)
- [Contributing](#contributing)
- [License](#license)
- [TODO](#todo)

# ⚠️ ATTENTION ⚠️
# BEFORE DEVELOPING MAKE SURE CHANGE AND RENAME REFERENCE PACKAGES USED IN PACKAGES.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites
To run this project, you will need to have the following installed on your system:

Golang\
MongoDB\
Redis\
Minio\
Docker

## Installation
To get started, first clone the repository:

`git clone https://github.com/dozheiny/backend-template.git`


`cd backend-template`

Then, run the following command to install the required dependencies:

`go mod download`

After that, you can set the required environment variables by creating a .env file in the root of the project and adding the necessary variables. You can use the .env.example file as a template.

# Usage
To run the project locally, you can use the following command:


`go run ./... .`
This will start the server and make it accessible at http://localhost:8080.

# Deployment
To deploy the project using Docker, you can use the following commands:

`docker build -t REPOSITORY-NAME -f build/Dockerfile .`

`docker run -p 8080:8080 -it REPOSITORY-NAME`

This will build the Docker image and start the container, making the server accessible at http://localhost:8080.

# Testing
To run the tests for the project, use the following command:

`go test ./...`

This will run all tests in the project.

# Linting
To lint the project code, use the following command:

`golangci-lint run`

This will run the linter on the project.

# Generate Documents
To generate a document with swagger, use the following command:

`swag init -g 'cmd/api/main.go' --output docs/swagger --parseDependency`

For more usage follow the official document.

# CI/CD
For using CI/CD pipelines, just set environments into github environment variables repository.

Set this `MONGODB` and `DB_NAME` environment variables into your github environment variable.

# Built With
- Golang - Programming language used
- MongoDB - Database used
- Redis - Caching service used
- Minio - Object storage used
- Docker - Deployment tool used
- Swagger - Document tool used

# Contributing
Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests.

# License
This project is licensed under the MIT License; see the LICENSE file for details.

# TODO
- [ ] Add reference links to README.md file.
- [ ] Add `docker-compose.yml` file for the fastest deployment. 
- [ ] Add more details to `README.md` file packages. 