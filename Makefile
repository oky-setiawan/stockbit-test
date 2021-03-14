tidy:
	@echo ">>> tidying..."
	@go mod tidy
	@go mod vendor

remock:
    @mockery --dir=./internal/repository --name=LogRepository --output=./mocks/mock_internal/mock_repository --outpkg=mock_repository
    @mockery --dir=./internal/domain --name=OMDBDomain --output=./mocks/mock_internal/mock_domain --outpkg=mock_domain
    @mockery --dir=./internal/usecase --name=MovieUsecase --output=./mocks/mock_internal/mock_usecase --outpkg=mock_usecase

run:
	@echo ">>> building services"
	@go run ./cmd/main.go