# easyjson example

[doc](https://github.com/euank/ifacepropagate)

```Go
go get -u https://github.com/euank/ifacepropagate/cmd/...
ifacepropagate . "l *closeLoggedConn.Conn" io.ReaderFrom,syscall.Conn > igen_generated.go
ifacepropagate . "l *closeLoggedConn.Conn" io.ReaderFrom,syscall.Conn
go build
```

ifacepropagate . "l *closeLoggedConn.Conn" io.ReaderFrom,syscall.Conn
## reference

