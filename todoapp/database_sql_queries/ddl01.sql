CREATE TABLE IF NOT EXISTS public.todo_details
(
    id integer NOT NULL DEFAULT nextval('books_id_seq'::regclass),
    title text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    completed boolean,
    CONSTRAINT todo_details_pkey PRIMARY KEY (id)
)
