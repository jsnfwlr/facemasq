CURRENTVERSION != git describe --tags --abbrev=0
RELEASEVERSION != semver -p $(CURRENTVERSION)
ALPHAVERSION != semver -d "alpha" $(RELEASEVERSION)
BETAVERSION != semver -d "beta" $(RELEASEVERSION)
RCVERSION != semver -d "rc" $(RELEASEVERSION)
ARCH != go env GOARCH

test-api:
	cd api; go test ./...

test-api-coverage:
	rm data/test.sqlite data/test2.sqlite; cd api; go test -v ./... -covermode=count -coverpkg=./... -coverprofile ../dist/coverage.out; go tool cover -html ../dist/coverage.out -o ../dist/coverage.html

# test-api-integration:
# 	cd api; go test --tags integration ./...

tag-alpha:
	git tag $(ALPHAVERSION)
	git push upstream --tags

tag-beta:
	git tag $(BETAVERSION)
	git push upstream --tags

tag-rc:
	git tag $(RCVERSION)
	git push upstream --tags

tag-release:
	git tag $(RELEASEVERSION)
	git push upstream --tags

build-container:
	docker buildx build --platform linux/arm64,linux/amd64 -t jsnfwlr/facemasq:dev -t jsnfwlr/facemasq:$(CURRENTVERSION) --push -f docker/Dockerfile.multiarch .

build-container-dev:
	docker buildx build --platform linux/arm64,linux/amd64 -t jsnfwlr/facemasq:dev --push -f docker/Dockerfile.multiarch .

build-container-basic:
	docker build -t jsnfwlr/facemasq:basic-$(ARCH) -f docker/Dockerfile .

build-web:
	cd web; pnpm build
