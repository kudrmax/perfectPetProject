create table if not exists "twits"
(
    id         serial primary key,
    user_id    integer   not null,
    constraint fk_user foreign key (user_id) references users (id),
    text       text      not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
--     TODO добавить удаление через архивирование
)
