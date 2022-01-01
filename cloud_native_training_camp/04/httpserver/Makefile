TAG = 1.0.0
REGISTRY = chkinglee

OS = linux
ARCH = amd64

build:
	@echo "building httpserver binary"
	@echo "GOOS=${OS} GOARCH=${ARCH}"
	@mkdir -p bin/${ARCH}
	CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -o bin/${ARCH} .

release: build
	@echo "building httpserver image"
	docker build -t ${REGISTRY}/httpserver:${TAG} .

push: release
	@echo "pushing httpserver image to docker hub"
	docker push ${REGISTRY}/httpserver:${TAG}