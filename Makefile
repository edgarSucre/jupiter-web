css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

templ:
	templ generate

sqlite:
	sqlite3 db/jupyter.db

sqlc:
	sqlc generate