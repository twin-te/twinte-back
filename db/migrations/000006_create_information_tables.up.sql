BEGIN;

-- Table: public.already_reads

-- DROP TABLE IF EXISTS public.already_reads;

CREATE TABLE IF NOT EXISTS public.already_reads
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    information_id uuid NOT NULL,
    read_user text COLLATE pg_catalog."default" NOT NULL,
    read_at timestamp without time zone NOT NULL,
    CONSTRAINT "PK_7b7e4dc4f3dac9076561acd7832" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.already_reads
    OWNER to postgres;

COMMIT;