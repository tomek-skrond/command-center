FROM golang:1.22

WORKDIR /usr/src/commandcenter

COPY go.mod ./

COPY . .

RUN go mod download && go mod verify

RUN go build -o commandcenter ./...

CMD ["./commandcenter"]

EXPOSE 2137