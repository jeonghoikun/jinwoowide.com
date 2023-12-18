.PHONY:
	tailwindc

tailwindc:
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/main.css --watch
