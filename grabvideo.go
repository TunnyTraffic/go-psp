package main

import (
        "fmt"
        "image/color"
        "gocv.io/x/gocv"
        "os"
)

func main() {
        //set source
        deviceID := os.Args[1]
        saveFile := os.Args[2]
        // deviceID := "http://root:pass@192.168.0.90/mjpg/video.mjpg"
        // saveFile := "testimage.jpg"
        webcam, err := gocv.OpenVideoCapture(deviceID)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer webcam.Close()
        img := gocv.NewMat()
        defer img.Close()

        if ok := webcam.Read(&img); !ok {
                fmt.Printf("Cannot Read device %v\n", deviceID)
                return
        }
        if img.Empty() {
                fmt.Printf("no image on device %v\n", deviceID)
                return
        }

        gocv.IMWrite(saveFile, img)


}
