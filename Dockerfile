#building the binary file
FROM golang AS builder

LABEL maintainer="Justin John"

WORKDIR /app

COPY go.mod .
COPY go.sum .

# Add the go mod download to pull in any dependencies
RUN go mod download

# Copy everything from this project into the filesystem of the container.
COPY . .

# Download all the dependencies
#RUN go get -d -v ./...

# Install the package
#RUN go install -v ./...

# Compile the binary exe for our app.
RUN go build -o main main.go

#installing goose migration
#RUN export GO111MODULE=on &&\
RUN go get github.com/pressly/goose/cmd/goose@latest
#RUN curl -L https://github.com/pressly/goose

#RUN mkdir /migrations
#WORKDIR /migrations

#goose version
#RUN goose status

#final run stage 
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ .
COPY . .
RUN ls -l


EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
