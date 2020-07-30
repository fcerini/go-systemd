# go-systemd
Run a Go binary as a daemon under systemd


Edit the paths in go-systemd.service

```
cp go-systemd.service /lib/systemd/system/

service go-systemd start

service go-systemd status

journalctl -f -u go-systemd

# on boot
systemctl enable go-systemd

```
