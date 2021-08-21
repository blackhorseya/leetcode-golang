.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20

.PHONY: clean # remove files
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: report # refresh goreportcard
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/leetcode-golang'

.PHONY: lint # execute golint
lint:
	@golint ./...

.PHONY: test-unit # execute unit test
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh
