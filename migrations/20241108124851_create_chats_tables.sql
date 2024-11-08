-- +goose Up
-- +goose StatementBegin
create table chat (
    id serial primary key,
    title text not null,
    created_at timestamp default now(),
    updated_at timestamp
);

create table user_chat (
    id serial primary key,
    chat_id integer not null references chat,
    user_id integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chat;
drop table user_chat;
-- +goose StatementEnd
