package main

/*
#cgo windows CFLAGS: -DGLEW_STATIC
#cgo windows LDFLAGS: -static -lnanovg -lglew32 -lopengl32

#include <stdio.h>
#include <GL/glew.h>
#include <nanovg/nanovg.h>
#define NANOVG_GL2_IMPLEMENTATION
#include <nanovg/nanovg_gl.h>
*/
import "C"

import (
	"github.com/yulon/go-oswnd"
	"github.com/yulon/go-glrc"
)

func main() {
	wnd := oswnd.New()
	wnd.SetTitle("NanoVG")
	wnd.SetClientSzie(oswnd.Size{800, 600})
	wnd.MoveToScreenCenter()

	rc := glrc.New(wnd.GetDisplayId())
	rc.Select()
	C.glewInit()

	C.glClearColor(1, 1, 1, 0)
	vg := C.nvgCreateGL2(C.NVG_ANTIALIAS | C.NVG_STENCIL_STROKES)
	println(vg)

	wnd.OnSize = func() {
		cs := wnd.GetClientSzie()
		C.glViewport(0, 0, C.GLsizei(cs.Width), C.GLsizei(cs.Height))
	}

	wnd.OnPaint = func(){
		cs := wnd.GetClientSzie()
		C.nvgBeginFrame(vg, C.int(cs.Width), C.int(cs.Height), 1)
		C.glClear(C.GL_COLOR_BUFFER_BIT | C.GL_DEPTH_BUFFER_BIT)
		C.nvgBeginPath(vg)
		C.nvgRect(vg, 100, 100, 120, 30)
		C.nvgCircle(vg, 120, 120, 5)
		C.nvgPathWinding(vg, C.NVG_HOLE)
		C.nvgFillColor(vg, C.nvgRGBA(255, 192, 0, 255))
		C.nvgFill(vg)
		C.nvgEndFrame(vg)
		rc.Render()
	}

	wnd.Show()

	oswnd.BlockAndHandleEvents()
}
