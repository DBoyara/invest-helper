-- migrate:up
create table if not exists users (
    id serial not null
        constraint users_pkey
            primary key,
    username varchar(64) unique,
    password varchar(255)
);

create table if not exists instruments (
    id serial not null
        constraint instruments_pkey
        primary key,
    tiker varchar(8) unique,
    type varchar(255)
);

create table if not exists trading_logs
(
    id serial not null
        constraint trading_logs_pkey
        primary key,
    datetime timestamptz not null default now(),
    tiker varchar(8),
    type varchar(16),
    is_open boolean default true not null,
    price decimal not null,
    count integer not null,
    lot integer not null,
    amount decimal not null,
    commission decimal not null,
    commission_amount decimal not null
);

create index if not exists trading_logs_datetime_idx on trading_logs(datetime);

create table if not exists commissions (
    id serial not null
        constraint commissions_pkey
        primary key,
    value decimal not null,
    type varchar(32)
);

INSERT INTO commissions (value, type)
VALUES 	(0.9, 'fixPrice'),
        (0.3, 'percent');
-- migrate:down
