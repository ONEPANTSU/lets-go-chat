create extension if not exists "uuid-ossp";

create table chat
(
    id         uuid               default uuid_generate_v4() primary key,
    name       text      not null,
    owner_id   uuid      not null,
    created_at timestamp not null default now(),
    deleted_at timestamp
);

create table user_chat
(
    user_id   uuid      not null,
    chat_id   uuid      not null references chat (id) on delete cascade,
    joined_at timestamp not null default now(),

    primary key (user_id, chat_id)
);

create table message
(
    id         serial primary key,
    user_id    uuid      not null,
    chat_id    uuid      not null references chat (id) on delete cascade,
    text       text      not null,
    created_at timestamp not null default now(),
    deleted_at timestamp
);