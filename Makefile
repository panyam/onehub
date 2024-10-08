
all: build

build: down copylinks dockerbuild resymlink

copylinks:
	rm -Rf locallinks/*
	cp -r ../dbsync locallinks/
	cp -r ../goutils locallinks/

dockerbuild:
	BUILDKIT_PROGRESS=plain docker compose build --no-cache

resymlink:
	rm -Rf locallinks/*
	cd locallinks && ln -s ../../dbsync
	cd locallinks && ln -s ../../goutils

up: ensurenetworks
	docker compose -f docker-compose.yml down
	BUILDKIT_PROGRESS=plain docker compose -f docker-compose.yml up -d

logs:
	docker compose -f docker-compose.yml logs -f

# Bring everything down
down:
	docker compose -f docker-compose.yml down --remove-orphans
	docker compose -f db-docker-compose.yml down --remove-orphans

# Bring up DB - only brings down DB containers from before
updb: dbdirs ensurenetworks
	BUILDKIT_PROGRESS=plain docker compose -f db-docker-compose.yml down --remove-orphans
	BUILDKIT_PROGRESS=plain docker compose -f db-docker-compose.yml up -d

dblogs:
	docker compose -f db-docker-compose.yml logs -f --tail 100

ensurenetworks:
	-docker network create onehubnetwork

dbdirs:
	mkdir -p ./data/pgdata ./data/typesensedata

#### Deprecated - only used in earlier versions before buf

# Some vars to detemrine go locations etc
GOROOT=$(which go)
GOPATH=$(HOME)/go
GOBIN=$(GOPATH)/bin

# Evaluates the abs path of the directory where this Makefile resides
SRC_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# Where the protos exist
PROTO_DIR:=$(SRC_DIR)/protos

# where we want to generate server stubs, clients etc
OUT_DIR:=$(SRC_DIR)/gen/go

oldway: createdirs printenv goprotos gwprotos openapiv2 cleanvendors

goprotos:
	echo "Generating GO bindings"
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative          	\
       --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative		\
       --proto_path=$(PROTO_DIR) 																			\
      $(PROTO_DIR)/onehub/v1/*.proto

gwprotos:
	echo "Generating gRPC Gateway bindings and OpenAPI spec"
	protoc -I . --grpc-gateway_out $(OUT_DIR)               \
		--grpc-gateway_opt logtostderr=true                   \
		--grpc-gateway_opt paths=source_relative              \
		--grpc-gateway_opt generate_unbound_methods=true      \
    --proto_path=$(PROTO_DIR) 																				\
      $(PROTO_DIR)/onehub/v1/*.proto

openapiv2:
	echo "Generating OpenAPI specs"
	protoc -I . --openapiv2_out $(SRC_DIR)/gen/openapiv2      \
    --openapiv2_opt logtostderr=true                    \
    --openapiv2_opt generate_unbound_methods=true           \
    --openapiv2_opt allow_merge=true                    \
    --openapiv2_opt merge_file_name=allservices             \
    --proto_path=$(PROTO_DIR) 															\
      $(PROTO_DIR)/onehub/v1/*.proto

printenv:
	@echo MAKEFILE_LIST=$(MAKEFILE_LIST)
	@echo SRC_DIR=$(SRC_DIR)
	@echo PROTO_DIR=$(PROTO_DIR)
	@echo OUT_DIR=$(OUT_DIR)
	@echo GOROOT=$(GOROOT)
	@echo GOPATH=$(GOPATH)
	@echo GOBIN=$(GOBIN)

createdirs:
	rm -Rf $(OUT_DIR)
	mkdir -p $(OUT_DIR)
	mkdir -p $(SRC_DIR)/gen/openapiv2
	cd $(PROTO_DIR) && (																							\
 		if [ ! -d google ]; then ln -s $(SRC_DIR)/vendors/google . ; fi	\
	)

cleanvendors:
	rm -f $(PROTO_DIR)/google
