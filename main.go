package main

import (
	"log"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/rbrick/fractal-explorer/engine"

	"github.com/go-gl/gl/v3.2-core/gl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize:", err) // failed to initialize GLFW
	}

	w, err := engine.NewWindow("Fractal Explorer", 800, 800, nil,
		engine.ContextVersion(3, 2),
		engine.ContextProfile(glfw.OpenGLCoreProfile),
		engine.ContextForwardCompatible(true),
		engine.WindowResizable(true),
		engine.Decorated(false),
	)

	if err != nil {
		log.Fatalln("failed to create window:", err)
	}

	w.WindowHandle.SetKeyCallback(func(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape {
			win.SetShouldClose(true)
		}
	})

	f, _ := os.Open("mandelbrot.glsl")
	p := &engine.Preprocessor{
		Directives: map[string]engine.ProcessFunc{
			"include": engine.Include,
		},
	}

	glfw.WindowHint(glfw.Floating, glfw.True)

	// w.WindowHandle.Maximize()
	w.WindowHandle.Iconify()
	w.Init()

	_, err = engine.ReadShader(f, gl.FRAGMENT_SHADER, p)

	if err != nil {
		log.Fatalln(err)
	}

	for !w.ShouldClose() {
		w.WindowHandle.SwapBuffers()

		glfw.PollEvents()
	}
}
