-- +goose Up
create table if not exists questions(
    id serial primary key,
    text varchar not null check (text <> ''),
    created_at timestamptz default current_timestamp
);

create table if not exists answers(
    id serial primary key,
    question_id int references questions on delete cascade,
    user_id uuid not null,
    text varchar not null check (text <> ''),
    created_at timestamptz default current_timestamp
);
-- +goose Down
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS questions;