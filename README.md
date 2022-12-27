# Go unit testing with Code Coverage

```makefile
go-test:
	go test .\test\...

test-verbose:
	go test -v .\test\...

test-coverage:
	go test .\test\... -coverpkg ./...

test-coverage-html:
	go test ./test/... -coverpkg ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```
