package util

import (
	"runtime"
	"strconv"
)

func PrintGoroutineStack() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	var s = ""
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		s += " "
		Logger.Error(s + "-> " + f.Name() + " in " + file + ":" + strconv.Itoa(line))
	}
}
