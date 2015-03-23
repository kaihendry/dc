all: clean
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
	sudo docker rm test || true
