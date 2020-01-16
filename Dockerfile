FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -mod=vendor -v -o ./anagram-service cmd/*.go
EXPOSE 8080
CMD ["./anagram-service"]