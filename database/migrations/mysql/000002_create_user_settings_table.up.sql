create table if not exists user_settings
(
    id                      bigint unsigned auto_increment primary key,
    user_id                 bigint unsigned                                                        not null,
    preferred_location_types  text                                                               null,
    preferred_distance       int          default 25                                            not null,
    primary_city             varchar(255)                                                       null,
    primary_state            varchar(255)                                                       null,
    primary_zip_code          varchar(255)                                                       null,
    primary_lat              double(10, 7)                                                      null,
    primary_long             double(10, 7)                                                      null,
    notify_comments          tinyint(1)   default 1                                             not null,
    notify_events            tinyint(1)   default 1                                             not null,
    notify_recommendations   tinyint(1)   default 1                                             not null,
    notification_follower_activity  tinyint(1)   default 1                                             not null,
    created_at               timestamp                                                          null,
    updated_at               timestamp                                                          null,
    foreign key (user_id) references users(id)
)

