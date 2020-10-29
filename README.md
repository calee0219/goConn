# goConn

OpenNetVM version: 27d9ed5d

## Folder structure:
```
/openNetVM
  /examples
    /goConn # this repo
      main.go
      go.mod
      go.sum
      /onvmConn
```

## Build

Modify path inside `onvmConn/conn.go`'s `#cgo` tag path.

Modify config path inside `onvmConn/conn.c`'s cmd's path.

make
