package glrc

import (
	"syscall"
	"unsafe"
)

var (
	opengl32, _ = syscall.LoadLibrary("opengl32.dll")
	user32, _ = syscall.LoadLibrary("user32.dll")
	gdi32, _ = syscall.LoadLibrary("gdi32.dll")

	getDC, _ = syscall.GetProcAddress(user32, "GetDC")
	choosePixelFormat, _ = syscall.GetProcAddress(gdi32, "ChoosePixelFormat")
	setPixelFormat, _ = syscall.GetProcAddress(gdi32, "SetPixelFormat")
	wglCreateContext, _ = syscall.GetProcAddress(opengl32, "wglCreateContext")
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
	PixelFormat, _, _ := syscall.Syscall(choosePixelFormat, 2, hDC, uintptr(unsafe.Pointer(&pfd)), 0)
	syscall.Syscall(setPixelFormat, 3, hDC, PixelFormat, uintptr(unsafe.Pointer(&pfd)))
}

type RenderingContext struct{
	hDC uintptr
	hRC uintptr
}

func New(id uintptr) *RenderingContext {
	hDC, _, _ := syscall.Syscall(getDC, 1, id, 0, 0)
	convDC(hDC)
	hRC, _, _ := syscall.Syscall(wglCreateContext, 1, hDC, 0, 0)
	return &RenderingContext{hDC, hRC}
}

var wglMakeCurrent, _ = syscall.GetProcAddress(opengl32, "wglMakeCurrent")

func (rc *RenderingContext) Select() {
	syscall.Syscall(wglMakeCurrent, 2, rc.hDC, rc.hRC, 0)
}

var wglSwapBuffers, _ = syscall.GetProcAddress(opengl32, "wglSwapBuffers")

func (rc *RenderingContext) Render() {
	syscall.Syscall(wglSwapBuffers, 1, rc.hDC, 0, 0)
}
