run:
	GOOS=windows go build -o bin/valclient.exe cmd/client/main.go && bin/valclient.exe

run-test-cover:
	GOOS=windows go test -v -coverprofile=/tmp/coverage.out -coverpkg=github.com/truearken/valclient/valclient ./...

run-test:
	GOOS=windows go test -v ./...
