NAME=buk-server

local-binary:
	go build -o ${NAME} main.go

local-binary-run:
	./buk-server

docker-build-image:
	docker build -t jayroy/bukserver .

docker-run:
	docker run -it -p 8877:8877 jayroy/bukserver

test:
	cd client && go test 
