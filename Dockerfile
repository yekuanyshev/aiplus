FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o /service

FROM gcr.io/distroless/base-debian10

COPY --from=build /service /service

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT ["./service"]