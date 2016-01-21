GOFLAGS = -tags netgo
GITHUB_ACCOUNT = ObjectIsAdvantag
DOCKER_ACCOUNT = objectisadvantag
CONFIG=-logtostderr=true -v=5
PROGRAM=retraite
PROJECT=github.com/$(GITHUB_ACCOUNT)/retraite

default: all

.PHONY:devenv
devenv:
	go get github.com/Sirupsen/logrus

.PHONY: all
all : devenv clean build run

.PHONY: build
build: clean
	go build $(GOFLAGS)

.PHONY: debug
debug:
	godebug build $(GOFLAGS) -instrument $(PROJECT)/depart
	./$(PROGRAM).debug $(CONFIG)

.PHONY: run
run: build
	./$(PROGRAM).exe -port 8080 $(CONFIG)

.PHONY: clean
clean:
	rm -f $(PROGRAM) $(PROGRAM).exe $(PROGRAM).zip $(PROGRAM).debug

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) $(PROGRAM).go

.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) $(PROGRAM).go

.PHONY: dist
dist: linux
	rm -rf dist
	mkdir dist
	cp $(PROGRAM) dist/
	mkdir dist/logs
	cp Dockerfile dist/

.PHONY: docker
docker: dist
	cd dist; docker build -t $(DOCKER_ACCOUNT)/bilanretraite .

.PHONY: archive
archive:
	git archive --format=zip HEAD > $(PROGRAM).zip

.PHONY: pkg
pkg: pkg-windows pkg-linux

.PHONY: graph
graph:
	godepgraph $(PROJECT) | dot -Tpng -o $(PROGRAM).png