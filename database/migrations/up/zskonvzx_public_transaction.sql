create table transaction
(
    id         serial  not null
        constraint transaction_pk
            primary key,
    user_id    integer not null
        constraint transaction_user_id_fk
            references "user"
            on update cascade on delete cascade,
    type       varchar(6),
    amount     bigint,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

alter table transaction
    owner to zskonvzx;

create unique index transaction_id_uindex
    on transaction (id);

