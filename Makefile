DIR := judge
SOURCE_FILE := $(wildcard $(DIR)/*.proto $(DIR)/**/*.proto)
TARGET_FILE := $(SOURCE_FILE:.proto=.pb.go)

all: $(TARGET_FILE)
	@echo Compile Done

.PHONY: all %.pb.go clean

%.pb.go: %.proto
	protoc -I $(GOPATH)/src/github.com/ysitd-cloud/grpc-schema $< --go_out=plugins=grpc:$(GOPATH)/src

clean:
	rm $(TARGET_FILE)
