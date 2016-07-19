package main

import (
    "fmt"
    "github.com/hashicorp/raft"
    "github.com/hashicorp/raft-boltdb"
    "io"
    "log"
    "os"
    "path"
    "time"
)

type Config struct {
    Bind    string `json:bind`
    DataDir string `json:data_dir`
}

type Word struct {
    words string
}

func (*Word) Apply(l *raft.Log) interface{} {
    return nil
}

func (*Word) Snapshot() (raft.FSMSnapshot, error) {
    return new(WordSnapshot), nil
}

func (*Word) Restore(snap io.ReadCloser) error {
    return nil
}

type WordSnapshot struct {
    words string
}

func (snap *WordSnapshot) Persist(sink raft.SnapshotSink) error {
    return nil
}

func (snap *WordSnapshot) Release() {

}

func main() {
    v := Config{Bind: os.Args[2], DataDir: os.Args[1]}
    os.MkdirAll(v.DataDir, 0755)

    cfg := raft.DefaultConfig()
    // cfg.EnableSingleNode = true
    fsm := new(Word)
    fsm.words = "hello"

    dbStore, err := raftboltdb.NewBoltStore(path.Join(v.DataDir, "raft.db"))
    if err != nil {
        log.Fatal(err)
    }

    fileStore, err := raft.NewFileSnapshotStore(v.DataDir, 1, os.Stdout)
    if err != nil {
        log.Fatal(err)
    }

    trans, err := raft.NewTCPTransport(v.Bind, nil, 3, 5*time.Second, os.Stdout)
    if err != nil {
        log.Fatal(err)
    }

    peers := make([]string, 0, 10)
    for i := 3; i < len(os.Args); i++ {
        peers = raft.AddUniquePeer(peers, os.Args[i])
    }

    peerStore := raft.NewJSONPeers(v.DataDir, trans)
    peerStore.SetPeers(peers)

    r, err := raft.NewRaft(cfg, fsm, dbStore, dbStore, fileStore, peerStore, trans)
    if err != nil {
        log.Fatal(err)
    }

    t := time.NewTicker(time.Duration(1) * time.Second)

    for {
        select {
        case <-t.C:
            fmt.Println(r.Leader())
        }
    }
}
