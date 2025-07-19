-- FILE: migrations/0001_initial_schema.up.sql
-- This file contains the SQL statements to apply the migration.

-- Create the users table based on the GORM model in the users module.
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    name TEXT,
    email TEXT UNIQUE
);

-- Add an index on the soft-delete column for performance.
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);

-- Create the sessions table required by the Fiber session middleware.
-- The structure is based on the gofiber/storage/postgres/v2 driver's requirements.
CREATE TABLE IF NOT EXISTS fiber_sessions (
    id VARCHAR(64) PRIMARY KEY,
    data BYTEA,
    exp BIGINT
);