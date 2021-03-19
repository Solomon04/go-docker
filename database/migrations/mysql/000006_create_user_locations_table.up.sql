create table if not exists user_locations
(
    id           bigint unsigned auto_increment
        primary key,
    location_id bigint unsigned not null,
    user_id     bigint unsigned not null,
    joined_on     datetime not null,
    foreign key (location_id) references locations (id),
    foreign key (user_id) references users (id)
)

