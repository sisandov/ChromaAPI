# ChromaAPI

## Setup and commands

### API Authentication

Make a .env file following the .env_sample format.

### Hot swap

Install air:

`go install github.com/cosmtrek/air@latest`

### Adjust dependencies:

`go mod tidy`

### Run the project:

`air`

## CI/CD

| Command              | Description                                              |
| -------------------- | -------------------------------------------------------- |
| `go test -v`         | Run all test on `_test.go` files recursively             |
| `go test -v ./utils` | Run all test on `_test.go` files recursively in a folder |
| `golangci-lint run`  | Run the linter for golang                                |

#### For testing be sure that you have installed the linter

You can install this with:

`curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2`

#### Husky

Install husky cli tool with

`go install github.com/automation-co/husky@latest`

### Docker

Run the project:

`docker-compose up`
