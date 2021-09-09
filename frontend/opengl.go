package main

import (
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 160
	height = 144
)

var (
	window  *glfw.Window
	program uint32
)

func init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	w, err := glfw.CreateWindow(width, height, "tamago", nil, nil)
	if err != nil {
		panic(err)
	}
	w.MakeContextCurrent()

	window = w

	if err := gl.Init(); err != nil {
		panic(err)
	}
	//version := gl.GoStr(gl.GetString(gl.VERSION))

	program = gl.CreateProgram()
	gl.LinkProgram(program)
}
