-include .env
export

init:
	goctl api go -api wiki.api  -dir .

test:
	go test -timeout 10s -v  ./...

docker:
	goctl docker -go server.go

serve:
	go run server.go -f ./server-api.yaml