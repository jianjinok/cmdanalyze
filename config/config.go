package config

import(
    "log"
    "flag"
)

const version = "v0.1"

var(
    RestAddr string
)

func RUN(){
    flag.Parse()
    log.Printf("tcp server addr: %s", RestAddr)
}

func init(){
    log.Printf("cmd analyze service version: %s", version)
    log.SetFlags(log.Llongfile | log.LstdFlags)
    flag.StringVar(&RestAddr, "R",":4444", "rest server addr")
}

