-- auto-generated definition
create table IF NOT EXISTS tasks
(
    task_id     bigint auto_increment primary key,
    description varchar(255) null,
    title       varchar(255) null,
    project_id  bigint null,
    created_at  timestamp            null,
    updated_at  timestamp            null,
    deleted_at  timestamp            null
);
