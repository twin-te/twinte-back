BEGIN;

-- Table: public.payment_users

-- DROP TABLE IF EXISTS public.payment_users;

CREATE TABLE IF NOT EXISTS public.payment_users
(
    id text COLLATE pg_catalog."default" NOT NULL,
    twinte_user_id uuid NOT NULL,
    display_name text COLLATE pg_catalog."default",
    link text COLLATE pg_catalog."default",
    CONSTRAINT "PK_1438bd4715f036ae4353ae95505" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.payment_users
    OWNER to postgres;

COMMIT;