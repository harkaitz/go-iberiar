#!/bin/sh
exec @PREFIX@/bin/iberiar-web@EXE@ @BINDHTTP@ >> /var/log/iberiar.log 2>&1
