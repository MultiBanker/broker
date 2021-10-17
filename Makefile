tidy:
	go mod tidy
	go mod vendor

swag-admin:
	swag init -g ./src/servers/adminhttp/resources/admin/resource.go -o ./swagger/admin/ --parseDependency --parseInternal --exclude ./src/servers/clienthttp/resources/market,./src/servers/clienthttp/resources/partner

swag-market:
	swag init -g ./src/servers/clienthttp/resources/market/resource.go -o ./swagger/market/ --parseDependency --parseInternal --exclude ./src/servers/adminhttp/resources/admin,./src/servers/clienthttp/resources/partner

swag-partner:
	swag init -g ./src/servers/clienthttp/resources/partner/resource.go -o ./swagger/partner/ --parseDependency --parseInternal --exclude ./src/servers/adminhttp/resources/admin,./src/servers/clienthttp/resources/market

swag-gen: swag-admin swag-market swag-partner

docs:
	swag init -d ./src/app -o ./swagger --parseDependency --parseInternal

mocks:
