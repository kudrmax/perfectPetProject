create table if not exists `twits`
(
    id serial primary key,
    user_id integer foreign key references users(id),
    text text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
--     TODO добавить удаление через архивирование
)
