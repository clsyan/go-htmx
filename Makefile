migration up:
	migrate -path ./migrations/ -database "mysql://root:hbx@tcp(172.17.0.1:3306)/go-htmx" up

back:
	migrate -path ./migrations/ -database "mysql://root:hbx@tcp(172.17.0.1:3306)/go-htmx" force $(to)
