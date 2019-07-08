package cmd

import(
    "fmt"
    "log"
    "encoding/json"
)

func(cmd *Cmd)cmdGetValueLen(param paramCfgSt) int{

    length := param.Len
    if len(param.BaseLen) > 0{
        value,ok := cmd.valueMap[param.BaseLen] 
        if !ok{return -1}
        baseLen := hexStringToLittleInt(value)
        log.Printf("baselen:%d len%d\n",baseLen, length)
        if baseLen < 0{return -1}
        length += int(baseLen)
    }

    return 2*length
}

func(cmd *Cmd)cmdGetCmdValue(param paramCfgSt)(string, bool){

    cmdstr := cmd.cmdstr
    analyzeLen := cmd.cmdGetValueLen(param)
    if analyzeLen < 0{
        return "get len error",false
    }
    if analyzeLen > len(cmdstr) {
        return fmt.Sprintf("analyze len %d cmd len:%d error\n", analyzeLen, len(cmdstr)), false
    }
    value := cmdstr[: analyzeLen]
    cmdstr = cmdstr[analyzeLen:]
    cmd.cmdstr = cmdstr

    return value, true
}

func (cmd *Cmd)cmdDataTypeProc(cmdstr string, param paramCfgSt)(string,bool){

    var cmdjson string
    cmdjson = cmdstr
    return cmdjson, true
}

func (cmd *Cmd)cmdDataOperateProc(cmdstr string, param paramCfgSt)(string, bool){

    var cmdjson string
    cmdjson = cmdstr
    return cmdjson, true
}

func (cmd *Cmd)cmdDataRecordProc(cmdstr string, param paramCfgSt)(string, bool){

    var cmdjson string
    cmdjson = cmdstr
    return cmdjson, true
}

func (cmd *Cmd)CmdAnalyze()(string,bool){
    params := cmd.cmdcfg.Params
    var ok bool
    var value string

    for _, param := range params{
        value,ok = cmd.cmdGetCmdValue(param)
        if !ok{return value,ok}
        value,ok = cmd.cmdDataTypeProc(value, param)
        if !ok{return value,ok}
        value,ok = cmd.cmdDataOperateProc(value, param)
        if !ok{return value,ok}
        value,ok = cmd.cmdDataRecordProc(value, param)
        if !ok{return value,ok}

        cmd.valueMap[param.Tag] = value
    }

    info, err := json.Marshal(cmd.valueMap)   
    if err != nil{
        log.Println(err)
        ok = false
    }
    log.Println(string(info),ok)
    return string(info),ok 
}

