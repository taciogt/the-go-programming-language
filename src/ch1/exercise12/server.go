// handler echoes the http request

// server up: go run . main.go &
// test: go run src/ch1/fetch/main.go http://localhost:8000/hello
// test: go run fetch/main.go http://localhost:8000/hello\?query\=123
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/lissajous", lissajousHandler) // each request call handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		log.Print(err)
	}

	lissajous(w, size)
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, size int) {
	if size == 0 {
		size = 100 // image canvas covers [-size..+size]
	}
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
