prefix=/usr/local
CGO_ENABLED?=0
GOOS?=darwin

all: build

build: hello

install: build
	mkdir -p $(DESTDIR)$(prefix)/bin
	cp hello $(DESTDIR)$(prefix)/bin/hello
	$(DESTDIR)$(prefix)/bin/hello

uninstall:
	rm $(DESTDIR)$(prefix)/bin/hello

clean:
	go clean

hello: main.go
	go fmt
	go build

.PHONY: all build install uninstall clean
