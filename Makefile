run:
	@echo Running...
	@go run main.go

install:
	@echo Downloading dependencies...
	@go get
	@echo Validating dependencies...
	@go mod tidy

vendor:
	@echo Generating vendor from dependencies...
	@go mod vendor

mock:
	@echo Generating mocks..
	@echo Mocking businesses...
	@mockgen -source=core/businesses/userBusiness.go -destination=core/businesses/mock/mockUserBusiness.go -package=mock
	@echo Mocking controllers...
	@mockgen -source=core/controllers/userController.go -destination=core/controllers/mock/mockUserController.go -package=mock
	@echo Mocks generated successfully!

test:
	@echo Running tests
	@go test -v ./...
	@echo test successfully!
