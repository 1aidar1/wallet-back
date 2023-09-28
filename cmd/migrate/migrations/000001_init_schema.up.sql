CREATE EXTENSION IF NOT EXISTS pgcrypto;


CREATE OR REPLACE FUNCTION generate_ulid() RETURNS uuid
    AS $$
        SELECT (lpad(to_hex(floor(extract(epoch FROM clock_timestamp()) * 1000)::bigint), 12, '0') || encode(gen_random_bytes(10), 'hex'))::uuid;
    $$ LANGUAGE SQL;

DROP TYPE IF EXISTS wallet_status;
CREATE TYPE wallet_status AS ENUM (
  'active',
  'blocked',
  'closed'
);

DROP TYPE IF EXISTS operation_type;
CREATE TYPE operation_type AS ENUM (
  'refill',
  'withdraw',
  'hold',
  'clear',
  'unhold',
  'refund'
);

DROP TYPE IF EXISTS transaction_type;
CREATE TYPE transaction_type AS ENUM (
    'pay',
    'transfer'
);

DROP TYPE IF EXISTS operation_status;
CREATE TYPE operation_status AS ENUM (
  'process',
  'success',
  'error'
);

DROP TYPE IF EXISTS wallet_identification_status;
CREATE TYPE "wallet_identification_status" AS ENUM (
  'full',
  'basic',
  'none'
);


create table if not exists currencies (
    id uuid primary key default gen_random_uuid(),
    code varchar(100) not null unique
);

create table if not exists countries (
    id uuid primary key default gen_random_uuid(),
    code varchar(100) not null unique
);


create table if not exists "wallets" (
    id uuid primary key default gen_random_uuid(),
    currency_id uuid not null,
    country_id uuid not null,
    balance bigint not null,
    hold bigint not null,
    status wallet_status not null,
    identification wallet_identification_status not null,
    phone varchar(20),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    constraint fk_currency_id foreign key (currency_id) references currencies (id),
    constraint fk_country_id foreign key (country_id) references countries (id),
    constraint unique_phone unique (phone)
);

create table if not exists consumers (
    id                 uuid primary key default gen_random_uuid(),
    code               varchar(50)  not null,
    slug               varchar(255) not null,
    secret             varchar(255) not null,
    white_list_methods text[],
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp
);
create unique index "unique_consumer_code" on consumers(code) where deleted_at is null;

create table if not exists transactions (
    id uuid primary key default generate_ulid(),
    consumer_id uuid not null,
    type transaction_type not null,
    description varchar(255),
    order_id varchar(255) not null,
    service_provider_id text not null,
    created_at timestamp not null default now(),
    constraint fk_consumer_id foreign key (consumer_id) references consumers(id),
    constraint unique_order_id unique (order_id,service_provider_id)
);

create table error_codes (
    code varchar(36) primary key,
    description text not null default '',
    http_code int not null
);

create table if not exists operations (
    id uuid primary key default generate_ulid(),
    transaction_id uuid not null,
    type operation_type not null ,
    wallet_id uuid not null,
    amount bigint not null,
    status operation_status not null ,
    error_code varchar(36),
    internal_log_message text not null default '',
    balance_before bigint not null ,
    balance_after bigint not null,
    hold_balance_before bigint not null ,
    hold_balance_after bigint not null,
    created_at timestamp not null default now(),
    constraint fk_transaction_id foreign key (transaction_id) references transactions(id),
    constraint fk_wallet_id foreign key (wallet_id) references wallets (id)
);

create table if not exists wallet_identification_history (
    id uuid primary key default generate_ulid(),
    wallet_id uuid not null,
    status wallet_identification_status not null,
    created_at timestamp not null default now(),
    constraint fk_wallet_id foreign key (wallet_id) references wallets (id)
);

create table wallet_status_history (
    id uuid primary key default generate_ulid(),
    wallet_id uuid not null,
    status wallet_status not null ,
    description text not null default '',
    created_at timestamp not null default now(),
    constraint fk_wallet_id foreign key (wallet_id) references wallets (id)
);