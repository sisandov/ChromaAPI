FROM golang:1.22-bookworm

RUN mkdir /app
WORKDIR /app
ADD . /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/cosmtrek/air@latest


EXPOSE 3333

ENTRYPOINT air
