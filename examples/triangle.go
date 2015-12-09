package main

import (
	"fmt"
	"github.com/yulon/go-oswnd"
	"github.com/yulon/go-gl"
	"github.com/yulon/go-glrc"
)

func main() {
	oswnd.Factory(func() {
		wnd := oswnd.New()
		wnd.SetTitle("Triangle")
		wnd.SetClientSzie(oswnd.Size{800, 600})
		wnd.MoveToScreenCenter()

		rc := glrc.New(wnd.GetId())
		rc.Select()

		version := gl.GoStr(gl.GetString(gl.VERSION))
		fmt.Println("OpenGL version", version)

		gl.ClearColor(1, 1, 1, 0)

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

		wnd.OnSize = func() {
			cs := wnd.GetClientSzie()
			gl.Viewport(0, 0, int32(cs.Width), int32(cs.Height))
		}

		wnd.SetLayout(oswnd.LayoutVisible)
	})
}
