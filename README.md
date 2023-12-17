# Attendance Apps

## Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/dhimweray222/Test-BE-UNINET.git
    cd Test-BE-UNINET
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```
3. Change Config in .env file
    ```bash
    # Server Settings:
    SERVER_URI=localhost
    SERVER_PORT=8080
    DEFAULT_LIMIT=10
    
    # Database Setting
    DB_HOST=host
    DB_PORT=port
    DB_NAME=db_name
    DB_USERNAME=postgres
    DB_PASSWORD=db_password
    DB_POOL_MIN=10
    DB_POOL_MAX=100
    DB_TIMEOUT=10
    DB_MAX_IDLE_TIME_SECOND=60
    
    # jwt
    SECRET_KEY=AaBbCcJKLMadhanirAAnNoOPqrstu23VWXYZ
    SECRET_KEY_REFRESH=AaBbCcJKLMadhaniRAaPqrstu23VWXYZ
    SESSION_LOGIN=24
    SESSION_REFRESH_TOKEN=720
    
    # Location
    TOKEN_MAPS=37c54fbcd8fb4f
    
    ```
4. Run the application:

    ```bash
    go run main.go
    ```

5. Open your web browser and visit [http://localhost:3000](http://localhost:8080) to see the app in action.

## Features
1. Create User
    ```
    Endpoint: POST /users
    Body:
    {
      "id": "123456789",
      "name": "dhim Doe",
      "email": "122@example.com",
      "password": "123123123"
    }
    ```
2. Login User
    ```
    Endpoint: POST /users/login
    Body:
    {
        "email":"123@example.com",
        "password":"123123123"
    }
    ```
3. Detail User
    ```
    Endpoint: POST /users/:id
        Params:
        {
            "id":"example_uuid_user",
        }
    ```

4. Check In
    ```
    Endpoint: POST /attendances
        Cookies:
        {
            "token":"example_token_from_login",
        }
    ```
5. Check Out 
    ```
    Endpoint: POST /attendances/:id
        Params:
        {
            "id":"example_uuid_attendance",
        }
    ```





