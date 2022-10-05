GITVERSION != git describe --tags --abbrev=0
CURRENTVERSION != npm version | jq -r .facemasq
ARCH != go env GOARCH
IMAGENAME != docker info | sed -r "/Username:/!d;s|.*: (.*)|\1/facemasq|"

.PHONY: all tag-alpha tag-beta tag-rc tag-release tag-git api-vet api-test api-coverage api-test-tests api container-release container-dev container-basic web-test web-coverage web docs docs-run
.DEFAULT: all

#DEFAULT
all: api web

# Versioning - semver control and commit tagging for different release types
tag-alpha:
	@npm version $(GITVERSION) --allow-same-version --git-tag-version=false
	@npm version prerelease --preid=alpha
	$(MAKE) tag-git
tag-beta:
	@npm version $(GITVERSION) --allow-same-version --git-tag-version=false
	@npm version prerelease --preid=beta
	$(MAKE) tag-git
tag-rc:
	@npm version $(GITVERSION) --allow-same-version --git-tag-version=false
	@npm version prerelease --preid=rc
	$(MAKE) tag-git
tag-release:
	@npm version $(GITVERSION) --allow-same-version --git-tag-version=false
	@npm version patch
	$(MAKE) tag-git
tag-git:
	@git push dev --all

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
