package cmd

import(
    "log"
    "encoding/json"
)

type paramCfgSt struct {
    Tag string `json:"tag"`
    BaseLen string `json:"baselen"`
    Len int `json:"len"`
    Type string `json:"type"`
    Operate string `json:"operate"`
    Record string `json:"record"`
}

type cmdCfgSt struct{
    Name string `json:"name"`
    Id  string `json:"id"`
    Params []paramCfgSt `json:"params"`
}

type Cmd struct{
    cmdstr string
    cmdcfg cmdCfgSt
    valueMap map[string]string
}

func cfgProc(cfgMap map[string]interface{})(cmdCfgSt, bool){
    var cmdcfg cmdCfgSt

    cfgjson,err := json.Marshal(cfgMap)
    if err != nil{goto ERR}
    err = json.Unmarshal(cfgjson, &cmdcfg)
    if err != nil{goto ERR}
    
    return cmdcfg, true
ERR:
    println(err)
    return cmdcfg, false
    
}

func New(cmdstr string, cfgMap map[string]interface{})(*Cmd, bool){

   var cmd Cmd

   cmd.cmdstr = cmdstr
   cmdcfg, ok := cfgProc(cfgMap)
   if !ok{ return &cmd, ok}
   cmd.cmdcfg = cmdcfg
   cmd.valueMap = make(map[string]string)
   log.Printf("new analyz cmd %#v\n",cmd)

   return &cmd, ok
}

func RestCmdProc(cmdstr string, cfgMap map[string]interface{})(string, bool){

    var ok bool
    var info string

    log.Println(cfgMap)

    if cmdPtr,ok := New(cmdstr, cfgMap);ok{
        info,ok = cmdPtr.CmdAnalyze()
    }

    return info, ok
}

func init(){
    log.Println("cmd package running...")
}

