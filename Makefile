BINARY := webui
.DEFAULT_GOAL := webui
ORG := geocodes

webui:
	cd cmd/$(BINARY) ; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)

docker:
	docker build  --tag="$(ORG)/$(BINARY)"  --file=./build/Dockerfile . ; \
	docker tag $(ORG)/$(BINARY) $(ORG)/$(BINARY):latest

publish: docker
	docker push $(ORG)/gleaner:$(VERSION) ; \
	docker push $(ORG)/gleaner:latest
