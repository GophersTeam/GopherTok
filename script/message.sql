create schema gophertok_chat;
create table gophertok_chat.message
(
    id           bigint auto_increment
        primary key,
    from_user_id bigint        not null,
    to_user_id   bigint        not null,
    content      varchar(255)  not null,
    create_time  int default 0 not null
);

