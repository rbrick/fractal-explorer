package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/rbrick/fractal-explorer/engine"

	"github.com/go-gl/gl/v3.2-core/gl"
)

var (
	DefaultPreprocessor = &engine.Preprocessor{
		Directives: map[string]engine.ProcessFunc{
			"include": engine.Include,
		},
	}
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize:", err) // failed to initialize GLFW
	}

	w, err := engine.NewWindow("Fractal Explorer", 1200, 900, nil,
		engine.ContextVersion(3, 2),
		engine.ContextProfile(glfw.OpenGLCoreProfile),
		engine.ContextForwardCompatible(true),
		engine.WindowResizable(false),
		engine.Decorated(false),
	)

	if err != nil {
		log.Fatalln("failed to create window:", err)
	}

	panX, panY := 0., 0.
	zoom := 1.
	iterations := 200

	w.WindowHandle.SetKeyCallback(func(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		switch key {
		case glfw.KeyEscape:
			win.SetShouldClose(true)
		case glfw.KeyP:
			// increment the zoom
			zoom *= 1.25
		case glfw.KeyO:
			zoom /= 1.25
		case glfw.KeyUp:
			// Pan up
			panY += 0.01 / zoom
		case glfw.KeyDown:
			// Pan down
			panY -= 0.01 / zoom
		case glfw.KeyRight:
			// pan right
			panX += 0.01 / zoom
		case glfw.KeyLeft:
			// pan left
			panX -= 0.01 / zoom
		case glfw.KeyI:
			if glfw.ModShift == mods {
				iterations -= 50
			} else {
				iterations += 50
			}
		}
	})

	v, _ := os.Open("shaders/vertex.glsl")
	f, _ := os.Open("shaders/mandelbrot.glsl")

	glfw.WindowHint(glfw.Floating, glfw.True)

	w.Init()

	vert, err := engine.ReadShader(v, gl.VERTEX_SHADER, DefaultPreprocessor)

	if err != nil {
		log.Fatalln(err)
	}

	frag, err := engine.ReadShader(f, gl.FRAGMENT_SHADER, DefaultPreprocessor)

	if err != nil {
		fmt.Println("fail")
		log.Fatalln(err)
	}

	program := engine.NewProgram()
	// set up our shader program
	program.Bind()
	program.Attach(vert)
	program.Attach(frag)
	program.Link()
	program.Unbind()

	// setup our VAO
	vao := engine.CreateBuffer(engine.BufferType{})

	for !w.ShouldClose() {

		program.Bind()

		fw, fh := glfw.GetCurrentContext().GetFramebufferSize()

		program.GetUniform("Resolution").Vecf(float32(fw), float32(fh))
		program.GetUniform("Pan").Vecf(float32(panX), float32(panY))
		program.GetUniform("Time").Float(float32(glfw.GetTime()))
		program.GetUniform("Zoom").Float(float32(zoom))
		program.GetUniform("MaxIterations").Int(int32(iterations))

		vao.Bind(gl.ARRAY_BUFFER)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		vao.Unbind(gl.ARRAY_BUFFER)

		program.Unbind()

		w.WindowHandle.SwapBuffers()
		glfw.PollEvents()
	}
}
