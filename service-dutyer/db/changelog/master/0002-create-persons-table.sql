-- drop table if exists persons;

create table if not exists persons (
    id             bigserial    primary key,
    team_id        bigint       not null references teams (id) on delete restrict,
    --
    first_name     text         not null,
    middle_name    text,
    last_name      text         not null,
    birthday       timestamp,
    email          text,
    phone          text,
    slack          text         not null,
    role           int          not null check (role>0),
    duty_order     int          not null,
    is_active      boolean      not null,
    --
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null,
    unique (team_id,duty_order) deferrable initially deferred
);