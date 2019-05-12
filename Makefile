.PHONY: all
all:
	make parser
	go build ./...

.PHONY: example
example:
	go run . '1+1*2'

.PHONY: docker-build
docker-build:
	docker build -t zakuro/antex-antlr -f antlr.dockerfile .

.PHONY: parser
parser:
	rm parser -rf
	docker run -it -v $(PWD):/usr/local/antex zakuro/antex-antlr
	sudo chown `id -u`:`id -g` parser -R
