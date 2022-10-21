-- Write your migrate up statements here

create table users
(
    id       uuid default gen_random_uuid() not null constraint users_pk unique primary key,
    login    varchar                        not null unique,
    password varchar                        not null
);

---- create above / drop below ----

drop table users;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
