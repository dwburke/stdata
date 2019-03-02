
APPNAME=stdata


default:
	$(MAKE) build

test:
	go test ./... -v

run:
	$(MAKE) build
	./${APPNAME}

build:
	go build

clean:
	rm ${APPNAME}

