swagger: 
	GO111MODULE=off swagger generate spec -o ./swagger.yaml
cleanmod:
	go clean --modcache