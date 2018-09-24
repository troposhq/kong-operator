PKGS := $(shell go list ./... | grep -v -e /vendor -e /pkg/apis -e /pkg/client/)

.PHONY: test
test:
	go test $(PKGS)

.PHONY: coverage
coverage: ## Generate global code coverage report
	./tools/coverage.sh;

# generates: 
# deepcopy funcs
# clientset
# listers
# informers
codegen:
	./hack/update-codegen.sh
