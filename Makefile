install:
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
alias:
	alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOCACHE=/tmp -e  GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger: 
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
cleanmod:
	go clean --modcache