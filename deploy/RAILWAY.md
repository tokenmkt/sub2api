# tokenMKT on Railway

This repo can be deployed into the existing Railway `tokenmkt` project as a
single code-deployed service backed by dedicated Railway managed PostgreSQL and
Redis services.

## Railway services

Create these services in the `tokenmkt` project:

- `sub2api`: GitHub/code deployment for this repository.
- `sub2api-pg`: Railway managed PostgreSQL for tokenMKT only.
- `sub2api-redis`: Railway managed Redis for tokenMKT only.

Do not share the existing `tokenmkt` Postgres or Redis services for a
high-traffic tokenMKT deployment. tokenMKT is a gateway workload with hot Redis
keys, request scheduling state, rate limits, session affinity, usage logs, and
monitoring writes; keeping its data services separate avoids latency, connection,
memory, and I/O contention with the storefront stack.

For the `sub2api` service, set:

- Root Directory: `/sub2api`
- Config File: `/sub2api/railway.toml`
- Volume mount: `/app/data`

Railway looks for Dockerfiles in the service root directory by default, so the
`/sub2api` root lets Railway build the existing `sub2api/Dockerfile`. The
`railway.toml` file pins the Dockerfile builder, limits autodeploys to
`/sub2api/**`, and enables the built-in `/health` endpoint as the deploy
healthcheck.

## Required variables

Set these variables on the `sub2api` service:

```bash
AUTO_SETUP=true
SERVER_HOST=0.0.0.0
SERVER_MODE=release
RUN_MODE=standard
TZ=Asia/Shanghai
DATA_DIR=/app/data

DATABASE_HOST=${{sub2api-pg.PGHOST}}
DATABASE_PORT=${{sub2api-pg.PGPORT}}
DATABASE_USER=${{sub2api-pg.PGUSER}}
DATABASE_PASSWORD=${{sub2api-pg.PGPASSWORD}}
DATABASE_DBNAME=${{sub2api-pg.PGDATABASE}}
DATABASE_SSLMODE=disable
DATABASE_MAX_OPEN_CONNS=20
DATABASE_MAX_IDLE_CONNS=5
DATABASE_CONN_MAX_LIFETIME_MINUTES=30
DATABASE_CONN_MAX_IDLE_TIME_MINUTES=5

REDIS_HOST=${{sub2api-redis.REDISHOST}}
REDIS_PORT=${{sub2api-redis.REDISPORT}}
REDIS_PASSWORD=${{sub2api-redis.REDISPASSWORD}}
REDIS_DB=0
REDIS_POOL_SIZE=128
REDIS_MIN_IDLE_CONNS=10
REDIS_ENABLE_TLS=false

ADMIN_EMAIL=admin@tokenmkt.cc
ADMIN_PASSWORD=<strong-random-password>
JWT_SECRET=<openssl-rand-hex-32>
TOTP_ENCRYPTION_KEY=<openssl-rand-hex-32>
```

`SERVER_PORT` is optional for Railway code deployments after this repository
change. Railway injects `PORT`, and `deploy/docker-entrypoint.sh` maps it to
`SERVER_PORT` when `SERVER_PORT` is not explicitly set. If you prefer a fixed
target port, set both values:

```bash
PORT=8080
SERVER_PORT=8080
```

## Recommended optional variables

Use these when you know the public domain before first boot:

```bash
SERVER_FRONTEND_URL=https://sub2api.tokenmkt.cc
LOG_FORMAT=json
LOG_OUTPUT_TO_STDOUT=true
LOG_OUTPUT_TO_FILE=false
```

OAuth, payment, Turnstile, and model-provider secrets can be added later from
the tokenMKT admin console or Railway variables.

## Deploy

From a linked workspace, a typical first deployment is:

```bash
railway link -p 0782eae8-4516-4b79-9425-70a7192942e0
railway add -s sub2api
railway add -d postgres -s sub2api-pg
railway add -d redis -s sub2api-redis
railway up /Users/aweminds/projects/token/tokenmkt/sub2api --path-as-root -s sub2api
```

Then bind the public domain to `sub2api`, for example:

```bash
railway domain -s sub2api
```

After first boot, confirm `https://<domain>/health` returns 200 and log in with
`ADMIN_EMAIL` / `ADMIN_PASSWORD`.

## Notes

- Do not reuse the root `railway/manage.sh` flow for tokenMKT; it is written for
  the existing `dujiaonext` `api/user/admin` stack.
- Keep the `/app/data` volume attached. tokenMKT writes `config.yaml`, the setup
  lock, logs, and runtime data there.
- `JWT_SECRET` and `TOTP_ENCRYPTION_KEY` should stay fixed across redeploys.
  Rotating them invalidates sessions and existing 2FA/payment encrypted data.
