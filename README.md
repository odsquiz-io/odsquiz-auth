# ODS Quiz Auth API

Authentication API for ODS Quiz. It uses Go, Fiber, GORM, and PostgreSQL.

## Run locally

Prerequisites:

- Go version declared in `go.mod`.
- PostgreSQL listening on `127.0.0.1:5432`.
- A local database and role that match your `.env.local` credentials.

Check that PostgreSQL is available:

```bash
pg_isready -h 127.0.0.1 -p 5432
```

Create the local configuration once:

```bash
cp .env.example .env.local
```

Set the values in `.env.local` without committing it:

```dotenv
PORT=8080
JWTSecret=local-development-secret
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=your_local_user
DB_PASSWORD=your_local_password
DB_NAME=odsquiz
DB_SSLMODE=disable
```

Run the API:

```bash
go run ./cmd/api
```

At startup, the API checks the database connection, retries transient failures
with exponential backoff, and runs its GORM migration. Verify it in another
terminal:

```bash
curl --fail-with-body http://127.0.0.1:8080/health
```

Expected response:

```text
ok
```

## Validation

```bash
go test ./...
go vet ./...
docker build -t odsquiz-auth .
```

## API paths

- `POST /api/auth/signup`
- `POST /api/auth/login`
- `GET /api/auth/me`
- `PATCH /api/auth/me`
- `GET /health`

The public load balancer exposes this service through `/api/auth/*`.

## Configuration convention

Use `DB_SSLMODE` in all new configuration. The application temporarily accepts
the older `DBSSLMode` spelling to avoid breaking existing developer machines.
