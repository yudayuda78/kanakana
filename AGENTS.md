# AGENTS.md

REST API for Japanese vocabulary (Go + GORM + PostgreSQL). Module name is `kanakana`.

## Commands

Use the `Makefile`, not the README (the README's `go run main.go` is wrong).

```bash
make run     # go run cmd/kanakana/main.go
make dev     # air auto-reload (install once: make install-air)
make build   # go build -o bin/kanakana cmd/kanakana/main.go
make test    # go test ./... -v
make seed    # go run cmd/seeder/main.go  (creates admin only)
```

No CI, linter, or formatter config exists. Verify changes with `go build ./... && go vet ./... && go test ./...` (or just `make test`). Use `gofmt`.

## Setup

- Copy env vars into a root `.env` (gitignored, but present locally). See `internal/config/config.go` for keys and defaults: `PORT`, `DB_*`, `JWT_SECRET`, `JWT_EXPIRATION_HOURS`.
- `godotenv.Load()` reads `.env` from the process CWD, so run binaries from the repo root (`make run`). Running `bin/kanakana` from elsewhere skips `.env`.
- DB is auto-migrated on startup (`db.AutoMigrate` in `internal/app/kanakana/app.go`). No migration files.
- `make seed` creates `admin@kanakana.com` / `admin123`. It does NOT seed levels or kosakata, and only migrates `User` — create `KosakataLevel` rows yourself.

## Tests

- `go test ./... -v` passes without a DB: DB tests skip when `DB_*` env vars are unset (`internal/pkg/database/db_test.go`).
- DB tests load `.env` via a path relative to the package dir (`../../../../.env`); run `go test` from the repo root, not an arbitrary CWD.
- Handler tests inject `MockKosakataService` (implements `services.KosakataService`) and call handlers directly — no server needed.

## Architecture

Layered, manual dependency injection wired in `internal/app/kanakana/app.go` (`StartServer`):

`handlers` -> `services` (interfaces) -> `repository` (interfaces) -> GORM.

- Entrypoint: `cmd/kanakana/main.go` -> `kanakana.StartServer()`. A second entrypoint `cmd/seeder/main.go` seeds the admin.
- All routes are registered in `registerRoutes` in `internal/app/kanakana/app.go`. Go 1.22+ method patterns (`GET /api/v1/...`) on `http.ServeMux`.
- Routes use query params, not path params: `?id=`, `?level=`, `?q=`.
- Auth: JWT HS256; claims `user_id` and `role`; middleware in `internal/middleware/auth_middleware.go`. `RequireRole` must run after `RequireAuth`.
- Adding an entity means touching every layer: model, repository interface+impl, service interface+impl, handler, route, and the `AutoMigrate` call. Many2many relations also need `db.SetupJoinTable` in `app.go` (see `UserLevel`, `UserKosakata`).

## Conventions / gotchas

- Code comments and most user-facing messages are written in Indonesian. Match this.
- Models embed `gorm.Model` (soft delete). Passwords use `json:"-"`; never serialize them.
- `README.md` and `API_DOCS.md` are stale: they mention `lib/pq` and omit GORM/JWT; register returns only a message, login returns `{token, user}`, and several user/levels endpoints are undocumented. Trust code and `app.go`.
- `internal/pkg/database` uses GORM's pgx driver (not `lib/pq`).
