package main

import (
    "os"
    "path/filepath"
    "time"
    "errors"
    //"fmt"
    //"log"
    )

type CacheNode struct{
        Name string       // base name of the file
        Size uint64        // length in bytes for regular files; system-dependent for others
        Mode os.FileMode     // file mode bits
        ModTime time.Time // modification time
        IsDir bool        // is dir?
    }

type CacheDir struct {
    CacheNode
    children []CacheNode
}

type DB struct {
    Debugger
    dirs map[string]CacheDir
}


func NewDatabase() *DB {
    d:=new(DB)
    d.dirs = make(map[string]CacheDir)
    d.Debpref=("database")
    return d
}

func (m *DB) OpenDir(path string) (cnode []CacheNode,err error){
    m.log("OpenDir(): %s",path)
    if m.dirs==nil {
    m.log("OpenDir(): making map")
    m.dirs=make(map[string]CacheDir)
    }
    a,ok := m.dirs[path]
    if ok{
        return a.children, nil
        m.log("cache")
    }
    cnode,err = m.updateDir(path)
    return
}

func (m *DB) updateDir(path string) (cnode []CacheNode, err error){
    nodes, err := wc.ReadDir(path)
    if err!= nil {
        return nil,err
    }
    for _,node := range nodes{
        var c CacheNode
        c.Size = uint64(node.Size())
        c.Name = node.Name()
        c.IsDir = node.IsDir()
        c.ModTime = node.ModTime()
        cnode = append(cnode,c)
    }
    
    var cd CacheDir
    cd.children = cnode
    m.dirs[path] = cd
    m.log("updateDir(): %s nodes: %d",path,cap(m.dirs[path].children))
    return
}

func (m *DB) getEnt(path string)(file CacheNode, err error){
    dir := filepath.Dir(path)
    fname := filepath.Base(path)
    children,err := m.OpenDir(dir)
    if err ==nil {
        for _,child := range children {
            if child.Name == fname {
                return child,nil
            }
    }
    }
    return file, errors.New("Merda")
}

func (m *DB) GetAttr(path string) (file CacheNode, err error){
    m.log("GetAttr() %s",path)
    file,err = m.getEnt(path)
    if file.IsDir {
        m.log("GetAttr(): %s ISDIR",path)
        db.OpenDir(path)
        file,err = m.getEnt(path)
    }
    return file, err
}

/*
func (m *DB) log (format string, v ...interface{}){
    format = m.debpref+format
    dbgPrintf(format,v...)
}*/
