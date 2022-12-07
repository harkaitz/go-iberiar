DESTDIR  =
PREFIX   =/usr/local
ETCDIR   =/etc
BINDHTTP =0.0.0.0:8083
SERVER   =m1
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
upload:
	@devh_makefile tar go-iberiar.tar.gz
	@tar-install -S -r $(SERVER) go-iberiar.tar.gz
	@echo "X Restarting..."
	@ssh $(SERVER) sudo sv restart iberiar

## Images
all-go: cmd/iberiar-web/img/favicon.ico
cmd/iberiar-web/img/favicon.ico: logo.png
	@echo "B favicon.ico ..."
	@favigen $< $@
## -- AUTO-SERVICE --
install: install-services
install-services:
	@echo 'I etc/sv/iberiar/run'
	@mkdir -p $(DESTDIR)$(ETCDIR)/sv/iberiar
	@sed "$${SERVICE_SED}" './iberiar.sv.sh' > $(DESTDIR)$(ETCDIR)/sv/iberiar/run
	@chmod +x $(DESTDIR)$(ETCDIR)/sv/iberiar/run
## -- AUTO-SERVICE --
## -- AUTO-GO --
all:     all-go
install: install-go
clean:   clean-go
all-go:
	@echo "B bin/iberiar-web$(EXE) ./cmd/iberiar-web"
	@go build -o bin/iberiar-web$(EXE) $(GO_CONF) ./cmd/iberiar-web
	@echo "B bin/iberiar$(EXE) ./cmd/iberiar"
	@go build -o bin/iberiar$(EXE) $(GO_CONF) ./cmd/iberiar
install-go: all-go
	@install -d $(DESTDIR)$(PREFIX)/bin
	@echo I bin/iberiar-web$(EXE)
	@cp bin/iberiar-web$(EXE) $(DESTDIR)$(PREFIX)/bin
	@echo I bin/iberiar$(EXE)
	@cp bin/iberiar$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	@echo "D Go binaries ..."
	@rm -f bin/iberiar-web
	@rm -f bin/iberiar
## -- AUTO-GO --
