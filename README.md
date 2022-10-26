# GoRedis
This is a microserves project using goswagger, mysql, and redis. All sql tables and procedures are within the MMO sql project.

# Build Swagger
```alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOCACHE=/tmp -e  GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'```

# Run Go
`go build`

`go run .`

# Api runs at
`http://localhost:9090` 
