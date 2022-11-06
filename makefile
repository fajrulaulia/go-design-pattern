test:
	@gotestsum --format testname  -- -coverprofile=cover.out ./... -v 