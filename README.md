# NAWNAW-API

This project is a RESTful API built using the Echo framework in Go. It provides authentication and user management functionalities.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed Go (1.20 or later).
- You have a running MySQL database.
- You have installed `viper` for configuration management.
- You have installed `gorm` for ORM.
- You have installed `go-migrate` for database migrations.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/nawafilhusnul/NAWNAW-API.git
    cd NAWNAW-API
    ```

2. Install the dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your MySQL database and update the `config.json` file with your database credentials.

4. Run the database migrations:
    ```sh
    make migrate-up
    ```

5. Run the server:
    ```sh
    go run main.go
    ```

## Configuration

Create a `config.json` file in the root directory of the project with the following structure:
```json
{
  "app": {
    "name": "NAWNAW-API",
    "host": "localhost",
    "port": "8080"
  },
  "database": {
    "driver": "",
    "host": "",
    "port": "",
    "username": "",
    "password": "",
    "database": "",
    "max_idle_conns": 0,
    "max_open_conns": 0,
    "conn_max_lifetime": 0
  },
  "migration": {
    "dir": "db/migrations"
  },
  "secret": {
    "jwt": "",
    "encrypt": "",
    "access_expired": 0,
    "refresh_expired": 0
  }
}
```


