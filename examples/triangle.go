package main

import (
	"fmt"
	"runtime"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/yulon/go-oswnd"
	"github.com/yulon/go-glrc"
)

func main() {
	runtime.LockOSThread()

	oswnd.Factory(func() {
		wnd := oswnd.New()
		wnd.SetTitle("Triangle")
		wnd.SetClientSzie(oswnd.Size{800, 600})
		wnd.MoveToScreenCenter()

		rc := glrc.New(wnd.GetId())
		rc.Select()

		if err := gl.Init(); err != nil {
			panic(err)
		}

		version := gl.GoStr(gl.GetString(gl.VERSION))
		fmt.Println("OpenGL version", version)

		wnd.OnPaint = func(){
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			gl.Begin(gl.TRIANGLES)

			gl.Color3f(1.0, 0.0, 0.0)
			gl.Vertex3f(0.0, 1.0, 0.0)

			gl.Color3f(0.0, 1.0, 0.0)
			gl.Vertex3f(-1.0, -1.0, 0.0)

			gl.Color3f(0.0, 0.0, 1.0)
			gl.Vertex3f(1.0, -1.0, 0.0)  

			gl.End()

			rc.Render()
		}

		wnd.SetLayout(oswnd.LayoutVisible)
	})
}
