Super simple Golang server.

Command to run:

```
$ PORT=xxxx go run main.go
```

Command to build to get _server.linux_ binary:

```
$ GOOS=linux GOARCH=amd64 go build -v
$ mv super_simple_go_server server.linux
```

[Optional] Refer the _server.service_ file to see the systemd unit file. Take a look at [this guide](https://medium.com/@benmorel/creating-a-linux-service-with-systemd-611b5c8b91d6) on how to set it up on Ubuntu machine.
