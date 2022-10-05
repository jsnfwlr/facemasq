CURRENTVERSION != npm version | jq -r .facemasq
ARCH != go env GOARCH
IMAGENAME != docker info | sed -r "/Username:/!d;s|.*: (.*)|\1/facemasq|"

.PHONY: all tag-alpha tag-beta tag-rc tag-release tag-git api-vet api-test api-coverage api-test-tests api container-release container-dev container-basic web-test web-coverage web docs docs-run
.DEFAULT: all

#DEFAULT
all: api web

# Versioning - semver control and commit tagging for different release types
tag-alpha:
	@npm set preid="alpha" && npm version prerelease --preid=alpha
	@make tag-git
tag-beta:
	@npm set preid="beta" && npm version prerelease --preid=beta
	@make tag-git
tag-rc:
	@npm set preid="rc" && npm version prerelease --preid=rc
	@make tag-git
tag-release:
	@npm set preid="" && npm version patch
	$make tag-git
tag-git:
  cd web; npm version $(CURRENTVERSION)
	@git tag v$(CURRENTVERSION)
	@git push upstream --tags

# API - run vet, tests or coverage generation against the API code
api-vet:
	@cd api; go vet -tags test ./...
api-test:
	@cd api; go test ./...

api-coverage:
	@rm data/test.sqlite data/test2.sqlite; cd api; go test -v --tags "full" ./... -covermode=count -coverpkg=./... -coverprofile ../dist/coverage.out; go tool cover -html ../dist/coverage.out -o ../dist/coverage.html

api-test-tests:
	@cd api; go test -tags test ./...
# API - build local version
api: api-test
	cd api; CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" --tags "linux sqlite_foreign_keys=1" -o ../dist/api/facemasq .

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
web: web-test
	cd web; pnpm build

# Docs - generation, serving
docs:
	pnpm run docs:build
docs-run:
	pnpm run docs:dev
