CREATE TABLE transactions (
    id TEXT PRIMARY KEY,
    sender TEXT, -- Should be something like bank account, but for simplicity
    receiver TEXT,
    amount float,
    status integer,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);