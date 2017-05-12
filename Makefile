GO_LDFLAGS=-ldflags " -w"
build:
	@GOOS=linux go build ${GO_LDFLAGS} -o fserver
image:build
	@docker build -t hub.faas.pro/fserver:latest .	