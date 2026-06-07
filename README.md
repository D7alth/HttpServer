# HttpServer

A simple HTTP server written in Go to practice the standard library's `net/http` package. The goal is to keep expanding it as I get more comfortable with the language.

## Routes

| Method | Path | Input | Description |
|--------|------|-------|-------------|
| GET | `/` | | Root greeting |
| GET | `/goodbye` | | Returns a goodbye message |
| GET | `/hello` | `?user=<name>` query param | Greets the given user |
| GET | `/responses/{user}/hello` | Path parameter | Greets the user from the URL path |
| GET | `/responses/hello` | `user` header | Greets the user from a request header |
| POST | `/responses/hello/json` | JSON body `{"Name": "..."}` | Greets the user from a JSON body |

## Running

```sh
go run main.go
```

## Tests

```sh
go test ./...
```
