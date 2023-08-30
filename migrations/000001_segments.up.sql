CREATE TABLE IF NOT EXISTS "Segments"(
    segments_id serial primary key,
    slug varchar(50) unique not null
);