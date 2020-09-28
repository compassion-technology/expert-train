CREATE TABLE "space".messages (
	id serial NOT NULL,
	"from" varchar(8) NULL,
	"to" varchar(8) NULL,
	group_id int8 NULL,
	"text" text NULL,
	content_url text NULL,
	created timestamp(0) NULL,
	CONSTRAINT messages_pk PRIMARY KEY (id)
)
TABLESPACE pg_default
;
