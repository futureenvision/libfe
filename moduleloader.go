package libfe

import (
	"syscall/js"
)

func (libfe *LIBFE) Load(url string) {
	js.Global().Call("moduleLoader", url)
	libfe.Println("Starting ...")
}
