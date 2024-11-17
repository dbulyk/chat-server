-- +goose Up
-- +goose StatementBegin
create table chats (
    id serial primary key,
    title text not null,
    created_at timestamp default now()
);

create table messages (
    id serial primary key,
    chat_id integer not null references chats,
    user_tag text not null,
    message text not null,
    created_at timestamp default now(),
    updated_at timestamp
);

create table users_chats (
    id serial primary key,
    chat_id integer not null references chats,
    user_tag text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats;
drop table messages;
drop table users_chats;
-- +goose StatementEnd
