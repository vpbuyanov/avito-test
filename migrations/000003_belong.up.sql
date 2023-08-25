CREATE TABLE IF NOT EXISTS "Belong"(
    belong_id serial primary key,
    user_id int,
    segments_id int,
    FOREIGN KEY (user_id) REFERENCES public."Users" (user_id) ON DELETE CASCADE,
    FOREIGN KEY (segments_id) REFERENCES public."Segments" (segments_id) ON DELETE CASCADE
);
