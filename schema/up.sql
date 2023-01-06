CREATE TABLE users (
    id serial not null unique, 
    fullname varchar(255) not null, 
    username varchar(255) not null, 
    password_hash varchar(255) not null
);

CREATE TABLE lists (
    id serial not null unique, 
    title varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE users_lists (
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    list_id int references lists(id) on delete cascade not null
);

CREATE TABLE items (
    id serial not null unique, 
    title varchar(255) not null,
    description varchar(255) not null,
    done boolean not null default false
);

CREATE TABLE lists_items (
    id serial not null unique,
    list_id int references lists(id) on delete cascade not null,
    item_id int references items(id) on delete cascade not null
);
 