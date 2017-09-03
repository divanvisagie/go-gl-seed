package main

import "github.com/go-gl/gl/v4.1-core/gl"

type DrawingObject struct {
	points []float32
	object uint32
}

/// makeVao make vertex array object
func MakeVertexArray(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

//Make a triangle to draw
func MakeTriangle() DrawingObject {

	triangle := []float32{
		0, 0.5, 0, //top
		-0.5, -0.5, 0, //left
		0.5, -0.5, 0, //right
	}

	vao := MakeVertexArray(triangle)
	return DrawingObject{triangle, vao}
}

//Make a square to draw
func MakeSquare(x, y, width, height float32) DrawingObject {
	square := []float32{
		x, y + height, 0,
		x, y, 0,
		x + width, y, 0,

		x, y + height, 0,
		x + width, y + height, 0,
		x + width, y, 0,
		// 0.5, 0.5, 0,
		// 0.5, -0.5, 0,
	}

	vao := MakeVertexArray(square)
	return DrawingObject{square, vao}
}
