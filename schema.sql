create table users (
    id serial primary key,
    uuid varchar(64) not null unique,
    name varchar(255),
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null
);

create table sessions (
    id serial primary key,
    uuid varchar(64) not null unique,
    email varchar(255),
    user_id integer references users(id),
    created_at timestamp not null
);

create table threads (
    id serial primary key,
    uuid varchar(64) not null unique,
    topic text,
    user_id integer references users(id),
    created_at timestamp not null
);

create table posts (
    id serial primary key,
    uuid varchar(64) not null unique,
    body text,
    user_id integer references users(id),
    thread_id integer references threads(id),
    created_at timestamp not null
);

create table posts2 (
    id serial primary key,
    content text,
    author varchar(255)
);

drop table if exists posts cascade;
drop table if exists comments;

create table posts (
    id serial primary key,
    content text,
    author varchar(255)
);

create table comments (
    id serial primary key,
    content text,
    author varchar(255),
    post_id integer references posts(id)
);
