.POSIX:
.SUFFIXES:
.PHONY: all clean install pack l10n
all:
PROJECT   =go-iberiar
VERSION   =1.0.1
DESTDIR   =
PREFIX    =/usr/local
ETCDIR    =/etc

help:
	@echo 'all     : Build binaries.'
	@echo 'clean   : Clean generated files.'
	@echo 'install : Install binaries.'
	@echo 'pack    : Create tar file.'
	@echo 'l10n    : Update locales.'

## -------------------------------------------------------------------
pack:
	hsrc-wd -n 'go-iberiar' -z -t x86_64-linux-gnu  -p $(PREFIX) hgmake clean all install
	hsrc-wd -n 'go-iberiar' -z -t aarch64-linux-gnu -p $(PREFIX) hgmake clean all install
l10n:
	gettext-h-i18n -ub


## -------------------------------------------------------------------
install: install-services
install-services:
	install -D iberiar.json $(DESTDIR)$(ETCDIR)/site/iberiar.json.example
	runit-h   -D $(DESTDIR) -a iberiar $(PREFIX)/bin/iberiar
	systemd-h -D $(DESTDIR) -a iberiar $(PREFIX)/bin/iberiar

## -------------------------------------------------------------------
build/iberiar$(EXE): html/img/favicon.ico
html/img/favicon.ico: html/img/logo.png
	favigen $< $@

## -- BLOCK:go --
build/iberiar$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/iberiar
all: build/iberiar$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/iberiar$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/iberiar$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
