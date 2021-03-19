create table locations
(
    id                  bigint unsigned auto_increment
        primary key,
    uuid                char(36) unique                                                                                                                                           not null,
    name                varchar(255)                                                                                                                                              not null,
    slug                varchar(255)                                                                                                                                              null,
    image              varchar(255) default 'https://www.lifetime.life/content/dam/ltp/images/locations/specific-club/mn-southdale/mkin20314486-edina-at-southdale-hero-xl3.jpg' not null,
    description         mediumtext                                                                                                                                                null,
    location_type        varchar(255) default 'lifetime_fitness'                                                                                                                   not null,
    street              varchar(255)                                                                                                                                              not null,
    city                varchar(255)                                                                                                                                              not null,
    state               varchar(255)                                                                                                                                              not null,
    zip                 varchar(255)                                                                                                                                              not null,
    lat                 double(10, 7)                                                                                                                                             not null,
    `long`              double(10, 7)                                                                                                                                             not null,
    schedule_rule        text                                                                                                                                                      null,
    is_active            tinyint(1)   default 0                                                                                                                                    not null,
    has_covid_restriction tinyint(1)   default 0                                                                                                                                    not null,
    appointment_type_id   varchar(255)                                                                                                                                              null,
    calendar_id          varchar(255)                                                                                                                                              null,
    created_at           timestamp                                                                                                                                                 null,
    updated_at           timestamp                                                                                                                                                 null
)

