FROM golang:1.17-alpine AS build
WORKDIR /build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o /app

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /app /app
EXPOSE 80
ENTRYPOINT ["/app"]