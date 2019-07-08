package cmd

import(
    "log"
    "encoding/hex"
    "encoding/binary"
)

func hexStringToLittleInt(value string) int {

    ilen := len(value)
    if ilen % 2 != 0{
        value = "0" + value
        ilen += 1    
    }

    for ;ilen < 8;{
        value = value + "00"
        ilen +=2
    }
    bs,err := hex.DecodeString(value)
    if err != nil{return -1}
    x := int32(binary.LittleEndian.Uint32(bs))

    return int(x)
}

func hexStringToBigInt(value string) int {

    ilen := len(value)
    if ilen % 2 != 0{
        value = "0"+value
        ilen += 1    
    }

    for ;ilen < 8;{
        value = "00"+value
        ilen +=2
    }

    log.Println(value,ilen)
    bs,err := hex.DecodeString(value)
    if err != nil{return -1}
    x := int32(binary.BigEndian.Uint32(bs))

    return int(x)
}
