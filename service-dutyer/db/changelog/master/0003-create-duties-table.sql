-- drop table if exists duties;

create table if not exists duties (
    team_id     bigint      unique references teams (id) on delete cascade,
    person_id   bigint      unique references persons (id) on delete cascade,
    month       integer     not null,
    day         integer     not null,
    --
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);