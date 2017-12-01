package main

import (
    "github.com/samuel/go-zookeeper/zk"
    "os"
    "os/signal"
    "strings"
    "strconv"
    "time"
    "log"
)

var RootPath string = "/zookeeper-for-myApp"
var zksStr string = os.Getenv("ZOOKEEPER_SERVERS")

func CreateRootPath(zks []string) *zk.Conn {
    flags := int32(0)
    acl := zk.WorldACL(zk.PermAll)
    conn, _, err := zk.Connect(zks, time.Second)
    conn.Create(RootPath, []byte{0}, flags, acl)
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    return conn
}


func JoinRootPath(zks []string) {
    acl := zk.WorldACL(zk.PermAll)
    conn, _, err := zk.Connect(zks, time.Second)
    if err != nil {
        panic(err)
    }
    lockPrefix := "/lock-"

    path, err := conn.CreateProtectedEphemeralSequential(RootPath+lockPrefix,
                                                         []byte{0}, acl)
    myData := getParse(path)
    ticker := time.NewTicker(time.Second)
    for {
        var isLeader bool = true
        children, _, err := conn.Children(RootPath)
        if err != nil {
            isLeader = false
        }
        if children == nil {
            isLeader = false
        }
        for _, c := range children {
            otherData := getParse(c)
            if myData > otherData {
                isLeader = false
                break
            }
        }

        if isLeader == true {
            log.Println("### I am Leader !! ###")
        }
        <-ticker.C
    }
}

func getParse(path string) int {
    splits := strings.Split(path, "-")
    seq := splits[len(splits)-1]
    data, _ := strconv.Atoi(seq)
    return data
}

func main() {
    zks := strings.Split(zksStr, ",")
    CreateRootPath(zks)
    quit_channel := make(chan os.Signal, 1)
    signal.Notify(quit_channel, os.Interrupt)
    go JoinRootPath(zks)
    <-quit_channel
}
