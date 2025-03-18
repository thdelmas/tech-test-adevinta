# Adevinta Technical Challenge

## Intro

### What's FizzBuzz

> Fizz buzz is a group word game for children to teach them about division.
>
> Players take turns to count incrementally, replacing any number divisible by three with the word "fizz", and any number divisible by five with the word "buzz", and any number divisible by both three and five with the word "fizzbuzz".
> 
> Players generally sit in a circle. The player designated to go first says the number "one", and the players then count upwards in turn.
>
> However, any number divisible by three is replaced by the word fizz and any number divisible by five is replaced by the word buzz.
>
> Numbers divisible by both three and five (i.e. divisible by fifteen) become fizz buzz.
> 
> For example, a typical round of fizz buzz would start as follows:
>
> 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, Fizz, Buzz, 26, Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz, ...


## Specifications
```
Write a simple fizz-buzz REST server.

"The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"".
The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...""."

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request
```

## Installation and Setup

### Requirements
- GNU Make
- Docker
- cURL or Postman for testing API endpoints


### Cloning the Repository
To get the source code, clone the repository:

```sh
git clone git@github.com:thdelmas/tech-test-adevinta.git
cd tech-test-adevinta
```

## Run & test
> The software comes with a Makefile and dockerfile


## Make && Deploy

The makefile allow you to perform the following actions:
1. build, run `make`
2. test, run `make test`
3. run, run `make run`
4. deploy, `make deploy`

### Building the Docker Image
> The compilation is made in the makefile and not in the dockerfile
>
> The container only contains the binary and not the source code

To build the Docker image, run:
```sh
 make &&
docker build --no-cache -t fizzbuzz-api .
```
or just run `make docker_build`

### Running the API
To start the API on port 8080:
```sh
 docker run -p 8080:8080 fizzbuzz-api
```

instead you can run `make deploy` which will compile and build the image first (no cache enabled)

## Running Tests
In order to test the code, run `make test`, alternatively you can use curl or the postman collection to send manual request once the server is running

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
