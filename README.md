# REST API transaction app

This app deploys on port 8080. Binary in repository was generated on MacOS. If using a different OS, another binary should be rebuilt using `go build`. Commands below must be executed in the root of the repository.

## Build

    go build

## Run the app

    ./test-api

## Run the tests

    go test ./...

# REST API

The REST API to the transaction app is described below.

## Create a new Transaction

### Request

`POST /transactions`

    curl -i -H 'Accept: application/json' -d '{"type":"credit","amount":30}' http://localhost:8080/transactions

### Response

    HTTP/1.1 201 Created
    Content-Type: application/json; charset=utf-8
    Date: Sun, 14 Feb 2021 21:41:36 GMT
    Content-Length: 124

    {
        "id": "11811d0a-8669-4d0b-ac55-51e875caabdf",
        "type": "credit",
        "amount": 30,
        "effectiveDate": "2021-02-14T18:30:45.435482-03:00"
    }

## Get Transaction by ID

### Request

`GET /transactions/:id`

    curl -i -H 'Accept: application/json' http://localhost:8080/transactions/11811d0a-8669-4d0b-ac55-51e875caabdf

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Sun, 14 Feb 2021 21:42:36 GMT
    Content-Length: 124

    {
        "id": "11811d0a-8669-4d0b-ac55-51e875caabdf",
        "type": "credit",
        "amount": 30,
        "effectiveDate": "2021-02-14T18:30:45.435482-03:00"
    }

## Get transaction history

### Request

`GET /transactions`

    curl -i -H 'Accept: application/json' http://localhost:8080/transactions

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Sun, 14 Feb 2021 21:44:36 GMT
    Content-Length: 126

    [
        {
            "id": "11811d0a-8669-4d0b-ac55-51e875caabdf",
            "type": "credit",
            "amount": 30,
            "effectiveDate": "2021-02-14T18:30:45.435482-03:00"
        }
    ]

## Get account balance 

### Request

`GET /balance`

    curl -i -H 'Accept: application/json' http://localhost:8080/balance

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Sun, 14 Feb 2021 21:46:36 GMT
    Content-Length: 2

    30
