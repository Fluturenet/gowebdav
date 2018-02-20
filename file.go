package main

import (
    "io"
    "time"
    "github.com/hanwen/go-fuse/fuse"
    "github.com/hanwen/go-fuse/fuse/nodefs"
    )

type WebDavFile struct {
    Debugger
    URL string
    offset int64
    size int64
    file io.ReadCloser
}

func (m *WebDavFile) Read(dest []byte, off int64) (fuse.ReadResult, fuse.Status) {
    m.log("Read() %s: requested_bytes %d offset %d",m.URL,cap(dest),off)
    if (off!=m.offset) {
        m.log("Read() %s: OffsetMismatch",m.URL)
        return fuse.ReadResultData(dest),fuse.EAGAIN
    }
    if(m.file ==nil) {
        //get file info
        info,_ := db.GetAttr(m.URL)
        m.size = int64(info.Size)
        m.file,_ = wc.ReadStream(m.URL)
        //check err
    }
    
    var i int
    var err error
    if m.size-m.offset > int64(cap(dest)) {
        i,err =io.ReadAtLeast(m.file,dest,cap(dest))
    } else {
        i,err =io.ReadAtLeast(m.file,dest,int(m.size-m.offset))
    }
    
    if err!=nil {
        m.log("Read() %s: read %d with err. %v",m.URL,i,err)
    }
    m.offset += int64(i)
    return fuse.ReadResultData(dest),fuse.OK
}

func (m *WebDavFile) Release(){
    m.log("Release(): %s",m.URL)
    if(m.file == nil) {return }
    m.file.Close()
}


func NewWebDavFile(path string) nodefs.File  {
    f:= new(WebDavFile)
    f.URL = path
    f.Debpref="WebDavFile"
    return f
}

/* nohing of interesting */

func (m *WebDavFile) SetInode(*nodefs.Inode) {
}

func (m *WebDavFile) InnerFile() nodefs.File {
    return nil
}

func (m *WebDavFile) String() string {
    return "WebDavFile"
}


func (m *WebDavFile) Write(data []byte, off int64) (uint32, fuse.Status) {
    return 0, fuse.ENOSYS
}

func (m *WebDavFile) Flock(flags int) fuse.Status { return fuse.ENOSYS }
func (m *WebDavFile) Flush() fuse.Status {
    return fuse.OK
}


func (m *WebDavFile) GetAttr(*fuse.Attr) fuse.Status {
    
    return fuse.ENOSYS
}

func (m *WebDavFile) Fsync(flags int) (code fuse.Status) {
    return fuse.ENOSYS
}

func (m *WebDavFile) Utimens(atime *time.Time, mtime *time.Time) fuse.Status {
    return fuse.ENOSYS
}

func (m *WebDavFile) Truncate(size uint64) fuse.Status {
    return fuse.ENOSYS
}

func (m *WebDavFile) Chown(uid uint32, gid uint32) fuse.Status {
    return fuse.ENOSYS
}

func (m *WebDavFile) Chmod(perms uint32) fuse.Status {
    return fuse.ENOSYS
}

func (m *WebDavFile) Allocate(off uint64, size uint64, mode uint32) (code fuse.Status) {
    return fuse.ENOSYS
}

