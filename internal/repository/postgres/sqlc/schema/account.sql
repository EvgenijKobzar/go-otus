CREATE SCHEMA if not exists movies_online;

create table if not exists movies_online.account
(
    id        serial primary key,
    firstName varchar(150) DEFAULT NULL,
    lastName varchar(150) DEFAULT NULL,
    login varchar(150) DEFAULT NULL,
    Password varchar(150) DEFAULT NULL,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
    );