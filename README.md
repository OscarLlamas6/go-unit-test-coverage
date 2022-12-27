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

## CircleCI

```bash

curl -X POST --header "Content-Type: application/json" --header "Circle-Token: $CIRCLE_TOKEN" -d '{
  "parameters": {
    "run-workflow-1": true,
    "run-workflow-2": false
  }
}' https://circleci.com/api/v2/project/gh/OscarLlamas6/go-unit-test-coverage/pipeline

```