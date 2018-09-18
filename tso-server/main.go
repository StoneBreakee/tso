package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"./server"
	"github.com/ngaut/log"
)

var (
	addr     = flag.String("addr", "127.0.0.1:1234", "server listening address")
	zk       = flag.String("zk", "127.0.0.1:2181", "zookeeper address")
	rootPath = flag.String("root", "/zk/tso", "tso root path in zookeeper, must have the prefix /zk first")
	interval = flag.Int64("interval", 2000, "interval milliseconds to save timestamp in zookeeper, default is 2000(ms)")
	logLevel = flag.String("L", "debug", "log level: info, debug, warn, error, fatal")
)

func main() {
	flag.Parse()

	log.SetLevelByString(*logLevel)

	cfg := &server.Config{
		Addr:         *addr,
		ZKAddr:       *zk,
		RootPath:     *rootPath,
		SaveInterval: *interval,
	}

  // create a timestamp oracle
  // timestamp oracle includes : a server socket , listenning on cfg.addr
  //                             empty connection pool
	oracle, err := server.NewTimestampOracle(cfg)
	if err != nil {
		log.Fatal(err)
	}

  // start a http web ?
	go http.ListenAndServe(":5555", nil)

  // create a non-buffered channel to store system interrupted
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
	  // this goroutine will block
	  // if system interrupted occur,this goroutine will shutdown the tso
		sig := <-sc
		log.Infof("Got signal [%d] to exit.", sig)
		oracle.Close()
		os.Exit(0)
	}()

	oracle.Run()
}
