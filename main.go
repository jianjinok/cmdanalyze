package main

import(
    "cmdanalyze/rest"
    "cmdanalyze/config"
)

func main(){
    config.RUN()
    rest.RUN(config.RestAddr)
}
