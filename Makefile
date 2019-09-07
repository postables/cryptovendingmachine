.PHONY: all
all: compile bindings

.PHONY: compile
compile:
	# compile vendor management contract
	solc \
	--optimize \
	--optimize-runs 200 \
	--bin \
	--abi \
	--overwrite \
	-o bin/vendorManagement \
	contracts/VendorManagement.sol
	# compile vendor factory
	solc \
	--optimize \
	--optimize-runs 200 \
	--bin \
	--abi \
	--overwrite \
	-o bin/vendorFactory \
	contracts/VendorFactory.sol

.PHONY: bindings
bindings:
	abigen \
	--abi bin/vendorManagement/VendorManagement.abi \
	--bin bin/vendorManagement/VendorManagement.bin \
	--pkg vendormanagement \
	--out bindings/vendorManagement/bindings.go
	abigen \
	--abi bin/vendorFactory/VendorFactory.abi \
	--bin bin/vendorFactory/VendorFactory.bin \
	--pkg vendorfactory \
	--out bindings/vendorFactory/bindings.go

# run standard go tooling for better readability
.PHONY: tidy
tidy: imports fmt
	go vet ./...
	golint ./...

# automatically add missing imports
.PHONY: imports
imports:
	find . -type f -name '*.go' -exec goimports -w {} \;

# format code and simplify if possible
.PHONY: fmt
fmt:
	find . -type f -name '*.go' -exec gofmt -s -w {} \;
