package main

import (
	"context"
	"flag"
	"log"
	"os/user"
	"strconv"

	"github.com/jacobsa/fuse"
	memfs "github.com/jacobsa/fuse/samples/memfs"
)

var (
	mountpoint string
)

func main() {
	mountpoint := flag.String("mountpoint", "testdir", "set directory as mountpoint")
	flag.Parse()

	serve := memfs.NewMemFS(currentUid(), currentGid())
	cfg := &fuse.MountConfig{}

	mfs, err := fuse.Mount(*mountpoint, serve, cfg)
	if err != nil {
		log.Fatalf("Mount: %v", err)
	}

	if err := mfs.Join(context.Background()); err != nil {
		log.Fatalf("Join %v", err)
	}

}

func currentUid() uint32 {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	uid, err := strconv.ParseUint(user.Uid, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint32(uid)
}

func currentGid() uint32 {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	gid, err := strconv.ParseUint(user.Gid, 10, 32)
	if err != nil {
		panic(err)
	}

	return uint32(gid)
}
