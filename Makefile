.PHONY : build txstorm
build :
	go build -o build/lachesis ./cmd/lachesis

txstorm :
	go build -o build/tx-storm ./cmd/tx-storm

.PHONY : test
test :
	go test ./...

.PHONY: coverage
coverage:
	go test -coverpkg=./... -coverprofile=cover.prof ./...
	go tool cover -func cover.prof | grep -e "^total:"


.PHONY: clean
clean:
	rm ./build/lachesis ./build/tx-storm
