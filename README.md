
## Run ##
    sh bootstrap.sh

## FindList ##
    curl "http://localhost:3000/servers?room_id=1"

    [
        {
            "id":3,
            "room_id":1,
            "address":"tcp://127.0.0.1:8080"
        },
        {
            "id":4,
            "room_id":1,
            "address":"tcp://127.0.0.1:8081"
        },
        {
            "id":5,
            "room_id":1,
            "address":"tcp://127.0.0.1:8082"
        }
    ]

## GET ##
    curl "http://localhost:3000/servers/3"

    {
        "id":3,
        "room_id":1,
        "address":"tcp://127.0.0.1:8080"
    }
