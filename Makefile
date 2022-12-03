SHELL := /bin/zsh
GITVERSION != git describe --tags --abbrev=0
CURRENTVERSION != npm version | jq -r .facemasq
ARCH != go env GOARCH
IMAGENAME != docker info | sed -r "/Username:/!d;s|.*: (.*)|\1/facemasq|"

.PHONY: all tag-alpha tag-beta tag-rc tag-release tag-git api-vet api-test api-coverage api-test-tests api container-release container-dev container-basic web-test web-coverage web docs docs-run docs-gen
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

# API - generate coverage for, vet, test, run, or build the API code
api-coverage:
	@rm data/test.sqlite data/test2.sqlite; cd api; go test -v --tags "full" ./... -covermode=count -coverpkg=./... -coverprofile ../dist/coverage.out; go tool cover -html ../dist/coverage.out -o ../dist/coverage.html
api-vet:
	@cd api; go vet -tags test ./...
api-test:
	@cd api; go test ./...
api-test-db: 					# API - run unit-tests against mysql, pgsql, and sqlite
	@cd api; go test -tags database ./...
api-test-testing:			# API - run unit-tests to ensure the functions used to facilitate testing are working
	@cd api; go test -tags testing ./...
api-test-full: 				# API - run  test-suite
	@cd api; go test --tags full ./...
api-extensions:
	cd api; find ./extensions -maxdepth 1 -type d | grep "\/.*\/" | xargs -n 1 -I {} go build -buildmode=plugin {}
	mv ./api/*.so ./extensions
api-dev: api-extensions
	source env.sh; cd api; go run ./
api: api-test api-extensions 				# API - build local version
	cd api; CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static" --tags "linux sqlite_foreign_keys=1" -o ../dist/api/facemasq .

# CONTAINER - build various versions of the container
container-release:
	docker buildx build --platform linux/arm64,linux/amd64 -t $(IMAGENAME):dev -t $(IMAGENAME):$(CURRENTVERSION) --push -f docker/multiarch.Dockerfile .
container-dev:
	docker buildx build --platform linux/arm64,linux/amd64 -t $(IMAGENAME):dev --push -f docker/multiarch.Dockerfile .
container-basic:
	docker build -t $(IMAGENAME):basic-$(ARCH) -t $(IMAGENAME):basic-$(ARCH)-$(CURRENTVERSION) -f docker/local.Dockerfile .

# WEB - generate coverage for, test, run, or build the web UI code
web-coverage:
	@echo "No Testing Yet"
web-test:
	@echo "No Testing Yet"
web-dev:              # WEB - start dev server
	cd web; pnpm run dev
web: web-test			    # WEB - build the web UI
	cd web; pnpm install; pnpm build

# Docs - generation, serving
docs:
	pnpm run docs:build
docs-run:
	pnpm run docs:dev
docs-gen:
