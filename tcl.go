package main

import (
	"fmt"
	"log"
	"os"

	"modernc.org/tcl"
)

func HandleTcl(filename string, data map[string]interface{}) (ret string, err error) {
	log.Println("HandleTcl: ", filename)
	dataDict := mapToTclDict(data)
	log.Println("dataDict: ", dataDict)
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	i, err := tcl.NewInterp()
	if err != nil {
		return
	}
	ret, err = i.Eval(dataDict + string(b))
	if err != nil {
		return
	}
	err = i.Close()
	return
}

func mapToTclDict(data map[string]interface{}) (ret string) {
	log.Printf("mapToTclDict: %+v", data)
	ret = "set data [dict create "
	for k, v := range data {
		ret = ret + fmt.Sprintf(` %s "%v"`, k, v)
	}
	ret = ret + "]\n"
	return
}
