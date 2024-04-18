# ChromaAPI

## Setup and commands

First install air:

`go install github.com/cosmtrek/air@latest`

### Adjust dependencies:

`go mod tidy`

### Run the project:

`air`

### Testing

The files that are tested ends with `_test.go`

`go test` (This maybe dont find the files)

If you want to test a folder recursively use

`go test -v ./handlers`

### Docker

Run the project:

```docker-compose up```

### API Authentication

Make a .env file following the .env_sample format.
