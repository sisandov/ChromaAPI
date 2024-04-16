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

Build the image:

```docker build -t chroma-api .```

Run the image:

```docker run --rm -p 3333:3333 chroma-api```

Or run in detached mode:

```docker run -d --rm -p 3333:3333 chroma-api```

