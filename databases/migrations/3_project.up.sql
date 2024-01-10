-- auto-generated definition
create table IF NOT EXISTS projects
(
    project_id  bigint auto_increment primary key,
    name        varchar(255) null,
    created_at  timestamp            null,
    updated_at  timestamp            null,
    deleted_at  timestamp            null
);

-- auto-generated definition
create table IF NOT EXISTS project_user
(
    project_id  bigint not null,
    user_id     bigint not null,
    created_at  timestamp            null,
    updated_at  timestamp            null,
    deleted_at  timestamp            null
);
