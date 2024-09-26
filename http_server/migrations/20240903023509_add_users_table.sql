-- +goose Up
-- +goose StatementBegin
create table users (
    id serial primary key,
    login varchar(255) not null unique,
    password varchar(255) not null,
    role_id int references users(id)
);
insert into users(login, password, role_id) values ('rasank', '$2a$12$DB6uqFXjukTcH4GJlEvU8.caJJN.zX6dgltnFphNJ3aAVmHtCtLhO', 1); -- pass --> "123123"
-- +goose StatementEnd

-- +goose Down
drop table if exists users;
