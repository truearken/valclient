run:
	GOOS=windows go build -o valclient.exe cmd/client/main.go && ./valclient.exe
