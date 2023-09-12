BEGIN;

-- Type: course_methods_method_enum

-- DROP TYPE IF EXISTS public.course_methods_method_enum;

CREATE TYPE public.course_methods_method_enum AS ENUM
    ('OnlineAsynchronous', 'OnlineSynchronous', 'FaceToFace', 'Others');

ALTER TYPE public.course_methods_method_enum
    OWNER TO postgres;

-- Type: course_schedules_day_enum

-- DROP TYPE IF EXISTS public.course_schedules_day_enum;

CREATE TYPE public.course_schedules_day_enum AS ENUM
    ('Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Intensive', 'Appointment', 'AnyTime', 'Unknown');

ALTER TYPE public.course_schedules_day_enum
    OWNER TO postgres;

-- Type: course_schedules_module_enum

-- DROP TYPE IF EXISTS public.course_schedules_module_enum;

CREATE TYPE public.course_schedules_module_enum AS ENUM
    ('SpringA', 'SpringB', 'SpringC', 'FallA', 'FallB', 'FallC', 'SummerVacation', 'SpringVacation', 'Annual', 'Unknown');

ALTER TYPE public.course_schedules_module_enum
    OWNER TO postgres;

-- SEQUENCE: public.course_methods_id_seq

-- DROP SEQUENCE IF EXISTS public.course_methods_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.course_methods_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;
    -- OWNED BY course_methods.id;

ALTER SEQUENCE public.course_methods_id_seq
    OWNER TO postgres;

-- SEQUENCE: public.course_recommended_grades_id_seq

-- DROP SEQUENCE IF EXISTS public.course_recommended_grades_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.course_recommended_grades_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;
    -- OWNED BY course_recommended_grades.id;

ALTER SEQUENCE public.course_recommended_grades_id_seq
    OWNER TO postgres;

-- SEQUENCE: public.course_schedules_id_seq

-- DROP SEQUENCE IF EXISTS public.course_schedules_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.course_schedules_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1;
    -- OWNED BY course_schedules.id;

ALTER SEQUENCE public.course_schedules_id_seq
    OWNER TO postgres;

-- Table: public.courses

-- DROP TABLE IF EXISTS public.courses;

CREATE TABLE IF NOT EXISTS public.courses
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    year smallint NOT NULL,
    code text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    instructor text COLLATE pg_catalog."default" NOT NULL,
    credit numeric NOT NULL,
    overview text COLLATE pg_catalog."default" NOT NULL,
    remarks text COLLATE pg_catalog."default" NOT NULL,
    last_update timestamp with time zone NOT NULL,
    has_parse_error boolean NOT NULL,
    is_annual boolean NOT NULL DEFAULT false,
    CONSTRAINT "PK_3f70a487cc718ad8eda4e6d58c9" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.courses
    OWNER to postgres;
-- Index: IDX_68ca51dc447bc2c03d5f1c44b8

-- DROP INDEX IF EXISTS public."IDX_68ca51dc447bc2c03d5f1c44b8";

CREATE UNIQUE INDEX IF NOT EXISTS "IDX_68ca51dc447bc2c03d5f1c44b8"
    ON public.courses USING btree
    (year ASC NULLS LAST, code COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;

-- Table: public.course_methods

-- DROP TABLE IF EXISTS public.course_methods;

CREATE TABLE IF NOT EXISTS public.course_methods
(
    id integer NOT NULL DEFAULT nextval('course_methods_id_seq'::regclass),
    method course_methods_method_enum NOT NULL,
    course_id uuid,
    CONSTRAINT "PK_412b40a7891d105d69847f77f72" PRIMARY KEY (id),
    CONSTRAINT "FK_b65806d6baf43fbfa16751d9d5c" FOREIGN KEY (course_id)
        REFERENCES public.courses (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.course_methods
    OWNER to postgres;

-- Table: public.course_recommended_grades

-- DROP TABLE IF EXISTS public.course_recommended_grades;

CREATE TABLE IF NOT EXISTS public.course_recommended_grades
(
    id integer NOT NULL DEFAULT nextval('course_recommended_grades_id_seq'::regclass),
    grade smallint NOT NULL,
    course_id uuid,
    CONSTRAINT "PK_832c34a40a52a56375e2796b3c6" PRIMARY KEY (id),
    CONSTRAINT "FK_54fbcd6eb4ed154ac8be7e7dc72" FOREIGN KEY (course_id)
        REFERENCES public.courses (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.course_recommended_grades
    OWNER to postgres;

-- Table: public.course_schedules

-- DROP TABLE IF EXISTS public.course_schedules;

CREATE TABLE IF NOT EXISTS public.course_schedules
(
    id integer NOT NULL DEFAULT nextval('course_schedules_id_seq'::regclass),
    module course_schedules_module_enum NOT NULL,
    day course_schedules_day_enum NOT NULL,
    period smallint NOT NULL,
    room text COLLATE pg_catalog."default" NOT NULL,
    course_id uuid,
    CONSTRAINT "PK_68118fc569f0c9ebb03fb79f80e" PRIMARY KEY (id),
    CONSTRAINT "FK_f0cd61820798323cb06ca91105d" FOREIGN KEY (course_id)
        REFERENCES public.courses (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.course_schedules
    OWNER to postgres;

COMMIT;
