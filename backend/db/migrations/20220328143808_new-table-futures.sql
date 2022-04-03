-- migrate:up
create table if not exists futures
(
    id serial not null
        constraint futures_pkey
            primary key,
    datetime timestamptz not null default now(),
    tiker varchar(16),
    is_open boolean default true not null,
    warranty_provision decimal not null,
    count integer not null,
    amount decimal not null,
    margin decimal default 0 not null,
    commission decimal not null,
    commission_amount decimal not null
);

create index if not exists futures_datetime_idx on futures(datetime);
create index if not exists futures_tiker_idx on futures(tiker);
create index if not exists futures_is_open_idx on futures(is_open);

-- migrate:down

