run: build
	
templ:
	@templ generate -watch -proxy=http://localhost:3000

tailwind:
	@tailwindcss -i components/assets/styles/input.css -o components/assets/styles/output.css --watch
