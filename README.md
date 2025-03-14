# tech-test-adevinta

## Overview
The FizzBuzz API is a simple web service that generates a custom FizzBuzz sequence based on user-defined parameters. It is built using Golang and runs inside a Docker container.

## Requirements
- Docker installed on your system
- cURL for testing API endpoints

## Installation and Setup

### Cloning the Repository
To get the source code, clone the repository:
```sh
git clone git@github.com:thdelmas/tech-test-adevinta.git
cd tech-test-adevinta
```

### Building the Docker Image
To build the Docker image, run:
```sh
 docker build --no-cache -t fizzbuzz-api .
```

### Running the API
To start the API on port 8080:
```sh
 docker run -p 8080:8080 fizzbuzz-api
```

## Running Tests
The application includes tests that generate a log file.

If the tests fails so will the build of the docker image

When the build is successful you can inspect the logs with:
```sh
 docker run --rm fizzbuzz-api cat test-results.log
```

## API Endpoints

### 1. Generate FizzBuzz Sequence
**Endpoint:**
```
GET /api/fizzbuzz
```

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| int1      | int  | First divisor |
| int2      | int  | Second divisor |
| limit     | int  | Upper limit for the sequence |
| str1      | str  | String to replace multiples of `int1` |
| str2      | str  | String to replace multiples of `int2` |

**Example Request:**
```sh
curl 'http://localhost:8080/api/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz'
```

**Example Response:**
```json
["1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"]
```

### 2. Retrieve API Statistics
**Endpoint:**
```
GET /api/stats
```

**Example Request:**
```sh
curl 'http://localhost:8080/api/stats'
```

**Example Response:**
```json
{
  "most_frequent_request": {
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "fizz",
    "str2": "buzz",
    "hits": 10
  }
}
```

## Additional Notes
- Ensure that port `8080` is available before running the container.
- The `/api/stats` endpoint helps track the most frequently requested FizzBuzz parameters.
- The application follows a stateless architecture, meaning each request is processed independently.
- A postman collection is present (`postman_fizzbuzz.json`) and can be imported for manual request using postman instead of curl 
