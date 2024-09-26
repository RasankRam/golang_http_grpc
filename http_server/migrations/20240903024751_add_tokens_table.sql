-- +goose Up
-- +goose StatementBegin
CREATE TABLE tokens (
    id serial primary key,
    user_id int not null references users(id),
    token text not null,
    user_agent text not null,
    ip varchar(15) not null,
    constraint unique_user_agent_ip unique (user_agent, ip)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tokens;
-- +goose StatementEnd
