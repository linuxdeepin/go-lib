PREFIX = /usr
GOSITE_DIR = ${PREFIX}/share/gocode
GOPKG_PERFIX = pkg.deepin.io/lib
LIBARAIES = \
app \
appinfo \
arch \
archive \
asound \
backlight \
calendar \
cgroup \
dbus \
dbus1 \
dbusutil \
encoding \
event \
fsnotify \
gdkpixbuf \
gettext \
graphic \
gsettings \
initializer \
iso \
keyfile \
locale \
log \
mime \
mobileprovider \
notify \
pam \
pinyin \
procfs \
profile \
proxy \
pulse \
sound \
sound_effect \
stb_vorbis \
strv \
tasker \
timer \
users \
utils \
imgutil \
xdg

all: build

build:
	echo ignore build

install: install-prepare install-dev
	cp *.go ${DESTDIR}${GOSITE_DIR}/src/${GOPKG_PERFIX}

install/lib/%:
	cp -r ${@F} ${DESTDIR}${GOSITE_DIR}/src/${GOPKG_PERFIX}

install-dev: ${addprefix install/lib/, ${LIBARAIES}}

install-prepare:
	mkdir -p ${DESTDIR}${GOSITE_DIR}/src/${GOPKG_PERFIX}
