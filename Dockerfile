FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum main.go ./
RUN go mod download

COPY ./src ./src
COPY ./docs ./docs

RUN CGO_ENABLED=0 GOOS=linux GIN_MODE=release go build -o /go-mongo-starter

EXPOSE 3000

CMD ["/go-mongo-starter"]