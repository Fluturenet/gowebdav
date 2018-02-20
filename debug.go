
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var dbgChan = make(chan string, 8)
func init () {
	go func() {
		for {
			line := <-dbgChan
			log.Println(line)
		}
	}()
}

func dbgPrintf(format string, args ...interface{}) {
	dbgChan <- fmt.Sprintf(format, args...)
}

func dbgJson(obj interface{}) string {
	r, err := json.Marshal(obj)
	if err == nil {
		return string(r)
	}
	return fmt.Sprintf("%+v", obj)
}

type Debugger struct {
    Debpref string
}

func (m *Debugger) log (format string, v ...interface{}){
    format = m.Debpref+" "+format
        dbgPrintf(format,v...)
        }
