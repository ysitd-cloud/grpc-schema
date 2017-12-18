SOURCE_FILE := $(shell find . -path ./node_modules -prune -o -type f -name '*.proto')
GO_TARGET_FILE := $(SOURCE_FILE:.proto=.pb.go)
NODE_TARGET_FILE := $(SOURCE_FILE:.proto=_pb.js)
PYTHON_TARGET_FILE := $(SOURCE_FILE:.proto=_pb2.py)
RUBY_TARGET_FILE := $(SOURCE_FILE:.proto=_pb.rb)

all: go js ruby python
	@echo Compile Done

.PHONY: all clean js python ruby go

js: $(NODE_TARGET_FILE)
	@echo Finish build for all js modules

%_pb.js: %.proto
	yarn generate:file $<

go: $(GO_TARGET_FILE)
	@echo Finish build for all go modules

%.pb.go: %.proto
	protoc -I $(GOPATH)/src/github.com/ysitd-cloud/grpc-schema $< --go_out=plugins=grpc:$(GOPATH)/src

python: $(PYTHON_TARGET_FILE)
	@echo Finish build for all python modules

%_pb2.py: %.proto
	python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. $<

ruby: $(RUBY_TARGET_FILE)
	@echo Finish build for all ruby modules

%_pb.rb: %.proto
	protoc -I . --ruby_out=. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` $<

clean:
	rm -f $(shell find . -type f -name '*.go')
	rm -f $(shell find . -type f -name '*.js')
	rm -f $(shell find . -type f -name '*.rb')
	rm -f $(shell find . -type f -name '*.py')
