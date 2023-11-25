FROM golang:1.20

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /fresheggs

EXPOSE 8000

CMD ["/fresheggs"]
