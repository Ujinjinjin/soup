create table privilege_to_subsystem
(
	privilege_to_subsystem_id serial not null,
	privilege_id int not null,
	subsystem_id int not null,
	primary key (privilege_to_subsystem_id)
);

create index i$privilege_to_subsystem$privilege_id$subsystem_id on privilege_to_subsystem (privilege_id, subsystem_id);
alter table privilege_to_subsystem add constraint fk$privilege_to_subsystem$privilege_id foreign key (privilege_id) references privilege (privilege_id);
