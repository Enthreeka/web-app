CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
    id uuid DEFAULT uuid_generate_v4(),
    login varchar(100) not null unique ,
    password varchar(200) not null,
    token varchar(200) DEFAULT 0,
    primary key (id)


);

create table account
(
    id int generated always as identity,
    user_id uuid,
    name varchar(50) constraint users_name null,
    email varchar(100) null ,
    photo varchar(200) null ,
    subscribe boolean null ,
    date_signup date DEFAULT now(),
    primary key (id),
    constraint fk_users
        foreign key (user_id)
            references users (id)

);

create table tasks
(
    id int generated always as identity ,
    user_id uuid,
    name_task varchar(250) null ,
    description_task text null,
    date_task date constraint date_create_task DEFAULT now(),
    primary key (id),
    constraint fk_account
        foreign key (user_id)
            references users (id)
);

-- CREATE TABLE tokens
-- (
--     id int generated always as identity,
--     user_id int not null,
--     token varchar(200) not null,
--     created_at timestamp default current_timestamp,
--     expires_at timestamp not null,
--     primary key (id),
--     constraint fk_users
--         foreign key (user_id)
--             references users (id)
-- );

--n3ksmirn = 1fsdjhj123hhjd
    --ilisTopskiy = ;;1afkasfo34
--

INSERT INTO tasks
(account_id,name_task,description_task)
VALUES
    (2,'Тест имени','Тест')
RETURNING id;

insert into tasks (account_id, name_task, description_task) values (1,'gqfz3123','gola123fasng');

delete from tasks where account_id= 2;

SELECT name_task , description_task
FROM tasks
WHERE account_id = 2;

drop table tasks;
drop table account;
drop table users;