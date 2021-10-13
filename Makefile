tidy:
	go mod tidy
	go mod vendor

swag-gen:
#	swag init -d ./src/app -o ./src/swagger --parseDependency
	swag init -g ./src/servers/http/resources/admin/resource.go -o ./swagger/admin/ --exclude ./src/servers/http/resources/market,./src/servers/http/resources/partner &&
	swag init -g ./src/servers/http/resources/market/resource.go -o ./swagger/market/ --exclude ./src/servers/http/resources/admin,./src/servers/http/resources/partner &&
	swag init -g ./src/servers/http/resources/partner/resource.go -o ./swagger/partner/ --exclude ./src/servers/http/resources/admin,./src/servers/http/resources/market

docs:
	swag init -d ./src/app -o ./src/swagger --parseDependency