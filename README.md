# go-receipt-processor
Written in Go 1.22.2 - hosts a local server, of which the port must be set during build. The server that accepts receipts in json format, parses and validates the information on the receipt, before returning the number of points it is worth.

# Installation and Setup


## Navigate to where you would like to store the directory

When you are calling the "docker run" command, you can modify what port you are publishing to. The port which you interact with the server through is set in the format "-p <host_port>:<container_port>", but do note that the container_port *must* be set to 8080 to work. This is the port that the docker container, which is holds the server, relays the HTTP requests to and from.
```
git clone https://github.com/Carson-McCombs/go-receipt-processor
docker build --tag go-receipt-processor go-receipt-processor
docker run -d -p 80:8080 go-receipt-processor
```


# Usuage Examples:

curl -X POST -H "Content-Type: application/json" -d '{"retailer": "Bestbuy","purchaseDate": "2023-10-15","purchaseTime": "15:30","items": [ {    "shortDescription": "CD",   "price": "10.20"   "price": "60.01" },], "total": "70.21" }' http://0.0.0.0:80/receipts/process

HTTP POST "http://0.0.0.0:80/receipts/process" sending 

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
