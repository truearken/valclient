run:
	GOOS=windows go build -o valclient.exe cmd/client/main.go && ./valclient.exe

run-test-cover:
	GOOS=windows go test -v -coverprofile=/tmp/coverage.out -coverpkg=github.com/truearken/valclient/valclient ./...

run-test:
	GOOS=windows go test -v ./...
