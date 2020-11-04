﻿create table privilege
(
	privilege_id serial not null,
	code varchar(128) not null,
	name varchar(128) not null
);

alter table privilege add constraint pk$privilege primary key (privilege_id);

create index i$privilege$code on privilege (code);