install:
	@echo "building toolbox ..."
	@go build -o bin/toolbox

test:
	@./bin/toolbox $*