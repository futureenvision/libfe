package libfe

import (
	"syscall/js"
	"time"
)

type LIBFE struct {
	allInputs   []interface{}
	inputCache  string
	inputBUFFER string
	dirlink     string
}

func (libfe *LIBFE) Print(input ...interface{}) {
	if len(libfe.allInputs) > 80 {
		libfe.allInputs = make([]interface{}, 0)
	}
	libfe.allInputs = append(libfe.allInputs, input...)
	js.Global().Call("Println", libfe.allInputs)
}

func (libfe *LIBFE) Println(input ...interface{}) {
	libfe.Print("\n", input)
}

func (libfe *LIBFE) PrintCl(input ...interface{}) {
	js.Global().Call("Println", append(libfe.allInputs, libfe.dirlink, input))
}

func (libfe *LIBFE) setScan(this js.Value, inputs []js.Value) interface{} {
	libfe.inputCache = inputs[0].String()
	if libfe.inputCache == "F1" ||
		libfe.inputCache == "F2" ||
		libfe.inputCache == "F3" ||
		libfe.inputCache == "F4" ||
		libfe.inputCache == "F5" ||
		libfe.inputCache == "F6" ||
		libfe.inputCache == "F7" ||
		libfe.inputCache == "F8" ||
		libfe.inputCache == "F9" ||
		libfe.inputCache == "F10" ||
		libfe.inputCache == "F11" ||
		libfe.inputCache == "F12" ||
		libfe.inputCache == "Shift" ||
		libfe.inputCache == "Meta" ||
		libfe.inputCache == "Enter" ||
		libfe.inputCache == "Control" ||
		libfe.inputCache == "Alt" {
	} else {
		libfe.inputBUFFER += libfe.inputCache
	}
	libfe.PrintCl(libfe.inputBUFFER + "_")
	return 0
}

func (libfe *LIBFE) Scan(dirlink string) string {
	libfe.dirlink = dirlink
	js.Global().Set("Scan", js.FuncOf(libfe.setScan))
	js.Global().Call("isScan", true)
	js.Global().Call("Scan", "")
	for {
		if libfe.inputCache != "" {
			if libfe.inputCache == "Enter" {
				js.Global().Call("isScan", false)
				var output = libfe.inputBUFFER
				libfe.inputBUFFER = ""
				libfe.inputCache = ""
				return output
			}
		}
		time.Sleep(1 * time.Millisecond)
	}
}
