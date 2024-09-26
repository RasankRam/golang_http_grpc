-- +goose Up
-- +goose StatementBegin
create table roles (
    id serial primary key,
    nm varchar(255) not null unique
);
insert into roles(nm) values ('admin');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists roles;
-- +goose StatementEnd
