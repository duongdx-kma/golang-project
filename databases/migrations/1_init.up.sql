-- auto-generated definition
CREATE TABLE IF NOT EXISTS  users
(
    id         bigint auto_increment primary key,
    name       varchar(255)         null,
    address    varchar(255)         null,
    password   varchar(255)         not null,
    age        int                  null,
    is_admin   tinyint(1) default 0 null,
    created_at timestamp            null,
    updated_at timestamp            null,
    deleted_at timestamp            null,
    constraint constraint_name      unique (name)
);

INSERT INTO users (
    name,
    address,
    password,
    age,
    is_admin,
    created_at,
    updated_at,
    deleted_at
) SELECT
    'duongdx',
    'VN',
    '$2a$10$mozDKCPtZJYFGzqsY/rmrOhG98m5WljKS/FTFW82oY8xwwwh1XFkG', -- password
    20,
    1,
    null,
    null,
    null
 FROM DUAL
WHERE NOT EXISTS
  (SELECT name FROM users WHERE name='duongdx');
