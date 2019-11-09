default: build

test:
	go test ./...

build:
	go build --buildmode=plugin -o tflint-ruleset-template.so main.go

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-template.so ~/.tflint.d/plugins
