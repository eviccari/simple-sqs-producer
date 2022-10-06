FROM golang:1.19-alpine as stage

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go get github.com/swaggo/swag/cmd/swag && \
    swag init -g handlers/*.go && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server ./cmd/main.go

FROM scratch as runner
COPY --from=stage /app/server server
COPY --from=stage /app/docs docs
COPY --from=stage /app/cmd/.env .env

EXPOSE 8080:8080

ENTRYPOINT ["./server"]

