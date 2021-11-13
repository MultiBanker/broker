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
	mockgen --source=src/database/repository/partner.go --destination=src/database/mock_repository/partner.go
	mockgen --source=src/database/repository/market.go --destination=src/database/mock_repository/market.go
	mockgen --source=src/database/repository/offer.go --destination=src/database/mock_repository/offer.go
	mockgen --source=src/database/repository/order.go --destination=src/database/mock_repository/order.go
	mockgen --source=src/database/repository/sequence.go --destination=src/database/mock_repository/sequence.go
	mockgen --source=src/database/repository/repository.go --destination=src/database/mock_repository/repository.go
	mockgen --source=src/database/drivers/datastore.go --destination=src/database/mock_drivers/datastore.go
	mockgen --source=src/manager/manager.go --destination=src/mock_manager/manager.go
	mockgen --source=src/manager/auth/auth.go --destination=src/mock_manager/mock_auth/auth.go
	mockgen --source=src/manager/market/market.go --destination=src/mock_manager/mock_market_manager/market.go
	mockgen --source=src/manager/offer/offers.go --destination=src/mock_manager/mock_offer_manager/offers.go
	mockgen --source=src/manager/order/order.go --destination=src/mock_manager/mock_order_manager/order.go
	mockgen --source=src/manager/partner/partner.go --destination=src/mock_manager/mock_partner_manager/partner.go
	mockgen --source=src/manager/loan/loan-program.go --destination=src/mock_manager/mock_loan_manager/loan-program.go
