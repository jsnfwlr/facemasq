CURRENTVERSION != git describe --tags --abbrev=0
RELEASEVERSION != semver -p $(CURRENTVERSION)
ALPHAVERSION != semver -d "alpha" $(RELEASEVERSION)
BETAVERSION != semver -d "beta" $(RELEASEVERSION)
RCVERSION != semver -d "rc" $(RELEASEVERSION)
ARCH != go env GOARCH
IMAGENAME != docker info | sed -r "/Username:/!d;s|.*: (.*)|\1/facemasq|"

# TAGGING - tag commits for different release types
tag-alpha:
	@git tag $(ALPHAVERSION)
	@git push upstream --tags

tag-beta:
	@git tag $(BETAVERSION)
	@git push upstream --tags

tag-rc:
	@git tag $(RCVERSION)
	@git push upstream --tags

tag-release:
	@git tag $(RELEASEVERSION)
	@git push upstream --tags
	
# API - run vet, tests or coverage generation against the API code
api-vet:
	@cd api; go vet -tags test ./...
api-test:
	@cd api; go test ./...

api-coverage:
	@rm data/test.sqlite data/test2.sqlite; cd api; go test -v ./... -covermode=count -coverpkg=./... -coverprofile ../dist/coverage.out; go tool cover -html ../dist/coverage.out -o ../dist/coverage.html

api-test-tests:
	@cd api; go test -tags test ./...
# API - build local version
api: 
	@cd api; CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" --tags "linux sqlite_foreign_keys=1" -o ../dist/api/facemasq .

# CONTAINER - build various versions of the container
container-release:
	docker buildx build --platform linux/arm64,linux/amd64 -t $(IMAGENAME):dev -t $(IMAGENAME):$(CURRENTVERSION) --push -f docker/Dockerfile.multiarch .

container-dev:
	docker buildx build --platform linux/arm64,linux/amd64 -t $(IMAGENAME):dev --push -f docker/Dockerfile.multiarch .

container-basic:
	docker build -t $(IMAGENAME):basic-$(ARCH) -t $(IMAGENAME):basic-$(ARCH)-$(CURRENTVERSION) -f docker/Dockerfile .
# WEB - testing
web-test:
	@echo "No Testing Yet"
web-coverage:
	@echo "No Testing Yet"
# WEB - build the web UI
web:
	@cd web; pnpm build

test-names:
	@echo $(IMAGENAME)-$(ARCH):$(CURRENTVERSION)

all:
	web-test
	api-test
	api
	web

