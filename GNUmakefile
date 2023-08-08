PROJECT   =go-iberiar
VERSION   =1.0.0
DESTDIR   =
PREFIX    =/usr/local
ETCDIR    =/etc
BINDHTTP  =127.0.0.1:8084
PLATFORMS =linux-arm64
SERVICES  =rpi3b:s=iberiar:p=linux-arm64
HBINDINGS =iberiar.eu:i=127.0.0.1:p=8084
HSERVICE  =rpi3b:s=https:i=192.168.1.6

all:
clean:
install:
update:
## -- SERVICES --
install-services:
	@lsetup-runit   -v -D "$(DESTDIR)" -a "iberiar" "$(PREFIX)/bin/iberiar-web" "$(BINDHTTP)"
	@lsetup-systemd -v -D "$(DESTDIR)" -a "iberiar" "$(PREFIX)/bin/iberiar-web" "$(BINDHTTP)"
## -- SERVICES --
## -- IMAGES --
deps: cmd/iberiar-web/img/favicon.ico
cmd/iberiar-web/img/favicon.ico: logo.png
	favigen $< $@
## -- IMAGES --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE README.md $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
update: update-license
update-license:
	ssnip README.md
## -- BLOCK:license --
## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps: deps-go

build/iberiar-web$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/iberiar-web
build/iberiar$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/iberiar

all-go:  build/iberiar-web$(EXE) build/iberiar$(EXE)
deps-go:
	mkdir -p build
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp  build/iberiar-web$(EXE) build/iberiar$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f  build/iberiar-web$(EXE) build/iberiar$(EXE)
## -- BLOCK:go --
