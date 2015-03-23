all: clean
	# this is a workaround for a non-**netgo** CGO_ENABLED=0 enabled build env
	# hopefully netgo will be enabled by default in future go >=1.5 distributions
	go build -a -tags netgo -ldflags "-linkmode external -extldflags -static"
	file ./dc | grep "statically linked"
	# name image "test"
	sudo docker build -t test .
	# name running container "test" after image "test"
	sudo docker run --publish 3000:3000 -d --name test test
	xdg-open http://localhost:3000
	sudo docker ps
	# check image is small <10M
	sudo docker images

clean:
	# stop running container
	sudo docker stop test || true
	# delete container
	sudo docker rm test || true
	# delete image
	sudo docker rmi test || true
	rm -f dc
