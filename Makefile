ifndef GOPATH
	GOPATH := ${PWD}/gopath
	export GOPATH
endif

pcd-ctl: *.go ${GOPATH}/src/github.com/spf13/cobra
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o pcd-ctl

${GOPATH}/src/github.com/spf13/cobra:
	echo "${GOPATH}"
	go get -v github.com/spf13/cobra

