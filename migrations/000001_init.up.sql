create schema social;

create table social.profiles(
	id uuid primary key default gen_random_uuid()
    version bigint not null default 1 /* оптимистичная блокировка */
	email varchar(320) unique /* TODO: создать индекс */
    phone_number varchar(15) check (
        phone_number ~ '^+[0-9]+$'
        and
        char_length(phone_number) between 10 and 15
    )
	username varchar(30) unique check (char_length(username) between 6 and 30)
	password varchar(250)
	name varchar(50) check (char_length(name) )
	birtday date
	gender smallint
	longitude real 
	latitude real
);


create table social.preferences(
	profile_id uuid not null references social.profiles(id) on delete cascade
	birthday date
	longitude real
	latitude real
);

