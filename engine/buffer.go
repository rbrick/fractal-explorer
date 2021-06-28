package engine

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type ComponentType uint16

const (
	Position ComponentType = iota
	Color
	Texture
)

type (
	VertexComponent struct {
		Length int
		Type   reflect.Type
		GLType uint32
	}

	BufferType struct {
		Components []VertexComponent
	}
)

func (b BufferType) Size() int {
	length := 0

	for _, c := range b.Components {
		length += c.Length * int(c.Type.Size())
	}
	return length
}

var (
	PositionComponent VertexComponent = VertexComponent{Length: 3, Type: reflect.TypeOf(float32(0)), GLType: gl.FLOAT}
	ColorComponent    VertexComponent = VertexComponent{Length: 3, Type: reflect.TypeOf(float32(0)), GLType: gl.FLOAT}
	TextureComponent  VertexComponent = VertexComponent{Length: 2, Type: reflect.TypeOf(float32(0)), GLType: gl.FLOAT}
)

type Vertex interface {
	PosX() float32
	PosY() float32
	PosZ() float32

	Write(buf *Buffer)
}

type BaseVertex struct {
	x, y, z float32
}

func (v *BaseVertex) PosX() float32 {
	return v.x
}

func (v *BaseVertex) PosY() float32 {
	return v.y
}

func (v *BaseVertex) PosZ() float32 {
	return v.z
}

func (v *BaseVertex) Write(buf *Buffer) {
	binary.Write(buf.buf, binary.LittleEndian, v.x)
	binary.Write(buf.buf, binary.LittleEndian, v.y)
	binary.Write(buf.buf, binary.LittleEndian, v.z)
}

func NewVertex(x, y, z float32) Vertex {
	return &BaseVertex{x, y, z}
}

type Buffer struct {
	BufId, ArrayId uint32
	bufferType     BufferType
	buf            *bytes.Buffer

	Vertices []Vertex
}

func (buf *Buffer) Vertex(vertex Vertex) {
	vertex.Write(buf)
	buf.Vertices = append(buf.Vertices, vertex)
}

func (b *Buffer) Bind(target uint32) {
	gl.BindVertexArray(b.ArrayId)
	gl.BindBuffer(target, b.BufId)
}

func (b *Buffer) Unbind(target uint32) {
	gl.BindBuffer(target, 0)
	gl.BindVertexArray(0)
}

func (b *Buffer) Upload(target uint32) {
	offset := 0

	for i, e := range b.bufferType.Components {
		gl.EnableVertexAttribArray(uint32(i))
		gl.VertexAttribPointer(uint32(i), int32(e.Length), e.GLType, false, int32(e.Type.Size())*int32(e.Length)*2, gl.PtrOffset(offset))
		offset += e.Length * int(e.Type.Size())
	}

	gl.BufferData(target, len(b.buf.Bytes()), gl.Ptr(b.buf.Bytes()), gl.STATIC_DRAW)
}

func (b *Buffer) Draw(glType uint32) {
	gl.BindVertexArray(b.ArrayId)
	gl.DrawArrays(glType, 0, int32(len(b.Vertices)))
}

func CreateBuffer(bufferType BufferType) *Buffer {
	var bufId, arrayId uint32
	gl.GenBuffers(1, &bufId)
	gl.GenVertexArrays(1, &arrayId)

	return &Buffer{
		BufId:      bufId,
		ArrayId:    arrayId,
		bufferType: bufferType,
		buf:        bytes.NewBuffer([]byte{}),
		Vertices:   make([]Vertex, 0),
	}
}
