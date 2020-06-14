BINARY := webui
.DEFAULT_GOAL := webui
ORG := geocodes
VERSION :=`cat VERSION`
DOCKERNAME :=  p418webui


webui:
	cd cmd/$(BINARY) ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)

docker:
	docker build  --tag="$(ORG)/$(BINARY)"  --file=./build/Dockerfile . ; \
	docker tag $(ORG)/$(BINARY) $(ORG)/$(DOCKERNAME):latest

publish: docker
	docker push $(ORG)/$(DOCKERNAME):latest ; \
        docker tag $(ORG)/$(DOCKERNAME) $(ORG)/$(DOCKERNAME):$(VERSION) ; \
	docker push $(ORG)/$(DOCKERNAME):$(VERSION)
