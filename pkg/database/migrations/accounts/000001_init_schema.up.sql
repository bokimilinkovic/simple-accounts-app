CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    owner TEXT NOT NULL,
    balance float,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);