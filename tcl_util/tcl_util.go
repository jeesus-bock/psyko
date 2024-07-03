package tcl_util

import (
	"fmt"
	"log"

	"modernc.org/tcl"
)

// RunTcl takes a TCL script and a data map and runs the script
// enhanced with a TCL dictionary object named data.
func RunTcl(script string, data map[string]interface{}) (ret string, err error) {
	log.Println("HandleTcl")
	dataDict := mapToTclDict(data)
	log.Println("dataDict: ", dataDict)
	i, err := tcl.NewInterp()
	if err != nil {
		return
	}
	ret, err = i.Eval(dataDict + string(script))
	if err != nil {
		return
	}
	err = i.Close()
	return
}

// mapToTclDict takes a map and creates a piece of TCL
// code defining a dictionary named data. All map values
// are reduced and cast to string in the process.
//
// For empty maps the function returns an empty string.
func mapToTclDict(data map[string]interface{}) (ret string) {
	log.Printf("mapToTclDict: %+v", data)
	if len(data) == 0 {
		return
	}
	ret = "set data [dict create "
	for k, v := range data {
		ret = ret + fmt.Sprintf(` %s "%v"`, k, v)
	}
	ret = ret + "]\n"
	return
}
