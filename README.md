# tg-relay

dead simple telegram message relay for webhooks and other services. useful for notifications and alerts from services that don't have native telegram support.

## usage

deployment can be done via coolify with the provided Dockerfile, fill out the env vars and deploy. you can also run it locally with `go run main.go`.

```sh
curl -X POST https://example.com -d '{"message":"fart"}'
```
