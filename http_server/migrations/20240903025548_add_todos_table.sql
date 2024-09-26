-- +goose Up
-- +goose StatementBegin
create table todos(
    id serial primary key,
    title varchar(255) not null,
    dsc varchar(1000),
    created_by int references users(id),
    updated_by int references users(id),
    created_at timestamptz,
    updated_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists todos;
-- +goose StatementEnd
