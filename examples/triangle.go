package main

/*
#cgo windows CFLAGS: -DGLEW_STATIC
#cgo windows LDFLAGS: -static -lglew32 -lopengl32
#include <GL/glew.h>
*/
import "C"

import (
	"fmt"
	"github.com/yulon/go-oswnd"
	"github.com/yulon/go-glrc"
	"unsafe"
)

func main() {
	wnd := oswnd.New()
	wnd.SetTitle("Triangle")
	wnd.SetClientSzie(oswnd.Size{800, 600})
	wnd.MoveToScreenCenter()

	rc := glrc.New(wnd.GetDisplayId())
	rc.Select()
	C.glewInit()

	fmt.Println("OpenGL version", C.GoString((*C.char)(unsafe.Pointer(C.glGetString(C.GL_VERSION)))))

	C.glClearColor(1, 1, 1, 0)

	wnd.OnPaint = func() {
		C.glClear(C.GL_COLOR_BUFFER_BIT | C.GL_DEPTH_BUFFER_BIT)

		C.glBegin(C.GL_TRIANGLES)

		C.glColor3f(1.0, 0.0, 0.0)
		C.glVertex3f(0.0, 1.0, 0.0)

		C.glColor3f(0.0, 1.0, 0.0)
		C.glVertex3f(-1.0, -1.0, 0.0)

		C.glColor3f(0.0, 0.0, 1.0)
		C.glVertex3f(1.0, -1.0, 0.0)

		C.glEnd()

		rc.Render()
	}

	wnd.OnSize = func() {
		cs := wnd.GetClientSzie()
		C.glViewport(0, 0, C.GLsizei(cs.Width), C.GLsizei(cs.Height))
	}

	wnd.Show()

	oswnd.BlockAndHandleEvents()
}
