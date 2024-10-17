FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o api .

EXPOSE 8000

CMD ["./api"]
