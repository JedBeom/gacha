drop table sessions;
drop table histories;
drop table users;

create table users (
	id serial primary key,
	uuid varchar(64) not null unique,
	name varchar(255) not null,
	student_id char(5) not null unique,
	email varchar(255) not null unique,
	password varchar(255) not null unique,
	is_admin boolean default false,
	coin integer default 0,
	created_at timestamp not null
);

create table sessions (
	id serial primary key,
	uuid varchar(64) not null unique,
	email varchar(255),
	user_id integer references users(id),
	created_at timestamp not null
);
