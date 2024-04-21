
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


## About the Project
Written in Go 1.22.2. This was written as a demo project to show an example of my ability to work with Go and Docker and is created based on Fetch's Receipt Proccessor Challenge. I will attempt to ommit some information about the challenge to try and avoid jeopardising the challenge itself. 

During development, I wanted to use as few dependencies as possible. This is why you see me with my own Date and Time package instead of using Go's standard library Time package. This does lead to a little more bloat as it required me to add more error checks, data validation, and testing. 

The project functions by hosting a local server. This server accepts receipts in json format and returns a json containing an id that corresponds to the provided receipt. This id can then be used to get the number of points - of which are determined by the rules given on Fetch's repository page.  that parses and validates the information on the receipt, returning an id. This id can be used to get the number of points the corresponding receipt is worth.

### Built With

* ![Golang 1.20.2](https://img.shields.io/badge/Golang-lightblue?style=for-the-badge&logo=go&logoColor=white&color=%2300ADD8&link=https%3A%2F%2Fgo.dev%2F)
* ![Docker](https://img.shields.io/badge/docker-lightblue?style=for-the-badge&logo=docker&logoColor=white&color=%232496ED&link=https%3A%2F%2Fwww.docker.com%2F)

## Getting Started

### Prerequisites

* [Docker Desktop](https://docs.docker.com/desktop/install/windows-install/)

### Installation

#### Navigate to where you would like to store the directory

*From Command Line:*

```
git clone https://github.com/Carson-McCombs/go-receipt-processor
docker build --tag go-receipt-processor go-receipt-processor
```

This clones the github and builds the Docker image.

### Usage

#### To Create and Start Docker Container

* When you are calling the "docker run" command, you can modify what port you are publishing to. The port which you interact with the server through is set in the format "-p <host_port>:<container_port>", but do note that the container_port *must* be set to 8080 to work. This is the port that the docker container, which is holds the server, relays the HTTP requests to and from.

*From Command Line:*

```
docker run -d -p 80:8080 go-receipt-processor
```

This creates a Docker Container from the Docker Image and runs it.

#### To Stop Running Docker Container

*From Command Line:*

```
docker stop go-receipt-processor
```

See [link](https://docs.docker.com/guides/walkthroughs/run-a-container/) for more information about running a Docker container.


#### Sending a Receipt to be Processed:

In this case, I am saving the json of the Receipt as a txt file named "textcase.txt".

<details>
  <summary>textcase.txt</summary>
  
  ```
  {
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
  }
  ```

</details>

*From Command Line:*
```
curl -X POST -H "Content-Type: application/json" -d @testcase.txt http://localhost:80/receipts/process
```
**Example output:**

```
{
  "id": "9d49ee51-1743-467a-8445-bc75cabe0b44"
}
```

#### Requesting the points that the Receipt is worth

where {id} is the value of the json id returned by the previous curl command

*From Command Line:*

```
curl -http://localhost:80/receipts/{id}
```

**Example output:**

```
{
  "points": "28"
}
```

## Contact

* Carson McCombs - carson.mccombs.work@gmail.com
* Project Link: https://github.com/Carson-McCombs/go-receipt-processor

## Acknowledgments

* [Best-README-Template](https://github.com/othneildrew/Best-README-Template)
* [Smallest Golang Docker Image](https://klotzandrew.com/blog/smallest-golang-docker-image)
