create schema social;

create table social.profiles(
	id uuid primary key default gen_random_uuid(),
    version bigint not null default 1 /* оптимистичная блокировка */,
	email varchar(320) not null unique,
    phone_number varchar(15) check (
        phone_number ~ '^+[0-9]+$'
        and
        char_length(phone_number) between 10 and 15
    ),
	username varchar(30) not null unique check (char_length(username) between 6 and 30),
	password varchar(250) not null,
	name varchar(50) not null,
	birthday date not null,
	gender smallint not null,
	longitude real,
	latitude real,
	like_ttl smallint not null default 7 check (like_ttl between 1 and 30)
);


create table social.preferences(
	profile_id uuid primary key references social.profiles(id) on delete cascade,
    age_from smallint default 18 check (age_from >= 18),
    age_to smallint default 99 check (age_to <= 150),
	raduis int default 50 /* km */
);
