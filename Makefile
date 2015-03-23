all: clean
	go build -a -tags netgo -ldflags "-linkmode external -extldflags -static"
	# name image "test"
	sudo docker build -t test .
	# name running container "test" after image "test"
	sudo docker run --publish 3000:3000 -d --name test test
	xdg-open http://localhost:3000
	sudo docker ps
	# why is the image SO BIG?
	sudo docker images

clean:
	sudo docker stop test || true
	# delete container
	sudo docker rm test || true
	# delete image
	sudo docker rmi test || true
	rm -f dc
