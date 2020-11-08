lint:
	golangci-lint run --disable-all -E golint -E staticcheck -E structcheck -E unused -E gocritic -E gofmt -E interfacer -E misspell -E stylecheck -E unconvert -E unparam -E scopelint -E prealloc
