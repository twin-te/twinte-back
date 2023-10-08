BEGIN;

-- Type: registered_courses_methods_enum

-- DROP TYPE IF EXISTS public.registered_courses_methods_enum;

CREATE TYPE public.registered_courses_methods_enum AS ENUM
    ('OnlineAsynchronous', 'OnlineSynchronous', 'FaceToFace', 'Others');

ALTER TYPE public.registered_courses_methods_enum
    OWNER TO postgres;

-- Table: public.registered_courses

-- DROP TABLE IF EXISTS public.registered_courses;

CREATE TABLE IF NOT EXISTS public.registered_courses
(
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    year smallint NOT NULL,
    course_id uuid,
    name text COLLATE pg_catalog."default",
    instractor text COLLATE pg_catalog."default",
    credit numeric,
    methods registered_courses_methods_enum[],
    schedules jsonb,
    memo text COLLATE pg_catalog."default" NOT NULL,
    attendance integer NOT NULL,
    absence integer NOT NULL,
    late integer NOT NULL,
    CONSTRAINT "PK_1aa4b72144999c7ac6f37e52fca" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.registered_courses
    OWNER to postgres;

-- Index: IDX_fbc9587b218000acc37d7c6385

-- DROP INDEX IF EXISTS public."IDX_fbc9587b218000acc37d7c6385";

CREATE UNIQUE INDEX IF NOT EXISTS "IDX_fbc9587b218000acc37d7c6385"
    ON public.registered_courses USING btree
    (user_id ASC NULLS LAST, course_id ASC NULLS LAST)
    TABLESPACE pg_default;

-- Table: public.tags

-- DROP TABLE IF EXISTS public.tags;

CREATE TABLE IF NOT EXISTS public.tags
(
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    "position" integer NOT NULL,
    CONSTRAINT "PK_e7dc17249a1148a1970748eda99" PRIMARY KEY (id),
    CONSTRAINT "UQ_b861cf6ec9af09190780481a311" UNIQUE (user_id, "position")
        DEFERRABLE INITIALLY DEFERRED
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tags
    OWNER to postgres;

-- Table: public.registered_course_tags

-- DROP TABLE IF EXISTS public.registered_course_tags;

CREATE TABLE IF NOT EXISTS public.registered_course_tags
(
    tag uuid NOT NULL,
    registered_course uuid NOT NULL,
    CONSTRAINT "PK_605d76e7c4a818b543daae4cdf2" PRIMARY KEY (tag, registered_course),
    CONSTRAINT "FK_41594bc21c7a7adff5f2a4574cc" FOREIGN KEY (registered_course)
        REFERENCES public.registered_courses (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_c102260e1ffed113fef8ead1481" FOREIGN KEY (tag)
        REFERENCES public.tags (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.registered_course_tags
    OWNER to postgres;

-- Index: IDX_41594bc21c7a7adff5f2a4574c

-- DROP INDEX IF EXISTS public."IDX_41594bc21c7a7adff5f2a4574c";

CREATE INDEX IF NOT EXISTS "IDX_41594bc21c7a7adff5f2a4574c"
    ON public.registered_course_tags USING btree
    (registered_course ASC NULLS LAST)
    TABLESPACE pg_default;

-- Index: IDX_c102260e1ffed113fef8ead148

-- DROP INDEX IF EXISTS public."IDX_c102260e1ffed113fef8ead148";

CREATE INDEX IF NOT EXISTS "IDX_c102260e1ffed113fef8ead148"
    ON public.registered_course_tags USING btree
    (tag ASC NULLS LAST)
    TABLESPACE pg_default;

COMMIT;
