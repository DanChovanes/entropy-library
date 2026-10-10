# entropy-library

A reusable Go library for generating random data, with an optional HTTP API.
The initial implementation uses Go's `crypto/rand`; specialized entropy
collection can be added behind this library API as the project evolves.

## Use as a library

Import the module in another Go project and call `entropy.Bytes`:

```go
data, err := entropy.Bytes(32)
if err != nil {
    // Handle the entropy source error.
}
```

The package rejects negative lengths.

## Run the HTTP server

```sh
go run ./cmd/entropy-server
```

The server listens on `:8080` by default. Set `ADDR` to change the listen
address, for example `ADDR=127.0.0.1:9000`.

Endpoints:

- `GET /healthz` returns `ok`.
- `GET /v1/entropy` returns 16 random bytes as JSON by default.
- `GET /v1/entropy?bytes=32` returns the requested number of bytes (1–1024).
  JSON encodes the `data` byte slice as base64.

The HTTP handler is also reusable independently via `server.NewHandler()`.
