.POSIX:
.SUFFIXES:
.PHONY: all clean install

PROJECT   =go-iberiar
VERSION   =1.0.1
DESTDIR   =
PREFIX    =/usr/local
BUILDDIR  =.build

## -------------------------------------------------------------------
all:
## -------------------------------------------------------------------
build/iberiar$(EXE): html/img/favicon.ico
html/img/favicon.ico: html/img/logo.png
	favigen $< $@
## -------------------------------------------------------------------
all: pkg/i-font-iberian pkg/i-font-uniedit-iberian
pkg/i-font-iberian:
	cp $$(getprj i)/bin/i-font-iberian $@
pkg/i-font-uniedit-iberian:
	cp $$(getprj i)/bin/i-font-uniedit-iberian $@
## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/iberiar$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/iberiar$(EXE)
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	install -c -m 755 $(BUILDDIR)/iberiar$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/iberiar$(EXE)
##
$(BUILDDIR)/iberiar$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/iberiar
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: README.md LICENSE
	install -d $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	install -c -m 644 README.md LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
