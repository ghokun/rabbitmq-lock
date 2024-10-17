FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /rabbitmq-lock

FROM gcr.io/distroless/static-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /rabbitmq-lock /rabbitmq-lock

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/rabbitmq-lock"]