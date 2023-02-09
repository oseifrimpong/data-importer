-- public."data" definition

-- Drop table

-- DROP TABLE public."data";

CREATE TABLE public."data" (
	"version" int4 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	id uuid NOT NULL,
	unix int8 NULL,
	symbol varchar(20) NULL,
	"open" varchar(25) NULL,
	high varchar(25) NULL,
	low varchar(25) NULL,
	"close" varchar(25) NULL,
	CONSTRAINT data_pkey PRIMARY KEY (id)
);