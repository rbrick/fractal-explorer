package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/mazznoer/colorgrad"
	"github.com/rbrick/fractal-explorer/engine"

	"github.com/go-gl/gl/v4.1-core/gl"
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
		engine.ContextVersion(4, 1),
		engine.ContextProfile(glfw.OpenGLCoreProfile),
		engine.ContextForwardCompatible(true),
		engine.WindowResizable(true),
		engine.Decorated(true),
	)

	if err != nil {
		log.Fatalln("failed to create window:", err)
	}

	panX, panY := 0., 0.
	zoom := 1.
	iterations := 200
	mode := 0

	w.WindowHandle.SetKeyCallback(func(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press || action == glfw.Repeat {
			switch key {
			case glfw.KeyEscape:
				win.SetShouldClose(true)
			case glfw.KeyP:
				// increment the zoom
				zoom *= 1.05
			case glfw.KeyO:
				zoom /= 1.05
			case glfw.KeyUp:
				// Pan up
				panY -= 0.02 / zoom
			case glfw.KeyDown:
				// Pan down
				panY += 0.02 / zoom
			case glfw.KeyRight:
				// pan right
				panX += 0.02 / zoom
			case glfw.KeyLeft:
				// pan left
				panX -= 0.02 / zoom
			case glfw.KeyI:
				if glfw.ModShift == mods {
					iterations -= 50
				} else {
					iterations += 50
				}
			case glfw.KeyM:
				mode += 1

				if mode > 1 {
					mode = 0
				}
			}

			fmt.Println("panX", panX, "panY", panY, "zoom", zoom)
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
		program.GetUniform("Mode").Int(int32(mode))

		var colors []int32

		for _, c := range colorgrad.Blues().ColorfulColors(256) {

			r, g, b := c.RGB255()

			colors = append(colors, ((int32(r)&0x0ff)<<16)|((int32(g)&0x0ff)<<8)|(int32(b)&0x0ff))
		}

		program.GetUniform("palette").IntArray(colors)

		vao.Bind(gl.ARRAY_BUFFER)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		vao.Unbind(gl.ARRAY_BUFFER)

		program.Unbind()

		w.WindowHandle.SwapBuffers()
		glfw.PollEvents()
	}
}
