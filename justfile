# list all receipts
@help:
  just -l

# run go mod vendor to fetch all dependencies for local build
@vendor:
  go mod vendor

# run all unit tests with ginkgo
@test:
  ginkgo run -r -race -cover -coverprofile=coverage.out

# build a binary executable
@build:
  go build -gcflags='-m=2' .

# docker build
@docker-build:
  docker build -t monkey .

# docker run
@docker-run:
  docker run -it --rm --name monkey zecheng/monkey

# run the Monkey language interpreter in interactive mode
@run:
  go run main.go

# run lint
@lint:
  golangci-lint run
