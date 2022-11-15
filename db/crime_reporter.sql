-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 1.0.0-beta
-- PostgreSQL version: 15.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: crime_report | type: DATABASE --
DROP DATABASE IF EXISTS crime_report;
CREATE DATABASE crime_report;
-- ddl-end --


-- object: role | type: SCHEMA --
DROP SCHEMA IF EXISTS role CASCADE;
CREATE SCHEMA role;
-- ddl-end --
ALTER SCHEMA role OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,role;
-- ddl-end --

-- object: public."user" | type: TABLE --
DROP TABLE IF EXISTS public."user" CASCADE;
CREATE TABLE public."user" (
	id uuid NOT NULL,
	username varchar(100),
	email varchar(100) NOT NULL,
	first_name varchar(200) NOT NULL,
	last_name varchar(200) NOT NULL,
	other_name varchar(200),
	user_type varchar(100) NOT NULL,
	mobile varchar(20) NOT NULL,
	about varchar(200),
	address varchar(220),
	status varchar(20) NOT NULL,
	role uuid NOT NULL,
	token varchar(220),
	active boolean NOT NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	CONSTRAINT pk_user_id PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public."user" OWNER TO postgres;
-- ddl-end --

-- object: public.crime | type: TABLE --
DROP TABLE IF EXISTS public.crime CASCADE;
CREATE TABLE public.crime (
	id uuid NOT NULL,
	title varchar(220) NOT NULL,
	description text,
	type varchar(50) NOT NULL,
	media varchar(50),
	location varchar(150) NOT NULL,
	CONSTRAINT pk_crime_id PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.crime OWNER TO postgres;
-- ddl-end --

-- object: public.role | type: TABLE --
DROP TABLE IF EXISTS public.role CASCADE;
CREATE TABLE public.role (
	id uuid NOT NULL,
	name varchar(100) NOT NULL,
	description text,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	CONSTRAINT pk_role_id PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.role OWNER TO postgres;
-- ddl-end --

-- object: public.crime_case | type: TABLE --
DROP TABLE IF EXISTS public.crime_case CASCADE;
CREATE TABLE public.crime_case (
	id uuid NOT NULL,
	crime uuid NOT NULL,
	reporter uuid NOT NULL,
	handler uuid,
	status varchar(50) NOT NULL,
	description text,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	CONSTRAINT pk_crime_case_id PRIMARY KEY (id)
);
-- ddl-end --

-- object: fk_role_id | type: CONSTRAINT --
ALTER TABLE public."user" DROP CONSTRAINT IF EXISTS fk_role_id CASCADE;
ALTER TABLE public."user" ADD CONSTRAINT fk_role_id FOREIGN KEY (role)
REFERENCES public.role (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_reporter_id | type: CONSTRAINT --
ALTER TABLE public.crime_case DROP CONSTRAINT IF EXISTS fk_reporter_id CASCADE;
ALTER TABLE public.crime_case ADD CONSTRAINT fk_reporter_id FOREIGN KEY (reporter)
REFERENCES public."user" (id) MATCH SIMPLE
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_handler_id | type: CONSTRAINT --
ALTER TABLE public.crime_case DROP CONSTRAINT IF EXISTS fk_handler_id CASCADE;
ALTER TABLE public.crime_case ADD CONSTRAINT fk_handler_id FOREIGN KEY (handler)
REFERENCES public."user" (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_crime_id | type: CONSTRAINT --
ALTER TABLE public.crime_case DROP CONSTRAINT IF EXISTS fk_crime_id CASCADE;
ALTER TABLE public.crime_case ADD CONSTRAINT fk_crime_id FOREIGN KEY (crime)
REFERENCES public.crime (id) MATCH SIMPLE
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --


