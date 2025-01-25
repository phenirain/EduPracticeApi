FROM golang:1.23-apline AS build
LABEL authors="phenirain"


WORKDIR /build

COPY go.mod go.sum /
RUN go mod download

COPY ./ ./

RUN go build -o api ./cmd/server/main.go
RUN go build -o api-migrator ./cmd/migrator/main.go

FROM golang:1.23-apline AS run

WORKDIR /app
COPY --from=build /build/.env ./.env
COPY --from=build /build/api ./api
COPY --from=build /build/api-migrator ./api-migrator
COPY --from=build /build/migrations ./migrations