setup:
	go get ./...
  	go get -u github.com/onsi/ginkgo/ginkgo
  	go get -u github.com/onsi/gomega/...  

test:
	ginkgo -r