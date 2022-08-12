BEGIN;

CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null unique,
    salt varchar(255),
    digest varchar(255)
);

CREATE TABLE roles
(
    id serial not null unique,
    name varchar(255) not null unique,
    rights varchar(255),
    user_id int references users (id) not null
);

CREATE TABLE pages
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    parent int references pages (id),
    _left int not null,
    _right int not null
);

CREATE TABLE page_types
(
    id serial not null unique,
    name varchar(255) not null,
    page_id int references pages (id) not null
);

COMMIT;
