# Resource Leak in Go HTTP Request Handler

This repository demonstrates a common resource leak in Go's HTTP request handlers.  The issue arises from improper handling of the `http.Request.Body`'s closure.  If an error occurs during request processing, the `defer r.Body.Close()` statement might not execute, leaving the request body open and potentially leading to resource exhaustion.

The `bug.go` file shows the problematic code.  The `bugSolution.go` file presents a corrected version.

## How to reproduce

1. Clone this repository.
2. Run the `bug.go` file (using a tool like `go run`).
3. Observe that the connection may be left open
4. Run the `bugSolution.go` file and observe that connection is always closed.

## Solution

The solution involves ensuring that `r.Body.Close()` is always called, regardless of whether an error occurs. This can be achieved by using a `defer` statement.