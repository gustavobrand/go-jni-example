GO_FILES=$(wildcard lib/*.go)
CFLAGS+=-I/usr/lib/jvm/java-9-openjdk-amd64/include
CFLAGS+=-I/usr/lib/jvm/java-9-openjdk-amd64/include/linux

.PHONY: run-linux
run-linux: app/Main.class libfetch.so
	@LD_LIBRARY_PATH=$(PWD) java app.Main

libfetch.so: lib/app_Main.h $(GO_FILES)
	CGO_CFLAGS="$(CFLAGS)" go build -buildmode=c-shared -o $(@) ./lib

app/%.class: app/%.java
	javac $(<)

lib/app_%.h: PACKAGE_NAME=$(subst /,.,$(<:%.class=%))
lib/app_%.h: app/%.class
	@javah -d $(@D) $(PACKAGE_NAME) > /dev/null
