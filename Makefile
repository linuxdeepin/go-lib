GOPATH_DIR = gopath
GOPKG_PREFIX = github.com/linuxdeepin/go-lib

prepare:
	@mkdir -p ${GOPATH_DIR}/src/$(dir ${GOPKG_PREFIX});
	@ln -sf ../../.. ${GOPATH_DIR}/src/${GOPKG_PREFIX};

print_gopath: prepare
	GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}"

clean:
	rm -rf ${GOPATH_DIR}

check_code_quality: prepare
	env GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}" go vet ./...

test: prepare
	env GOPATH="${CURDIR}/${GOBUILD_DIR}:${GOPATH}" go test -v ./...

test-coverage: prepare
	env GOPATH="${CURDIR}/${GOBUILD_DIR}:${GOPATH}" go test -cover -v ./... | awk '$$1 ~ "^(ok|\\?)" {print $$2","$$5}' | sed "s:${CURDIR}::g" | sed 's/files\]/0\.0%/g' > coverage.csv
