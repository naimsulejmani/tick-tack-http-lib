# ticktackhttp - A Simple Generic HTTP Client Library for Go

`ticktackhttp` is a lightweight, generic HTTP client library for Go that allows you to make HTTP requests with support for different HTTP methods (`GET`, `POST`, `PUT`, `PATCH`, `DELETE`) and handle responses in a generic way using Go's type parameters.

## Features

- Generic support for request bodies and responses.
- Handles all common HTTP methods (`GET`, `POST`, `PUT`, `PATCH`, `DELETE`).
- Flexible response handling for both structs and slices.
- Automatically handles JSON marshalling and unmarshalling.

## Installation

To use this library, you need to import it into your Go module. You can install it with:

```bash
go get github.com/yourusername/ticktackhttp
```

## Usage
1. Import the library
```go 
import "github.com/yourusername/ticktackhttp"
```

2. Use the library for HTTP requests
Below are examples of how to use ticktackhttp for GET, POST, and other requests.
**Example of GET Request**
```go
package main

import (
    "fmt"
    "github.com/yourusername/ticktackhttp"
)

type Todo struct {
    UserID    int    `json:"userId"`
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func main() {
    url := "https://jsonplaceholder.typicode.com/todos/1"
    headers := map[string]string{}

    // Call GET request expecting a Todo object
    var todo Todo
    resp, err := ticktackhttp.GenericRequest[any, Todo]("GET", url, headers, nil)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Fetched TODO: %+v\n", resp)
}
```


**Example of POST Request**
```go
package main

import (
    "fmt"
    "github.com/yourusername/ticktackhttp"
)

type Todo struct {
    UserID    int    `json:"userId"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func main() {
    url := "https://jsonplaceholder.typicode.com/todos"
    headers := map[string]string{
        "Content-Type": "application/json",
    }

    // Request body for the POST
    reqBody := &Todo{
        UserID:    1,
        Title:     "Learn Golang",
        Completed: false,
    }

    // Call POST request expecting a Todo object
    var respBody Todo
    resp, err := ticktackhttp.GenericRequest("POST", url, headers, reqBody)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Created TODO: %+v\n", resp)
}

```


**Example of DELETE Request**
```go
package main

import (
    "fmt"
    "github.com/yourusername/ticktackhttp"
)

func main() {
    url := "https://jsonplaceholder.typicode.com/todos/1"
    headers := map[string]string{}

    // Call DELETE request
    var respBody string
    resp, err := ticktackhttp.GenericRequest[any, string]("DELETE", url, headers, nil)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Deleted TODO, Response: %s\n", resp)
}

```


## API Documentation
#### Parameters
* method (string): The HTTP method (GET, POST, PUT, PATCH, DELETE).
url (string): The URL to make the request to.
* headers (map[string]string): The HTTP headers to include with the request.
* body (T): The request body (use nil for GET and DELETE requests).
Returns
R: The parsed response body.
* error: Any error encountered during the request or unmarshalling process.

#### Handling Responses
The library automatically handles responses by unmarshalling JSON responses into the provided generic type R. If the expected response type is a string (e.g., for a DELETE request that returns no body), the raw response body is returned as a string.

License
This project is licensed under the MIT License - see the LICENSE file for details.
