package main

import (
	"flag"
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/ongyx/tamago"
)

const (
	width  = 500
	height = 500
)

var (
	window  *glfw.Window
	program uint32

	debug        bool
	rom, bootrom string
)

func init() {
	flag.BoolVar(&debug, "debug", false, "enter debug shell")
	flag.StringVar(&rom, "rom", "", "rom file")
	flag.StringVar(&bootrom, "bootrom", "", "bootrom file")
}

type OpenGLRenderer struct{}

func (rr OpenGLRenderer) Write(x, y int, c tamago.Colour) {}

func main() {
	/*
		runtime.LockOSThread()
		initGLFW()
		defer glfw.Terminate()
		initGL()
	*/

	cpu := tamago.NewCPU(OpenGLRenderer{})

	flag.Parse()

	if debug {
		cpu.DebugRun()
	}

	if bootrom != "" {
		cpu.LoadBoot(bootrom)
	}

	if rom != "" {
		cpu.Load(rom)
	}

	/*
		for !window.ShouldClose() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			window.SwapBuffers()
			glfw.PollEvents()
		}
	*/

	if err := cpu.Run(); err != nil {
		fmt.Println(err)
	}

}

func initGLFW() {
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
}

func initGL() {
	if err := gl.Init(); err != nil {
		fmt.Println("WARNING: gl.Init failed to load function: " + err.Error())
	}
	//version := gl.GoStr(gl.GetString(gl.VERSION))

	program = gl.CreateProgram()
	gl.LinkProgram(program)
	gl.UseProgram(program)
}
