drop table if exists wallet_status_history;
drop table if exists wallet_identification_history;
drop table if exists operations;
drop table if exists error_codes;
drop table if exists transactions;
drop table if exists consumers;
drop table if exists wallets;
drop table if exists countries;
drop table if exists currencies;
drop type if exists wallet_status;
drop type if exists operation_type;
drop type if exists transaction_type;
drop type if exists operation_status;
drop function if exists generate_ulid();
drop extension if exists pgcrypto;

