## Makefile commands

| Command         | Description                               |
|-----------------|-------------------------------------------|
| `make dev`      | Run app with Air (hot reload)             |
| `make run`      | Run app normally (`go run ./cmd/api`)     |
| `make build`    | Build the project                        |
| `make fmt`      | Format code with `go fmt`                 |
| `make imports`  | Fix imports with `goimports`              |
| `make lint`     | Run linters (local, may auto-fix)         |
| `make lint-ci`  | Run linters (CI mode, check only)         |
| `make test`     | Run all tests                             |
| `make tidy`     | Clean up `go.mod` and `go.sum`            |
| `make check`    | Format + imports + lint                   |
| `make check-all`| Full cycle: fmt, lint, build, test, tidy  |
| `make migrate-*`| DB migrations (up, down, reset, new, ...) |
