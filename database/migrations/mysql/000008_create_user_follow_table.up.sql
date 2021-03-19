create table if not exists user_follow
(
    id                int unsigned auto_increment
        primary key,
    follower_user_id  bigint unsigned not null,
    following_user_id  bigint unsigned not null,
    accepted_at        timestamp null,
    created_at         timestamp null,
    updated_at         timestamp null,
    foreign key (following_user_id) references users(id),
    foreign key (follower_user_id) references users(id)
)

