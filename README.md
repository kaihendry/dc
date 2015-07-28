Run a statically built golang binary in a systemd-nspawn container example

	make container

Tested successfully on Archlinux

# Pros over Docker

* simpler
* faster
* lightweight

## Size of contianer

	$ sudo du -sh foobar
	17M     foobar

# Related issues

* <https://github.com/systemd/systemd/issues/750#issuecomment-125402766>

journalctl is not working.
