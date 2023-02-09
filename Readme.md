# OHLC Data API

## Introduction

A RESTful api to upload and process OHLC data into database.

### Features

- Loads OHLC data into data
- Search Query with pagination to retrieve data

## Tech Stack

- Go 1.19 (Gin + Gorm)
- PostgreSQL
- Docker

## Data Model Design

| column     | Description                                    | Data type |
| ---------- | ---------------------------------------------- | --------- |
| unix       | unix                                           | integer   |
| symbol     | symbol                                         | string    |
| open       | open                                           | string    |
| close      | close                                          | string    |
| timestamps | default time stamps of a database model        | date      |
| low        | low                                            | string    |
| version    | keep track of record changes in the database   | integer   |
| id         | unique identifier and primary key of the table | uuid      |


## Setup Locally

- Make a copy of `env.example.com` to `env`
- Install [Go](https://go.dev/doc/install)
- Install [PostgreSQL](https://www.postgresql.org/download/)
- Run sql statement in this [file](api/static/migrations/init_db.sql). *You might need change values to your database credentials and not run this. Don't run this if your postgres server is a fresh install.* 
- Run `go test -v ./...` to run tests
- Run `go run main.go` to start up application

## Docker Setup

- Make a copy of `env.example.com` to `env`
- Install [Docker](https://docs.docker.com/engine/install/)
- Run `make start` in your terminal
- To stop the containers run `make stop`

## Access

- Access health check endpoint on [localhost](http://localhost:3000/api/healthcheck)
- Access swagger document on [Swagger](http://localhost:3000/api/swagger/index.html) in a browser of your choice.

## Improvement

- Notification on when file is done uploading
- Log table in the database to log information about the uploaded file. (eg: file path in the storage directory)
- Add more unit and integration test.
- Feature to avoid duplication.
- Indexing records to enable faster query
- Add `pattern matching` query to search api symbol field.

