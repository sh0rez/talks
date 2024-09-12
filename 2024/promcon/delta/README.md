# PromCon24: Sending delta to Prometheus

This was given as a lightning talk during PromCon 2024 Berlin

### Running

```sh
$ docker-compose up -d
$ OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318 go run ./deltagen
```
