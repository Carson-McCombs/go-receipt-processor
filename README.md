# go-receipt-processor
Written in Go 1.22.2 - hosts a local server, of which the port must be set during build. The server that accepts receipts in json format, parses and validates the information on the receipt, before returning the number of points it is worth.

Currently, the internal port must be 8080 -> "docker run -p 8081:8080"

Currently, the container is only tested on AMD64 processors.

Usuage Examples:


HTTP POST "http://0.0.0.0:8081/receipts/process" sending 

example data: 

"{
  "retailer": "Bestbuy",
  "purchaseDate": "2023-10-15",
  "purchaseTime": "15:30",
  "items": [
    {
      "shortDescription": "CD",
      "price": "10.20"
    },{
      "shortDescription": "Game",
      "price": "60.01"
    },
  ],
  "total": "70.21"
}"

if statuscode 200 - returns a json containings an id corresponding to this receipt.

HTTP GET "http://0.0.0.0:8081/receipts/{id}"

if statuscode 200 - returns number of points the corresponding receipt is worth
