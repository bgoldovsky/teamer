-- drop table if exists persons;

create table if not exists persons (
    id             bigserial    primary key,
    team_id        bigint       references teams (id) on delete cascade,
    --
    first_name     text         not null,
    middle_name    text,
    last_name      text         not null,
    birthday       timestamp    not null,
    email          text,
    phone          text,
    slack          text         not null,
    role           int          not null check (role>0),
    duty_order     int          unique,
    is_active      boolean      not null,
    --
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);