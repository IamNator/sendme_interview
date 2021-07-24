create table "user"
(
    id               serial not null
        constraint user_pk
            primary key,
    username         varchar(255),
    hashed_password  integer,
    token            text,
    token_expiration timestamp,
    created_at       integer,
    updated_at       timestamp,
    deleted_at       timestamp
);

alter table "user"
    owner to zskonvzx;

create unique index user_hashed_password_uindex
    on "user" (hashed_password);

create unique index user_id_uindex
    on "user" (id);

create unique index user_token_uindex
    on "user" (token);

