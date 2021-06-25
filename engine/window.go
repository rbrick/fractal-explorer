package engine

import (
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Hint func()

func ContextVersion(maj, min int) Hint {
	return func() {
		glfw.WindowHint(glfw.ContextVersionMajor, maj)
		glfw.WindowHint(glfw.ContextVersionMinor, min)
	}
}

func ContextProfile(profile int) Hint {
	return func() {
		glfw.WindowHint(glfw.OpenGLProfile, profile)
	}
}

func ContextForwardCompatible(b bool) Hint {
	return func() {
		if b {
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		} else {
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.False)
		}
	}
}

func WindowResizable(b bool) Hint {
	return func() {
		if b {
			glfw.WindowHint(glfw.Resizable, glfw.True)
		} else {
			glfw.WindowHint(glfw.Resizable, glfw.False)
		}
	}
}

func Decorated(b bool) Hint {
	return func() {
		if b {
			glfw.WindowHint(glfw.Decorated, glfw.True)
		} else {
			glfw.WindowHint(glfw.Decorated, glfw.False)
		}
	}
}

type Window struct {
	WindowHandle *glfw.Window
}

//Size returns the current width and height of the window
func (w *Window) Size() (int, int) {
	return w.WindowHandle.GetSize()
}

//FramebufferSize returns the current width and height of the framebuffer
func (w *Window) FramebufferSize() (int, int) {
	return w.WindowHandle.GetFramebufferSize()
}

func (w *Window) Init() error {
	w.WindowHandle.MakeContextCurrent()
	// initialize OpenGL
	if err := gl.Init(); err != nil {
		return err
	}
	return nil
}

func (w *Window) ShouldClose() bool {
	return w.WindowHandle.ShouldClose()
}

func NewWindow(name string, width, height int, monitor *glfw.Monitor, hints ...Hint) (*Window, error) {
	for _, hint := range hints {
		hint() // call the hints
	}
	// attempt to create a window
	window, err := glfw.CreateWindow(width, height, name, monitor, nil)

	if err != nil {
		return nil, err
	}
	return &Window{WindowHandle: window}, nil
}
