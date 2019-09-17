dir = ./...

travis: setup gofmt golint tests;

tests: ;
	go test $(dir)

benchmark: ;
	go test $(dir) -v --bench . --benchmem

gofmt: ;
	bash -c "diff -u <(echo -n) <(gofmt -d ./)"

golint: ;
	bash -c "diff -u <(echo -n) <(golint ./...)"

setup: ;
	go get -u golang.org/x/lint/golint
