	echo "tw" && npx tailwindcss -i ./input.css -o ./components/styles/output.css &&
	echo "templ" && templ generate && 
	echo "go" && go run main.go
