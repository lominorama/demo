front-build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o build/front main.go

front-docker-build:
	docker build . -t rvazquez/gitops-demo-front:$(VERSION)

front-docker-push:
	docker push rvazquez/gitops-demo-front:$(VERSION)

front-image: front-build-linux front-docker-build front-docker-push
