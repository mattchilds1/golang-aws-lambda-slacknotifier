.PHONY: test clean

test: apply

clean:
	cd iac; terraform destroy
	rm -f init apply build.zip build.app

init:
	cd iac; terraform init

apply: init build.zip
	cd iac; terraform apply

build.zip: build.app
	zip $@ $<

build.app: main.go
	go get .
	GOOS=linux GOARCH=amd64 go build -o $@


	