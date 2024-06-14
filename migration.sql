create table todos (
	id SERIAL PRIMARY key,
	created_at timestamp,
	content varchar,
	done bool
)