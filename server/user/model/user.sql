create table user
(
    id               bigint auto_increment
        primary key,
    username         varchar(32)  not null,
    password         varchar(32)  not null,
    avatar           varchar(255) null,
    background_image varchar(255) null,
    signature        varchar(255) null,
    constraint id
        unique (id),
    constraint username
        unique (username)
);

