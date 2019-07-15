# schat

`schat` is a small chat application with gRPC bidirectional streaming RPC.

## Usage 

```
$ schat --help
Usage of schat:
  -h string
        host of server (default "localhost:8080")
  -s    run as a server
  -u string
        username of client (default "anonymous")
```

### Run as Server

```
$ schat -s
2019/07/15 00:35:19 server: starting
2019/07/15 00:35:23 micnncim | hello
```

### Run as Client

```
$ schat -u micnncim
2019/07/15 00:35:21 client: starting
2019/07/15 00:35:21 client: connected to stream
>>> hello
2019/07/15 00:35:23 micnncim | hello
```
