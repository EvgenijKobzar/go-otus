CREATE SCHEMA if not exists movies_online;

create table if not exists movies_online.serial
(
    id        serial primary key,
    sort      integer,
    active    char(1) default 'Y',
    file_id int DEFAULT NULL,
    title varchar(50) DEFAULT NULL,
    production_period varchar(50) DEFAULT NULL,
    rating float DEFAULT NULL,
    quality varchar(50) DEFAULT NULL,
    duration float DEFAULT NULL,
    description varchar(150) DEFAULT NULL,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
    );