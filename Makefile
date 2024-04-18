dev: go run .
	
build: go build -o ./tmp/ .
	

watch:css:
	@tailwindcss -i ./input.css -o components/styles/output.css --watch
