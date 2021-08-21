CREATE TABLE "fib" (
    "id" bigserial PRIMARY KEY,
    "index" bigint NOT NULL,
    "value"  bigint NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT (now())
)