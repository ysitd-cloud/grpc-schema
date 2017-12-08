SOURCE_FILE := $(shell find . -type f -name '*.proto')
GO_TARGET_FILE := $(SOURCE_FILE:.proto=.pb.go)
NODE_TARGET_FILE := $(SOURCE_FILE:.proto=_pb.js)
PYTHON_TARGET_FILE := $(SOURCE_FILE:.proto=_pb2.py)

all: $(GO_TARGET_FILE) $(NODE_TARGET_FILE)
	@echo Compile Done

.PHONY: all %.pb.go clean

%.pb.go: %.proto
	protoc -I $(GOPATH)/src/github.com/ysitd-cloud/grpc-schema $< --go_out=plugins=grpc:$(GOPATH)/src
	grpc_tools_node_protoc --js_out=import_style=commonjs,binary:. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` $<
	python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. $<

clean:
	rm -f $(GO_TARGET_FILE) $(NODE_TARGET_FILE) $(shell find . -type f -name '*.js')
