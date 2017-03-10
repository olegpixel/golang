package main

import (
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320
	cells = 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {

	// open output file
	fo, err := os.Create("output.html")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	//buf := make([]byte, 1024)
	//for {
	//	// write a chunk
	//	if _, err := fo.Write(buf[:n]); err != nil {
	//		panic(err)
	//	}
	//}

	fo.WriteString("<svg xmlns='http://www.w3.org/2000/svf' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7; " +
		"width=" + strconv.Itoa(width) + "; height=" + strconv.Itoa(height) + "' >")
	for i:=0; i < cells; i++ {
		for j:=0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fo.WriteString("<polygon points='")
			fo.WriteString(FloatToString(ax) + ", " + FloatToString(ay))
			fo.WriteString(FloatToString(bx) + ", " + FloatToString(by))
			fo.WriteString(FloatToString(cx) + ", " + FloatToString(cy))
			fo.WriteString(FloatToString(dx) + ", " + FloatToString(dy))

			fo.WriteString("' />")
		}
	}
	fo.WriteString("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x,y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}