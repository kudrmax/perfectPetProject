create table if not exists "users"
(
    id           serial primary key,
    name         text not null,
    username     text not null unique,
    passwordHash text not null
)
