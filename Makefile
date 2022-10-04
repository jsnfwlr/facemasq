CURRENTVERSION != git describe --tags --abbrev=0
RELEASEVERSION != semver -p $(CURRENTVERSION)
ALPHAVERSION != semver -d "alpha" $(RELEASEVERSION)
BETAVERSION != semver -d "beta" $(RELEASEVERSION)
RCVERSION != semver -d "rc" $(RELEASEVERSION)
ARCH != go env GOARCH
IMAGENAME != docker info | sed -r "/Username:/!d;s|.*: (.*)|\1/facemasq|"

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
	
api-test:
	@cd api; go test ./...

api-generate-coverage:
	@rm data/test.sqlite data/test2.sqlite; cd api; go test -v ./... -covermode=count -coverpkg=./... -coverprofile ../dist/coverage.out; go tool cover -html ../dist/coverage.out -o ../dist/coverage.html

api-test-tests:
	@cd api; go test -tags test ./...

# test-api-integration:
# 	cd api; go test --tags integration ./...

container-release-build:
	docker buildx build --platform linux/arm64,linux/amd64 -t $(IMAGENAME):dev -t $(IMAGENAME):$(CURRENTVERSION) --push -f docker/Dockerfile.multiarch .

container-dev-build:
	docker buildx build --platform linux/arm64,linux/amd64 -t $(IMAGENAME):dev --push -f docker/Dockerfile.multiarch .

container-basic-build:
	docker build -t $(IMAGENAME):basic-$(ARCH) -t $(IMAGENAME):basic-$(ARCH)-$(CURRENTVERSION) -f docker/Dockerfile .

web-build:
	@cd web; pnpm build

test-names:
	@echo $(IMAGENAME)-$(ARCH):$(CURRENTVERSION)

