FROM golang:1.23-alpine AS build
LABEL authors="phenirain"

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /build/api-migrator ./cmd/migrator/main.go
RUN go build -o /build/api ./cmd/server/main.go

FROM golang:1.23-alpine AS run

WORKDIR /app
COPY --from=build /build/.env ./.env
COPY --from=build /build/api ./api
COPY --from=build /build/api-migrator ./api-migrator
COPY --from=build /build/migrations ./migrations

RUN chmod +x /app/api-migrator
RUN chmod +x /app/api
