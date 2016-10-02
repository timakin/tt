package tt

import (
	"log"
	"runtime"
	"time"
)

func Trace() (string, time.Time) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	funcName := f.Name()
	log.Println("START: ", funcName)
	return funcName, time.Now()
}

func Untrace(funcName string, startTime time.Time) {
	endTime := time.Now()
	log.Println("END:   ", funcName, "ElapsedTime in seconds:", endTime.Sub(startTime))
}
