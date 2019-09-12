GOPATH_DIR = gopath
GOPKG_PREFIX = pkg.deepin.io/lib

prepare:
	@if [ ! -d ${GOPATH_DIR}/src/${GOPKG_PREFIX} ]; then \
		mkdir -p ${GOPATH_DIR}/src/$(dir ${GOPKG_PREFIX}); \
		ln -sf ../../.. ${GOPATH_DIR}/src/${GOPKG_PREFIX}; \
		fi

clean:
	rm -rf ${GOPATH_DIR}

check_code_quality: prepare
	env GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}" go vet ./...
