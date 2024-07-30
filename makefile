## help: prints this help message
.PHONY: help
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: run for running app on local
.PHONY: run
run:
	@go run cmd/main.go

.PHONY: mocks
mocks:
	mockgen -source internal/usecase/user_usecase.go -destination internal/mocks/user_usecase_mock.go -package=mocks
	mockgen -source internal/repositories/user_repository.go -destination internal/mocks/user_repository_mock.go -package=mocks
	mockgen -source internal/repositories/mailer_repository.go -destination internal/mocks/mailer_repository_mock.go -package=mocks

mock-httpclient:
	mockgen -source pkg/httpclient/httpclient.go -destination pkg/httpclient/mock/httpclient_mock.go -package=mock_httpclient

mock-database:
	mockgen -source internal/infastructure/database/database.go -destination internal/infastructure/database/mock/database_mock.go -package=mock_database
mock-datasource:
	mockgen -source pkg/datasource/gorm_wrapper.go -destination pkg/datasource/mock/gorm_wrapper_mock.go -package=mock_datasource

.PHONY: tests
unit-test:
	@go clean -testcache
	@go test --race -count=1 -cpu=1 -parallel=1 -timeout=90s -failfast -vet= -cover -covermode=atomic -coverprofile=./.coverage/unit.out ./...