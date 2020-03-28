.check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

.swagger: .check_install
	GO111MODULE=off swagger generate spec -o ./specs/swagger.json --scan-models

krait: .swagger
	go install -ldflags "-X 'github.com/hutsharing/krait/cmd.Version=0.0.9'"
