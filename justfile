# list all receipts
@help:
  just -l

# run go mod vendor to fetch all dependencies for local build
@vendor:
  go mod vendor

# run all tests with ginkgo
@test:
  ginkgo run -r -cover -coverprofile=coverage.out

# build a binary executable
@build:
  go build .

# run the Monkey language interpreter in an interactive shell env
@run:
  go run main.go

# run lint
@lint:
  golangci-lint run
