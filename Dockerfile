#building the binary file
FROM golang AS builder

LABEL maintainer="Justin John"

WORKDIR /app

COPY go.mod go.sum ./

# Add the go mod download to pull in any dependencies
RUN go mod download && go mod verify

# Copy everything from this project into the filesystem of the container.
COPY . .

# Compile the binary exe for our app.
RUN go build -o main main.go

#installing goose migration
#RUN go get github.com/pressly/goose/cmd/goose@latest

#final run stage 
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ .
COPY . .
COPY .env .
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]