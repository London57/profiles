create table profiles(
	id uuid primary key default gen_random_uuid()
	name varchar(50)
	Age smallint
	Gender varchar(15)
	Longitude real 
	Latitude real
)
