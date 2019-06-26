\c archive

DROP SCHEMA public;

CREATE SCHEMA public;

COMMENT ON SCHEMA public IS 'standard public schema';

DROP TABLE IF EXISTS public.app_config;

CREATE TABLE public.app_config (
	id int4 NOT NULL,
	configuration json NULL,
	"role" varchar(20) NULL,
	CONSTRAINT app_config_pk PRIMARY KEY (id)
);

INSERT INTO public.app_config (id,configuration,"role") VALUES 
(1,'{
	"routes": [
	{"order":1,"to": "/", "icon": "home", "text": "Home","dropdown":false},
	{"order":2,"to": "/upload","icon": "backup", "text": "Upload data","dropdown":false},
	{"order":4,"to": "/query","icon": "help","text": "Query","dropdown":false},
	{"order":3,"to": "/datasets", "icon": "view_column", "text": "My Files","dropdown":true,"dropdown_data":"datasets"}
	]
}','user'),
(2, '{
	"routes": [
	{"order":1,"to": "/", "icon": "home", "text": "Home","dropdown":false},
	{"order":2,"to": "/upload","icon": "backup", "text": "Upload data","dropdown":false},
	{"order":4,"to": "/query","icon": "help","text": "Query","dropdown":false},
	{"order":3,"to": "/datasets", "icon": "view_column", "text": "My Files","dropdown":true,"dropdown_data":"datasets"},
	{"order":5,"to": "/accounts", "icon": "account_circle", "text": "Create account","dropdown":false}
	]
}', 'admin')
;

-- Drop table

DROP TABLE IF EXISTS public.minio;

CREATE TABLE public.minio (
	user_id varchar(50) NULL,
	user_bucket_name varchar(60) NULL,
	bucket_name varchar(100) NULL,
	bucket_secret varchar(100) NULL,
	bucket_access_key varchar(100) NULL
);

-- Drop table

DROP TABLE IF EXISTS public.users;

CREATE TABLE public.users (
	user_id varchar(50) NOT NULL,
	username varchar(100) NOT NULL,
	user_pass varchar(100) NOT NULL
);
CREATE UNIQUE INDEX users_userid_idx ON public.users USING btree (user_id);
CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);

INSERT INTO public.users (user_id,username,user_pass) VALUES 
('f5547634-2c08-4780-bda3-e446771c6a2c','admin','$2a$10$wZR.HqIdG1PUJUkRLGoXB.ZSDV3i/pDJ0QmCZLiXoHJFiRl1r6Cz2')
;

\c policydb

DROP SCHEMA IF EXISTS policies;

CREATE SCHEMA policies;

COMMENT ON SCHEMA policies IS 'schema for authorization policies';

DROP TABLE IF EXISTS policies.policies;

CREATE TABLE policies.policies (
	ptype varchar(10) NULL,
	sub varchar(256) NULL,
	obj varchar(256) NULL,
	"oid" varchar(256) NULL,
	act varchar(256) NULL,
	CONSTRAINT policies_un UNIQUE (ptype, sub, obj, oid, act)
);

INSERT INTO policies.policies (ptype,sub,obj,"oid",act) VALUES 
('p','user','/user/*','1','(.*)')
,('p','user','/datasets/*','1','(.*)')
,('g','f5547634-2c08-4780-bda3-e446771c6a2c','admin',NULL,NULL)
,('g','f5547634-2c08-4780-bda3-e446771c6a2c','user',NULL,NULL)
,('p','admin','/*','1','(.*)')
,('p','f5547634-2c08-4780-bda3-e446771c6a2c','create-account','*','(.*)')
;
