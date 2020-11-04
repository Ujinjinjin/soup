create table "user"
(
	user_id serial not null,
	username varchar(128) not null,
	email varchar(128) not null,
	first_name varchar(256) null,
	last_name varchar(256) null,
	middle_name varchar(256) null,
	primary key (user_id)
);
