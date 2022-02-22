-- migrate:up
alter table trading_logs add column if not exists tiker_type varchar(64) default 'equity';
alter table trading_logs add column if not exists currency varchar(64) default 'rub';

-- migrate:down

