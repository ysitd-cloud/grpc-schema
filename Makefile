SOURCE_FILE := $(shell find . -path ./node_modules -prune -o -type f -name '*.proto')
GO_TARGET_FILE := $(SOURCE_FILE:.proto=.pb.go)
NODE_TARGET_FILE := $(SOURCE_FILE:.proto=_pb.js)

all: go js
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

clean:
	rm -f $(shell find . -type f -name '*.go')
	rm -f $(shell find . -type f -name '*.js')
