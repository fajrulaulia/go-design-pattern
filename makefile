test:
	@gotestsum --format testname  -- -coverprofile=coverage.out ./... -v 
	@go tool cover -html=coverage.out
