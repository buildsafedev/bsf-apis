PROTO ?= buildsafe/v1/search.proto

lint: buf-produce-bin
	api-linter --set-exit-status --output-format yaml  --descriptor-set-in descriptorset.bin $(PROTO)
 

lint-summary: buf-produce-bin
	api-linter --set-exit-status --output-format summary --descriptor-set-in descriptorset.bin $(PROTO)

buf-produce-bin:
	buf build  --output descriptorset.bin --as-file-descriptor-set
