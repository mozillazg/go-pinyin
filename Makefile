help:
	@echo "test             run test"
	@echo "lint             run lint"
	@echo "gen_pinyin_dict  gen pinyin dict"

.PHONY: test
test:
	@echo "run test"
	@go test -v -cover

.PHONY: gen_pinyin_dict
gen_pinyin_dict:
	@go run _tools/gen_pinyin_dict/main.go _tools/pinyin-data/pinyin.txt pinyin_dict.go

.PHONY: gen_phrase_dict
gen_phrase_dict:
	@go run _tools/gen_phrase_dict/main.go _tools/phrase-data/data/phrases-dict.js phrase_dict.go
	@goreturns -w phrase_dict.go

.PHONY: lint
lint:
	gofmt -s -w . cmd/pinyin _tools
	golint .
	golint cmd/pinyin
	golint _tools
