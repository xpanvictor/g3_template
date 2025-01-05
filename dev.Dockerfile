FROM golang:1.23-alpine

WORKDIR /app

# Install development tools
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

# Use air for hot reload
CMD ["air"]