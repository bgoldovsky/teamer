-- drop table if exists persons;

create table if not exists teams (
    id          bigserial   primary key,
    --
    person_id   text    references persons (id) on delete cascade,
    --
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);