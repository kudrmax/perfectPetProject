create table if not exists "user"
(
    id           serial primary key,
    name         text not null,
    username     text not null,
    passwordHash text not null
)
