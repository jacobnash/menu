FROM golang:1.23 AS Builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o menu .

FROM alpine:3.20.0

WORKDIR /app

COPY --from=Builder /app/menu/ .

EXPOSE 8080 

CMD ["./menu"]
