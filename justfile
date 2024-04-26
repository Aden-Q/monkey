# list all receipts
@help:
  just -l

# run go mod vendor to fetch all dependencies for local build
@vendor:
  go mod vendor

# run all tests with ginkgo
@test:
  ginkgo run -r -cover -coverprofile=coverage.out

# run lint
@lint:
  golangci-lint run

# generate RESTful API documentation with Swagger 2.0
@swag-init:
  swag init
