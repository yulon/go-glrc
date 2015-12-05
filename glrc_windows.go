package glrc

import (
	"syscall"
	"unsafe"
)

var (
	opengl32 = syscall.NewLazyDLL("opengl32.dll")
	user32 = syscall.NewLazyDLL("user32.dll")
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	getDC = user32.NewProc("GetDC").Call
	choosePixelFormat = gdi32.NewProc("ChoosePixelFormat").Call
	setPixelFormat = gdi32.NewProc("SetPixelFormat").Call
	wglCreateContext = opengl32.NewProc("wglCreateContext").Call
)

type pixelformatdescriptor struct{
	nSize uint16
	nVersion uint16
	dwFlags uint32
	iPixelType byte
	cColorBits byte
	cRedBits byte
	cRedShift byte
	cGreenBits byte
	cGreenShift byte
	cBlueBits byte
	cBlueShift byte
	cAlphaBits byte
	cAlphaShift byte
	cAccumBits byte
	cAccumRedBits byte
	cAccumGreenBits byte
	cAccumBlueBits byte
	cAccumAlphaBits byte
	cDepthBits byte
	cStencilBits byte
	cAuxBuffers byte
	iLayerType byte
	bReserved byte
	dwLayerMask uint32
	dwVisibleMask uint32
	dwDamageMask uint32
}

const (
	pfd_draw_to_window = 0x00000004
	pfd_support_opengl = 0x00000020
	pfd_doublebuffer = 0x00000001
	pfd_type_rgba = 0x000

	pfd_main_plane = 0x000
)

var pfd = pixelformatdescriptor{
	40,
	1,
	pfd_draw_to_window |
	pfd_support_opengl |
	pfd_doublebuffer,
	pfd_type_rgba,
	24,
	0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0,
	32,
	0, 0,
	pfd_main_plane,
	0,
	0, 0, 0,
}

func convDC(hDC uintptr) {
	PixelFormat, _, _ := choosePixelFormat(hDC, uintptr(unsafe.Pointer(&pfd)))
	setPixelFormat(hDC, PixelFormat, uintptr(unsafe.Pointer(&pfd)))
}

type RenderingContext struct{
	hDC uintptr
	hRC uintptr
}

func New(id uintptr) *RenderingContext {
	hDC, _, _ := getDC(id)
	convDC(hDC)
	hRC, _, _ := wglCreateContext(hDC)
	return &RenderingContext{hDC, hRC}
}

var wglMakeCurrent = opengl32.NewProc("wglMakeCurrent").Call

func (rc *RenderingContext) Select() {
	wglMakeCurrent(rc.hDC, rc.hRC)
}

var wglSwapBuffers = opengl32.NewProc("wglSwapBuffers").Call

func (rc *RenderingContext) Render() {
	wglSwapBuffers(rc.hDC)
}
