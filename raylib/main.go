package raylib

/*
#cgo android LDFLAGS: -llog -landroid -lEGL -lGLESv2 -lOpenSLES -lm
#cgo android CFLAGS: -DPLATFORM_ANDROID -DGRAPHICS_API_OPENGL_ES2 -DMAL_NO_SDL -Iexternal/android/native_app_glue

#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation
#cgo darwin CFLAGS: -x objective-c -Iexternal/glfw/include -D_GLFW_COCOA -D_GLFW_USE_CHDIR -D_GLFW_USE_MENUBAR -D_GLFW_USE_RETINA -Wno-deprecated-declarations -DPLATFORM_DESKTOP -DMAL_NO_COREAUDIO

#cgo darwin,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo darwin,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo darwin,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33

#cgo linux,arm LDFLAGS: -L/opt/vc/lib -L/opt/vc/lib64 -lbrcmGLESv2 -lbrcmEGL -lpthread -lrt -lm -lbcm_host -lvcos -lvchiq_arm -ldl
#cgo linux,arm CFLAGS: -DPLATFORM_RPI -DGRAPHICS_API_OPENGL_ES2 -I/opt/vc/include -I/opt/vc/include/interface/vcos -I/opt/vc/include/interface/vmcs_host/linux -I/opt/vc/include/interface/vcos/pthreads

#cgo linux CFLAGS: -Iexternal/glfw/include -DPLATFORM_DESKTOP -Wno-stringop-overflow

#cgo linux,!wayland LDFLAGS: -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux,wayland LDFLAGS: -lGL -lm -pthread -ldl -lrt -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon

#cgo linux,!wayland CFLAGS: -D_GLFW_X11
#cgo linux,wayland CFLAGS: -D_GLFW_WAYLAND

#cgo linux,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo linux,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo linux,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33

#cgo windows LDFLAGS: -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -D_GLFW_WIN32 -Iexternal -Iexternal/glfw/include -Iexternal/glfw/deps/mingw -DPLATFORM_DESKTOP -DRAYGUI_IMPLEMENTATION

#cgo windows,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo windows,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo windows,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33

#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

// Color type, RGBA (32bit)
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// cptr returns C pointer
func (color *Color) cptr() *C.Color {
	return (*C.Color)(unsafe.Pointer(color))
}

const (
	// Set to show raylib logo at startup
	FlagShowLogo = 1
	// Set to run program in fullscreen
	FlagFullscreenMode = 2
	// Set to allow resizable window
	FlagWindowResizable = 4
	// Set to disable window decoration (frame and buttons)
	FlagWindowUndecorated = 8
	// Set to allow transparent window
	FlagWindowTransparent = 16
	// Set to try enabling MSAA 4X
	FlagMsaa4xHint = 32
	// Set to try enabling V-Sync on GPU
	FlagVsyncHint = 64
)

func init() {
	//We have to lock the OS Thread as raylib is sensitive to that stuff.
	//runtime.LockOSThread()
}

// SetConfigFlags - Setup some window configuration flags
func SetConfigFlags(flags int) {
	cflags := (C.uint)(flags)
	C.SetConfigFlags(cflags)
}

//WindowShouldClose Should the window clsoe
func WindowShouldClose() bool {
	return bool(C.WindowShouldClose())
}

// InitWindow - Initialize Window and OpenGL Graphics
func InitWindow(width int32, height int32, title string) {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)

	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	C.InitWindow(cwidth, cheight, ctitle)
}

func BeginDrawing() {
	runtime.LockOSThread()
	C.BeginDrawing()
}

func EndDrawing() {
	C.EndDrawing()
	runtime.UnlockOSThread()
}

func ClearBackground(color Color) {
	clr := *color.cptr()
	C.ClearBackground(clr)
}

func DrawText(label string, x int32, y int32, w int32, color Color) {
	ctext := C.CString(label)
	clr := *color.cptr()
	cx := C.int(x)
	cy := C.int(y)
	cw := C.int(w)

	defer C.free(unsafe.Pointer(ctext))
	C.DrawText(ctext, cx, cy, cw, clr)
}

func CloseWindow() {
	C.CloseWindow()
}

func TestWorld() {
	InitWindow(800, 400, "Hello world!")

	color := Color{R: 255, G: 0, B: 255, A: 255}
	color2 := Color{R: 255, G: 255, B: 255, A: 255}

	for !WindowShouldClose() {
		BeginDrawing()
		ClearBackground(color2)
		DrawText("How are you today? ", 10, 10, 20, color)
		EndDrawing()
	}

	CloseWindow()
}

func HelloWorld() {
	fmt.Println("Testing")
	cs := C.CString("Hello from stdio")
	label := C.CString("Congrats! First WIndow!")

	fmt.Println("Unsafe String", cs)

	//defer C.free(cs)
	//defer C.free(label)

	color := Color{R: 255, G: 0, B: 255, A: 255}
	color2 := Color{R: 255, G: 255, B: 255, A: 255}
	C.InitWindow(800, 400, cs)

	for !C.WindowShouldClose() {
		C.BeginDrawing()
		C.ClearBackground(*color2.cptr())
		C.DrawText(label, 10, 10, 20, *color.cptr())
		C.EndDrawing()
	}

	C.CloseWindow()

}
