FROM golang:1.19-alpine as stage

WORKDIR /app

COPY . .

RUN go mod tidy && \
   go get github.com/swaggo/swag/cmd/swag && \
   go install github.com/swaggo/swag/cmd/swag@latest && \
   swag init -g handlers/*.go && \
   CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server ./cmd/main.go

#RUN go mod tidy
#RUN go get github.com/swaggo/swag/cmd/swag
#RUN go install github.com/swaggo/swag/cmd/swag@latest
#RUN swag init -g handlers/*.go
#RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server ./cmd/main.go

FROM scratch as runner
COPY --from=stage /app/server server
COPY --from=stage /app/docs docs
COPY --from=stage /app/cmd/.env .env

#EXPOSE 8080:8080

ENTRYPOINT ["./server"]

