package tcl_util

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"modernc.org/tcl"
)

var inter *tcl.Interp

func SetupTcl() (err error) {
	inter, err = tcl.NewInterp()
	addEcho()
	return
}

// RunTcl takes a TCL script and a data map and runs the script
// enhanced with a TCL dictionary object named data.
func RunTcl(script string, data map[string]interface{}) (ret string, err error) {
	log.Println("HandleTcl")
	dataDict := mapToTclDict(data)
	log.Println("dataDict: ", dataDict)
	if err != nil {
		return
	}
	ret, err = inter.Eval(dataDict + string(script))
	if err != nil {
		return
	}
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

func GetTclEndPoints() (ret map[string]string) {
	ret = make(map[string]string)
	files, err := os.ReadDir("scripts")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Println("TCL script: ", file.Name())
		b, err := os.ReadFile("scripts/" + file.Name())
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(bytes.NewReader(b))
		scanner.Scan()
		ret["/tcl/"+strings.TrimSuffix(file.Name(), ".tcl")] = strings.TrimPrefix(scanner.Text(), "# ")
	}
	return
}

func addEcho() {
	log.Printf("%+v", inter)
	cmd, err := inter.NewCommand(
		"::go::echo",
		func(clientData interface{}, in *tcl.Interp, args []string) int {
			log.Println("added echo func")
			// Go implementation of the Tcl ::go::echo command
			args = append(args[1:], fmt.Sprint(clientData))
			in.SetResult(strings.Join(args, " "))
			return 0
		},
		42, // client data
		func(clientData interface{}) {
			// Go implemetation of the command delete handler
			log.Printf("%+v", clientData)
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%+v", cmd)
}
