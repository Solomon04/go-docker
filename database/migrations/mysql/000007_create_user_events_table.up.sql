create table user_events
(
    id               bigint unsigned auto_increment
        primary key,
    event_id        bigint unsigned   not null,
    user_id         bigint unsigned   not null,
    responded_at      datetime   not null,
    is_going          tinyint(1) not null,
    attended_at       datetime   null,
    foreign key (event_id) references events (id),
    foreign key (user_id) references users (id)
)

