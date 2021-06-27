package engine

import "github.com/go-gl/gl/v3.2-core/gl"

type FrameBuffer struct {
	Id uint32
}

func (fb *FrameBuffer) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, fb.Id)
}

func (*FrameBuffer) Unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func CreateFrameBuffer() *FrameBuffer {
	var fbo uint32

	gl.GenFramebuffers(1, &fbo)
	return &FrameBuffer{
		Id: fbo,
	}
}
