# getir-app

## Prerequisite

1. Golang version 1.14 or above

## Instructions to run getir-app in local

1. Clone the repository to your local machine.

    ```
     $ git clone git@github.com:sakshi29/getir-test.git
    ```

2. To run all test cases in the repo, use command

    ```
     $ go test ./...
    ```

3. To run application, use command inside `cmd/getir-app`

    ```
    $ go build
    $ ./getir-app
    
    OR

    $ go run app.go    
    ```

    It will start serving the application on http://localhost:9000

    To check your application is up and running.

    ```
    curl http://localhost:9000/healthcheck
    ```

## Heroku

This application is also being served on Heroku under the domain.

https://boiling-bastion-20752.herokuapp.com

## Sample Requests

1. Getting documents from mongoDB

    ```
    curl -X POST 'https://boiling-bastion-20752.herokuapp.com/documents' \         
    --header 'Content-Type: application/json' \
    --data-raw '{
        "startDate": "2015-01-26",
        "endDate": "2018-02-02",
        "minCount": 2000,
        "maxCount": 2788
    }'
    ```

2. Add a record to In-Memory Database

    ```
    curl -X POST 'https://boiling-bastion-20752.herokuapp.com/in-memory' \
    --header 'Content-Type: text/plain' \
    --data-raw '{
        "key":"active-tabs",
        "value":"getir"
    }'
    ```

3. Fetch a record from In-Memory Database

    ```
    curl -X GET 'https://boiling-bastion-20752.herokuapp.com/in-memory?key=active-tabs'
    ```
