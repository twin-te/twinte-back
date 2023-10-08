BEGIN;

-- Table: public.session

-- DROP TABLE IF EXISTS public.session;

CREATE TABLE IF NOT EXISTS public.session
(
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    expired_at timestamp(3) without time zone NOT NULL,
    CONSTRAINT session_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.session
    OWNER to postgres;
-- Index: session.id_expired_at_index

-- DROP INDEX IF EXISTS public."session.id_expired_at_index";

CREATE INDEX IF NOT EXISTS "session.id_expired_at_index"
    ON public.session USING btree
    (id ASC NULLS LAST, expired_at ASC NULLS LAST)
    TABLESPACE pg_default;

COMMIT;