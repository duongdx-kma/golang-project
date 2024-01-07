-- auto-generated definition
create table IF NOT EXISTS tasks
(
    task_id     int auto_increment
        primary key,
    description varchar(255) null,
    title       varchar(255) null,
    user_id     varchar(255) null
);

