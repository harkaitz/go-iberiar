PROJECT   =go-iberiar
VERSION   =1.0.1
DESTDIR   =
PREFIX    =/usr/local
ETCDIR    =/etc
#HBINDINGS =iberiar.eu:i=127.0.0.1:p=8084
#HSERVICE  =rpi3b:s=https:i=192.168.1.6

all:
clean:
install: install-services

##
pack:
	hsrc -z -t x86_64-linux-gnu  -p $(PREFIX) iberiar
	hsrc -z -t aarch64-linux-gnu -p $(PREFIX) iberiar
install-services:
	install -D iberiar.json $(DESTDIR)$(ETCDIR)/site/iberiar.json.example
	runit-h   -D $(DESTDIR) -a iberiar $(PREFIX)/bin/iberiar
	systemd-h -D $(DESTDIR) -a iberiar $(PREFIX)/bin/iberiar
##
deps: html/img/favicon.ico
	gettext-h-i18n -ub
html/img/favicon.ico: html/img/logo.png
	favigen $< $@
## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps:
build/iberiar$(EXE): deps
	@mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/iberiar
all-go: build/iberiar$(EXE)
install-go:
	install -D -t $(DESTDIR)$(PREFIX)/bin build/iberiar$(EXE)
clean-go:
	rm -f build/iberiar$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	install -D -t $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT) LICENSE
## -- BLOCK:license --
