start-db:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

create-tables:
	docker cp ./schema/up.sql todo-db:/docker-entrypoint-initdb.d/up.sql
	docker cp ./schema/down.sql todo-db:/docker-entrypoint-initdb.d/down.sql
	docker exec -u postgres todo-db psql postgres postgres -f docker-entrypoint-initdb.d/up.sql
delete-tables:
	docker exec -u postgres todo-db psql postgres postgres -f docker-entrypoint-initdb.d/down.sql