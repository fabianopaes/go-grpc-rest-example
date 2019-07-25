include help.mk

GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	golang.org/x/tools/cmd/cover \
	github.com/axw/gocov/gocov \
	gopkg.in/matm/v1/gocov-html

.PHONY: clean
clean: ##@development Remove all generated files
	-@rm -f $(NAME); \
	rm -rf .build; \
	rm -rf  $(pwd)/api/*.pb.go

.PHONY: dep
dep: tools ##@development Install external godep packages
	dep ensure -v 

run: ##@development Use it when you want to get your local enviroment up
	PORT=9999 \
	go run cmd/go-rest-sample/main.go	