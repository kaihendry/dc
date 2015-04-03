all: clean
	# golang 1.4 needed
	go build -a -tags netgo -installsuffix netgo .
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
