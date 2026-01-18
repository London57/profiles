create table profiles(
	id uuid primary key default gen_random_uuid()
	name varchar(50)
	age smallint
	gender varchar(15)
	longitude real 
	latitude real
)
