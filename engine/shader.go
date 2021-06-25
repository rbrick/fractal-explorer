// Shader is used to represent and work with GL Shaders in Go.
// Provides things like Uniforms, Shader models,
package engine

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
	"gonum.org/v1/gonum/mat"
)

//Program represents a GL Shader Program
type Program struct {

	// The GL ID for this program
	Id uint32
	// The Uniforms for this program
	Uniforms map[string]*Uniform

	// The shaders linked to this program
	Shaders []*Shader
}

func (p *Program) GetUniform(name string) *Uniform {
	if v, ok := p.Uniforms[name]; ok {
		return v
	}

	v := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(p.Id, v)

	uniform := &Uniform{
		Location: location,
	}

	return uniform
}

func (p *Program) Bind() {
	gl.UseProgram(p.Id)
}

func (p *Program) Unbind() {
	gl.UseProgram(0)
}

func (p *Program) Attach(s *Shader) {
	gl.AttachShader(p.Id, s.Id)
}

func (p *Program) Link() {
	gl.LinkProgram(p.Id)
}

//NewProgram creates a new shader program
func NewProgram() *Program {
	return &Program{
		Id:       gl.CreateProgram(),
		Uniforms: map[string]*Uniform{},
		Shaders:  []*Shader{},
	}
}

type Shader struct {
	Id     uint32
	Source string
	Status int32 // the compile status of this shader
}

//ReadShader returns a compiled new shader & an error if anything fails for to
// load or compile
func ReadShader(reader io.Reader, shaderType uint32, preprocessor *Preprocessor) (*Shader, error) {
	shaderId := gl.CreateShader(shaderType)

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	source := preprocessor.Process(strings.NewReader(string(data)))
	fmt.Println(source)
	src, free := gl.Strs(source + "\x00")

	gl.ShaderSource(shaderId, 1, src, nil)
	free()

	gl.CompileShader(shaderId)

	var status int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderId, logLength, nil, gl.Str(log))

		return &Shader{Status: status}, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return &Shader{
		Id:     shaderId,
		Source: source,
		Status: status,
	}, nil
}

type Uniform struct {
	Location int32
}

func (u *Uniform) Int(i int32) {
	gl.Uniform1i(u.Location, i)
}

func (u *Uniform) Float(f float32) {
	gl.Uniform1f(u.Location, f)
}

func (u *Uniform) Double(f float64) {
	gl.Uniform1d(u.Location, f)
}

func (u *Uniform) Vecf(vector ...float32) {
	switch len(vector) {
	case 2:
		gl.Uniform2f(u.Location, vector[0], vector[1])
	case 3:
		gl.Uniform3f(u.Location, vector[0], vector[1], vector[2])
	case 4:
		gl.Uniform4f(u.Location, vector[0], vector[1], vector[2], vector[3])
	}
}

//Matrix puts a matrix as an uniform for a given shader
func (u *Uniform) Matrix(m mat.Matrix, transpose bool) {
	x, y := m.Dims()

	matrix := make([]float32, x*y)

	// store the matrix in the array
	MatrixToArray(m, matrix)

	if x == y {
		switch x {
		case 2:
			gl.UniformMatrix2fv(u.Location, 4, transpose, &matrix[0])
		case 3:
			gl.UniformMatrix3fv(u.Location, 9, transpose, &matrix[0])
		case 4:
			gl.UniformMatrix4fv(u.Location, 16, transpose, &matrix[0])
		}
	}
}
