create table events
(
    id           bigint unsigned auto_increment
        primary key,
    uuid         char(36) unique not null,
    creator_id  bigint unsigned       null,
    location_id bigint unsigned       not null,
    start_time    datetime        not null,
    end_time      datetime        not null,
    lat          double(10, 7)   not null,
    `long`       double(10, 7)   not null,
    live_count    int default 0   not null,
    created_at    timestamp       null,
    updated_at    timestamp       null,
    foreign key (creator_id) references users (id),
    foreign key (location_id) references locations (id)
)

