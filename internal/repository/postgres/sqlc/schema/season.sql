CREATE SCHEMA if not exists movies_online;

create table if not exists movies_online.season
(
    id        serial primary key,
    serial_id int references movies_online.serial(id),
    sort      integer,
    active    char(1) default 'Y',
    title varchar(50) DEFAULT NULL,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
    );