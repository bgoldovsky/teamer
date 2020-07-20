-- drop table if exists teams;

create table if not exists teams (
    id           bigserial  primary key,
    --
    name         text   not null,
    description  text,
    slack        text   not null,
    --
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);