
/****************************************
************** Accounts *****************
****************************************/
CREATE TABLE IF NOT EXISTS accounts
(
  id bigserial NOT NULL,
  version bigint,
  created bigint,
  updated bigint,
  name text,
  description text,
  email text,
  CONSTRAINT accounts_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE accounts
  OWNER TO thingbricks;


/****************************************
********** Account Groups ***************
****************************************/
CREATE TABLE IF NOT EXISTS account_groups
(
  id bigserial NOT NULL,
  version bigint,
  created bigint,
  updated bigint,
  name text,
  default_group boolean,
  account_id bigint,
  CONSTRAINT account_groups_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE account_groups
  OWNER TO thingbricks;

/****************************************
*************** Logins ******************
****************************************/
CREATE TABLE IF NOT EXISTS logins
(
  id bigserial NOT NULL,
  version bigint,
  created bigint,
  updated bigint,
  email text,
  first_name text,
  last_name text,
  gender text,
  picture text,
  access_token text,
  network text,
  account_group_id bigint,
  CONSTRAINT logins_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE logins
  OWNER TO thingbricks;

/****************************************
************** API Keys *****************
****************************************/
CREATE TABLE IF NOT EXISTS api_keys
(
  id bigserial NOT NULL,
  version bigint,
  created bigint,
  updated bigint,
  access_key text,
  label text,
  active boolean,
  hashed bytea,
  login_id bigint,
  CONSTRAINT key_pairs_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE key_pairs
  OWNER TO thingbricks;

/****************************************
************** Projects *****************
****************************************/

CREATE TABLE projects
(
  id bigserial NOT NULL,
  version bigint,
  created bigint,
  updated bigint,
  name text,
  description text,
  account_id bigint,
  CONSTRAINT projects_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE projects
  OWNER TO thingbricks;


/****************************************
********* Data Stream Groups ************
****************************************/

CREATE TABLE data_stream_groups
(
  id bigserial NOT NULL,
  version bigint,
  created bigint,
  updated bigint,
  name text,
  description text,
  project_id bigint,
  account_id bigint,
  CONSTRAINT data_stream_groups_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE data_stream_groups
  OWNER TO thingbricks;

/****************************************
************ Data Streams ***************
****************************************/

CREATE TABLE data_streams
(
  id bigserial NOT NULL,
  version bigint,
  created bigint,
  updated bigint,
  name text,
  description text,
  project_id bigint,
  account_id bigint,
  data_stream_group_id bigint,
  CONSTRAINT data_streams_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE data_streams
  OWNER TO thingbricks;