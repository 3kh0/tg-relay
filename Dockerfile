FROM golang:alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
COPY templates/ ./templates/
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /relay .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /relay /relay
EXPOSE 8080
ENTRYPOINT ["/relay"]
