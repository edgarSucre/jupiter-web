css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

sqlite:
	sqlite3 db/jupyter.db

sqlc:
	sqlc generate