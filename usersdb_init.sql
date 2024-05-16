-- usersdb_init.sql
-- Create a table using SERIAL type and insert data into the usersdb database

CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY,
    person_id bigint,
    username text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    role bigint,
    is_active boolean,
    reset_password_token text COLLATE pg_catalog."default",
    email_verification_token text COLLATE pg_catalog."default"
);

INSERT INTO public.users(
    id, username, password, role, is_active, reset_password_token, email_verification_token, person_id)
VALUES 
    (-4, 'turista', '$2a$10$DE9tdj0AIpgAN1.TtfVHDuinqJqwU57fxejSZyCmertgqhzXJi90K', 2, true, null, null, -4);
INSERT INTO public.users(
    id, username, password, role, is_active, reset_password_token, email_verification_token, person_id)
VALUES 
    (-1, 'admin', '$2a$10$DE9tdj0AIpgAN1.TtfVHDuinqJqwU57fxejSZyCmertgqhzXJi90K', 0, true, null, null, -1);
    INSERT INTO public.users(
    id, username, password, role, is_active, reset_password_token, email_verification_token, person_id)
VALUES 
    (-2, 'autor', '$2a$10$DE9tdj0AIpgAN1.TtfVHDuinqJqwU57fxejSZyCmertgqhzXJi90K', 1, true, null, null, -2);