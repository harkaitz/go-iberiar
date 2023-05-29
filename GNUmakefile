DESTDIR  =
PREFIX   =/usr/local
ETCDIR   =/etc
BINDHTTP =0.0.0.0:8083

define SERVICE_SED
s|@PREFIX@|$(PREFIX)|
s|@EXE@|$(EXE)|
s|@BINDHTTP@|$(BINDHTTP)|
endef
export SERVICE_SED
all:
clean:
install:
update:

## Images
deps: cmd/iberiar-web/img/favicon.ico
cmd/iberiar-web/img/favicon.ico: logo.png
	@echo "B favicon.ico ..."
	@favigen $< $@
## -- AUTO-GO --
GO_PROGRAMS += bin/iberiar-web$(EXE) bin/iberiar$(EXE) 
.PHONY all-go: $(GO_PROGRAMS)
all:     all-go
install: install-go
clean:   clean-go
deps:
bin/iberiar-web$(EXE): deps 
	go build -o $@ $(IBERIAR_WEB_FLAGS) $(GO_CONF) ./cmd/iberiar-web
bin/iberiar$(EXE): deps 
	go build -o $@ $(IBERIAR_FLAGS) $(GO_CONF) ./cmd/iberiar
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp bin/iberiar-web$(EXE) $(DESTDIR)$(PREFIX)/bin
	cp bin/iberiar$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(GO_PROGRAMS)
## -- AUTO-GO --
## -- license --
install: install-license
install-license: LICENSE
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/go-iberiar
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/go-iberiar
## -- license --
## -- AUTO-SERVICE --
install: install-services
install-services:
	mkdir -p $(DESTDIR)$(ETCDIR)/sv/iberiar
	sed "$${SERVICE_SED}" './iberiar.sv.sh' > $(DESTDIR)$(ETCDIR)/sv/iberiar/run
	chmod +x $(DESTDIR)$(ETCDIR)/sv/iberiar/run
## -- AUTO-SERVICE --
