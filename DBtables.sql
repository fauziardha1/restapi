-- Database Name : SOLID
-- Database using PostgreSQL

-- public.measurements definition

-- Drop table

-- DROP TABLE public.measurements;

CREATE TABLE public.measurements (
	id serial4 NOT NULL,
	project_id int4 NOT NULL,
	file_name varchar NOT NULL,
	srp_val float8 NULL DEFAULT 0.0,
	ocp_val float8 NULL DEFAULT 0.0,
	lsp_val float8 NULL DEFAULT 0.0,
	isp_val float8 NULL DEFAULT 0.0,
	dip_val float8 NULL DEFAULT 0.0,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT measurements_pkey PRIMARY KEY (id)
);


-- public.projects definition

-- Drop table

-- DROP TABLE public.projects;

CREATE TABLE public.projects (
	id int4 NOT NULL,
	user_id int4 NOT NULL,
	"name" varchar NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	CONSTRAINT projects_pkey PRIMARY KEY (id)
);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id int4 NOT NULL,
	username varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	"password" varchar(100) NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);