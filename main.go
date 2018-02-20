package main

import (
    "flag"
    "fmt"
    "os"
    "log"


//    "github.com/hanwen/go-fuse/fuse"
    "github.com/hanwen/go-fuse/fuse/nodefs"
    "github.com/hanwen/go-fuse/fuse/pathfs"
    webdav "github.com/studio-b12/gowebdav"
    )

var db *DB
var wc *webdav.Client

func usage(){
    fmt.Println("Usage:  gowebdavfs [options] URL MOUNTPOINT")
    flag.PrintDefaults()
}

func main(){
    username := flag.String("u","","webdav username")
    password := flag.String("p","","webdav password")
    flag.Parse()
    if len(flag.Args()) < 2 {
        usage()
        os.Exit(1)
    }

    db = NewDatabase()

    wc=webdav.NewClient(flag.Arg(0),*username,*password)
    err :=wc.Connect()
    if err != nil {
        log.Panicf("Error conetcing to server: %v",err)
    } else {
        log.Printf("Connected to server.")
    }

    nfs := pathfs.NewPathNodeFs(&GoWebDav{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
    server, _, err := nodefs.MountRoot(flag.Arg(1), nfs.Root(), nil)
    if err != nil {
        log.Fatalf("Mount fail: %v\n", err)
    }
    server.Serve()
}
