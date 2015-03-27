GO=go 
GC=gc
MAIN=vde.go
TARGET=vde.test

all:
	${GO} build -compiler ${GC} github.com/kurojishi/vde-go

install:
	${GO} install -compiler ${GC} github.com/kurojishi/vde-go

test:
	${GO} test -compiler ${GC} github.com/kurojishi/vde-go

clean:
	rm -f ${TARGET}

