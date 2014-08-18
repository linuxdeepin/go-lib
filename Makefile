PREFIX = /usr
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GOPATH_DIR = gopath
GOPKG_PREFIX = pkg.linuxdeepin.com/lib

LIBS =  \
    dbus \
    gdkpixbuf \
    gettext \
    gio-2.0 \
    glib-2.0 \
    gobject-2.0 \
    graphic \
    log \
    pinyin \
    proxy \
    pulse \
    utils \

all: build

prepare:
	@if [ ! -d ${GOPATH_DIR}/src/${GOPKG_PREFIX} ]; then \
		mkdir -p ${GOPATH_DIR}/src/$(dir ${GOPKG_PREFIX}); \
		ln -sf ../../.. ${GOPATH_DIR}/src/${GOPKG_PREFIX}; \
	fi

build/lib/%:
	env GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}" go install ${GOPKG_PREFIX}/${@F}

build: prepare $(addprefix build/lib/, ${LIBS})
	env GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}" go install ${GOPKG_PREFIX}

install-prepare:
	mkdir -p ${DESTDIR}${PREFIX}/share/go/src/pkg/${GOPKG_PREFIX}
	mkdir -p ${DESTDIR}${PREFIX}/lib/go/pkg

install/src/%:
	cp -r ${CURDIR}/${@F} ${DESTDIR}${PREFIX}/share/go/src/pkg/${GOPKG_PREFIX}/

install: build install-prepare $(addprefix install/src/, ${LIBS})
	cp -r ${CURDIR}/*.go ${DESTDIR}${PREFIX}/share/go/src/pkg/${GOPKG_PREFIX}/
	cp -r ${CURDIR}/${GOPATH_DIR}/pkg/* ${DESTDIR}${PREFIX}/lib/go/pkg/

uninstall:
	rm -rvf ${DESTDIR}${PREFIX}/share/go/src/pkg/${GOPKG_PREFIX}
	rm -rvf ${DESTDIR}${PREFIX}/lib/go/pkg/${GOOS}_${GOARCH}/${GOPKG_PREFIX}

clean:
	rm -rf ${GOPATH_DIR}

rebuild: clean build
