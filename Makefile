SERVER_OUT := "bin/servidor"
CLIENT_OUT := "bin/cliente"
API_OUT := "api/api.pb.go"
API_REST_OUT := "api/api.pb.gw.go"
API_SWAG_OUT := "api/api.swagger.json"
PKG := "tracksale.prova"
SERVER_PKG_BUILD := "${PKG}/servidor"
CLIENT_PKG_BUILD := "${PKG}/cliente"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all api servidor cliente

all: servidor cliente

api/api.pb.go: api/api.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/tracksale.prova/api \
		--go_out=plugins=grpc:api \
		api/api.proto

api/api.pb.gw.go: api/api.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/tracksale.prova/api \
		--grpc-gateway_out=logtostderr=true:api \
		api/api.proto

api/api.swagger.json: api/api.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/tracksale.prova/api \
		--swagger_out=logtostderr=true:api \
		api/api.proto

api: api/api.pb.go # api/api.pb.gw.go api/api.swagger.json ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

servidor: dep api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

cliente: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

# calculapi: 
#	@go build -i -v -o 

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT) # $(API_REST_OUT) $(API_SWAG_OUT)

all: calculapi cliente servidor api


help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
