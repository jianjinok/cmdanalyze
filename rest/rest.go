package rest

import (
    "log"
    "io/ioutil"
    "net/http"
    "cmdanalyze/cmd"
    "github.com/ant0ine/go-json-rest/rest"
    "encoding/json"
)

var url []string = []string{"/cmd/analyze","/service/getstatus"}
var routes = [] *rest.Route{
    rest.Post(url[0], cmd_analyze),
    rest.Get(url[1], service_getstatus),
}

type requestSt struct{
    Cmd string   `json:"cmd"`
    Cfg map[string]interface{}  `json:"cfg"`
}

type responseSt struct{
    Msg string  `json:"msg"`
    Status string   `json:"status"`
    Info string     `json:"info"`
}

func service_getstatus(w rest.ResponseWriter, req *rest.Request){
    var response responseSt
    response.Msg = "no support"

    w.WriteJson(response)
}

func cmd_analyze(w rest.ResponseWriter, req *rest.Request){

    var request requestSt

    response := responseSt{Status:"ok"}
    jsonbytes, _ := ioutil.ReadAll(req.Body)
    err := json.Unmarshal(jsonbytes, &request)
    if err!=nil || len(request.Cmd)==0{
        log.Println(err)
        log.Println("jsonbytes",string(jsonbytes))
        response.Status = "fail"
        response.Msg = "format error"
    }else{
        log.Println(request)
        recvmsg, ok := cmd.RestCmdProc(request.Cmd, request.Cfg)
        if !ok{
            response.Status = "fail"
            response.Info = recvmsg
            recvmsg = ""
        }
        response.Msg = recvmsg
    }

    w.WriteJson(response)
}

func restserver(addr string){
    
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(routes...)
    if err != nil{
        log.Fatal(err)
    }
    api.SetApp(router)
    http.Handle("/", http.StripPrefix("", api.MakeHandler()))
    log.Fatal(http.ListenAndServe(addr,nil))
}

func RUN(addr string){
    log.Printf("start rest server %s", addr)
    restserver(addr)
}

func init(){
    log.Printf("rest server running...\n")
    log.Println("url: ",url)
}

