GOFLAGS = -tags netgo
GITHUB_ACCOUNT = ObjectIsAdvantag
DOCKER_ACCOUNT = objectisadvantag
CONFIG=-logtostderr=true -v=5
PROGRAM=maretraite
PROJECT=github.com/$(GITHUB_ACCOUNT)/maretraite

default: all

.PHONY: test
test:
	go test $(PROJECT)/depart

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
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS)

.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS)

.PHONY: dist
dist: linux
	rm -rf dist
	mkdir dist
	cp $(PROGRAM) dist/
	cp Dockerfile dist/

.PHONY: docker-build
docker-build: dist
	cd dist; docker build -t $(DOCKER_ACCOUNT)/maretraite .

.PHONY: docker-push
docker-push: docker-build
	docker push $(DOCKER_ACCOUNT)/maretraite:latest

.PHONY: docker-pull
docker-pull:
	docker pull $(DOCKER_ACCOUNT)/maretraite:latest

.PHONY: docker-run
docker-run: docker-pull
	docker run -it $(DOCKER_ACCOUNT)/maretraite

.PHONY: archive
archive:
	git archive --format=zip HEAD > $(PROGRAM).zip

.PHONY: pkg
pkg: pkg-windows pkg-linux

.PHONY: graph
graph:
	godepgraph $(PROJECT) | dot -Tpng -o $(PROGRAM).png