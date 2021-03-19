create table personal_access_tokens
(
    ID         bigint unsigned auto_increment
        primary key,
    user_id   bigint unsigned  not null,
    token      varchar(255) not null,
    abilities  text         null,
    last_used_at timestamp    null,
    created_at  timestamp    null,
    updated_at  timestamp    null,
    foreign key (user_id) references users (id)
)

