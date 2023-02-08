all: validate clean generate build

validate:
	swagger validate ./api/swagger.yml

spec:
	swagger generate spec -o ./api/swagger-gen.yml

clean:
	rm -rf crud-goswagger-server
	rm -rf crud-goswagger-migrator
	go clean -i .

generate:validate clean
	swagger -q generate server -f api/swagger.yml -A Crud_Goswagger -s internal/apis -m pkg/models
	swagger -q generate client -f api/swagger.yml -A Crud_Goswagger -c pkg/client -m pkg/models

doc:
	swagger validate ./api/swagger.yml
	swagger serve api/swagger.yml --no-open --host=0.0.0.0 --port=8000 --base-path=/
	
build: clean
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/crud-goswagger-server

run: 
	./crud-goswagger-server --port=8000 --host=0.0.0.0

.PHONY: validate spec clean generate doc build run 
