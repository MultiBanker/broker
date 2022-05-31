tidy:
	go mod tidy
	go mod vendor

docs:
	swag fmt -d src/app
	swag init --parseDependency --parseDepth 5 -g servers/adminhttp/router.go -o swagger/admin -d src --exclude ./src/clienthttp/
	swag init --parseDependency --parseDepth 5 -g servers/clienthttp/router.go -o swagger/users -d src --exclude ./src/adminhttp/

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
