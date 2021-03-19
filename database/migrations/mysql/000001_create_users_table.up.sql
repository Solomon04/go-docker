create table if not exists `users`
(
    id                      bigint unsigned auto_increment primary key,
    uuid                    char(36)    unique                                                 not null,
    first_name               varchar(255)                                                       null,
    last_name                varchar(255)                                                       null,
    email                   varchar(255) unique                                                      not null,
    oauth_provider          varchar(255)                                                       null,
    oauth_identifier          varchar(255)                                                       null,
    password                varchar(255)                                                       null,
    avatar                  varchar(255) default 'https://hoopspots.s3.amazonaws.com/user.png' not null,
    last_seen_at              datetime                                                           null,
    last_seen_location        text                                                               null,
    created_at               timestamp                                                          null,
    updated_at               timestamp                                                          null,
    deleted_at               timestamp                                                          null
)