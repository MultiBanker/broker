tidy:
	go mod tidy
	go mod vendor

swag-gen:
	swag init -d ./src/app -o ./src/swagger --parseDependency