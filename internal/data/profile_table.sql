create table profiles(
	id uuid primary key default gen_random_uuid()
	email varchar(320) unique /* TODO: create index */
	username varchar(30) unique
	password varchar(250)
	name varchar(50)
	birtday date
	gender varchar(15)
	longitude real 
	latitude real
)
