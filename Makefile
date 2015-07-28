CONTAINER_DIR=foobar

dc: clean
	go build -a -tags netgo -installsuffix netgo .
	file ./dc | grep "statically linked"

container: dc
	mkdir $(CONTAINER_DIR) || true
	sudo pacstrap -idc $(CONTAINER_DIR) filesystem
	cp dc $(CONTAINER_DIR)
	sudo systemd-nspawn -D $(CONTAINER_DIR) /dc

clean:
	sudo rm -rf dc $(CONTAINER_DIR)
