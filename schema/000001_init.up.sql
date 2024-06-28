CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    username varchar(255) not null unique,
    name varchar(255) not null,
    password_hash varchar(255) not null
);
CREATE TABLE chats
(
    id SERIAL PRIMARY KEY,
    first_user_id int references users (id) on delete cascade not null,
    second_user_id int references users (id) on delete cascade not null,
    first_user_name varchar(255) not null,
    second_user_name varchar(255) not null
);
CREATE TABLE messages
(
    id SERIAL PRIMARY KEY,
    chat_id int references chats (id) on delete cascade not null,
    text varchar(300) not null,
    user_id int references users (id) on delete cascade not null
);
CREATE TABLE rooms
(
    id SERIAL PRIMARY KEY,
    users_quantity BIGINT not null,
    first_user_id int references users (id) on delete cascade not null,
    second_user_id int references users (id) on delete cascade not null
);