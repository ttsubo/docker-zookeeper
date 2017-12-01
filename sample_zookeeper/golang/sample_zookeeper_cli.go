package main

import (
    "github.com/samuel/go-zookeeper/zk"
    "os"
    "strings"
    "time"
    "fmt"
)

var RootPath string = "/zookeeper-for-myApp"
var zksStr string = os.Getenv("ZOOKEEPER_SERVERS")

func Connect(zks []string) *zk.Conn {
    conn, _, err := zk.Connect(zks, time.Second)
    if err != nil {
        panic(err)
    }
    return conn
}

func main() {
    zks := strings.Split(zksStr, ",")
    conn := Connect(zks)
    defer conn.Close()

    data, _, err := conn.Get(RootPath)
    fmt.Printf("%s: %s\n", RootPath, string(data))

    children, _, err := conn.Children(RootPath)
    if err != nil {
        panic(err)
    }
        for _, name := range children {
            data, _, err := conn.Get(RootPath+"/"+name)
            if err != nil {
                panic(err)
            }
        fmt.Printf("%s: %s\n", name, string(data))
    }
}
