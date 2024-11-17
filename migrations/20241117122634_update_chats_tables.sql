-- +goose Up
-- +goose StatementBegin
alter table chats alter column title type varchar(50);
alter table messages alter column user_tag type varchar(50);
alter table users_chats alter column user_tag type varchar(50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
