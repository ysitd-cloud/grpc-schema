DIR := judge
SOURCE_FILE := $(wildcard $(DIR)/*.proto)
TARGET_FILE := $(SOURCE_FILE:.proto=.pb.go)

all: $(TARGET_FILE)
	@echo Compile Done

.PHONY: all %.pb.go clean

%.pb.go: %.proto
	protoc -I . $< --go_out=plugins=grpc:.

clean:
	rm $(TARGET_FILE)
