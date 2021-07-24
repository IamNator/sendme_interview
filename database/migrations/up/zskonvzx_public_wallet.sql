create table wallet
(
    id         serial                              not null
        constraint wallet_pk
            primary key,
    user_id    integer
        constraint wallet_user_id_fk
            references "user"
            on update cascade on delete cascade,
    balance    integer,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp
);

alter table wallet
    owner to zskonvzx;

create unique index wallet_id_uindex
    on wallet (id);

create unique index wallet_user_id_uindex
    on wallet (user_id);

