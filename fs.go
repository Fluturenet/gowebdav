package main

import (
    "log"
    "github.com/hanwen/go-fuse/fuse"
    "github.com/hanwen/go-fuse/fuse/nodefs"
    "github.com/hanwen/go-fuse/fuse/pathfs"
    )


func NewFileSystem() (*pathfs.PathNodeFs){
    GWD:= &GoWebDav{FileSystem: pathfs.NewDefaultFileSystem()}
    GWD.Debpref = "filsesystem"
    fs := pathfs.NewPathNodeFs(GWD,nil)
    return fs
}


type GoWebDav struct {
    pathfs.FileSystem
    Debugger
    }

func (me *GoWebDav) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
    me.log("GetAttr(): %s",name)

    if name == "" {
        return &fuse.Attr{Mode: fuse.S_IFDIR | 0555}, fuse.OK
    }

    file,err := db.GetAttr("/"+name)
    if err==nil{
        var Mode uint32
        var Size uint64
        if file.IsDir {
            Mode = fuse.S_IFDIR | 0555
            //&fuse.Attr{Mode: fuse.S_IFDIR | 0755}
        } else {
            Mode = fuse.S_IFREG | 0444
            Size = file.Size
            //&fuse.Attr{Mode: fuse.S_IFREG | 0644}
        }
        return &fuse.Attr{Mode: Mode,Size: Size}, fuse.OK
    }
    return nil, fuse.ENOENT
}

func (me *GoWebDav) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
    me.log("Open(): %s",name)
    f := NewWebDavFile("/"+name)
    return f, fuse.OK
}

func (me *GoWebDav) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, status fuse.Status) {
    me.log("OpenDir(): %s",name)
    files,err := db.OpenDir("/"+name)
    var r []fuse.DirEntry
    if err != nil {
        log.Panicf("Error opening Dir:%s\n %v\n",name,err)
        }
    for _,file := range files {
        var dr fuse.DirEntry
        dr.Name = file.Name
        if file.IsDir {
            dr.Mode = fuse.S_IFDIR | 0555 
        } else {
            dr.Mode=fuse.S_IFREG | 0444
        }
        r = append(r,dr)
    }
    return r, fuse.OK
}


