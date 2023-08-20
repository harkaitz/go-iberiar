PROJECT   =go-iberiar
VERSION   =1.0.1
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

install-services:
	@lsetup-runit   -v -D "$(DESTDIR)" -a "iberiar" "$(PREFIX)/bin/iberiar" "$(BINDHTTP)"
	@lsetup-systemd -v -D "$(DESTDIR)" -a "iberiar" "$(PREFIX)/bin/iberiar" "$(BINDHTTP)"

deps: cmd/iberiar/img/favicon.ico
cmd/iberiar/img/favicon.ico: logo.png
	favigen $< $@

## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps: deps-go

build/iberiar$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/iberiar

all-go:  build/iberiar$(EXE)
deps-go:
	mkdir -p build
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp  build/iberiar$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f  build/iberiar$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE README.md $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
update: update-license
update-license:
	ssnip README.md
## -- BLOCK:license --
