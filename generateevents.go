/*
Detect faces, take a snapshot, write event data and filename to influxDB

via https://github.com/hybridgroup/gocv/blob/master/cmd/facedetect/main.go

classifier is in /home/go/src/gocv.io/x/gocv/data/
*/

package main

import (
	"fmt"
	"github.com/influxdata/influxdb/tree/1.8/client/v2"
	"time"
	"image"
	"image/color"
	"os"
	"gocv.io/x/gocv"
)

func main()
	if len(os.Args) <4 {
			fmt.Println("How to run:\n\tfacedetect [Camera ID] [classifier XML file] [saved image dir]")
			return
	}

	// parse args
	deviceID := os.Args[1]
	xmlFile := os.Args[2]
	imageLoc := os.Args[3]

	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// rectangle colour
	green := color.RGBA{0, 255, 0, 0}

	// load classifier
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image,
		// along with text identifing as "Human"
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}
	}
}
