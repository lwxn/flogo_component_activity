package sample

import (
	"fmt"
	"github.com/project-flogo/core/activity"
	"log"
	"gocv.io/x/gocv"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func init() {
	_ = activity.Register(&VideoDecode{},New) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &VideoDecode{} //a new instance
	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type VideoDecode struct {
	VideoSrcPath string
	FrameDesPath string
}

// Metadata returns the activity's metadata
func (a *VideoDecode) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *VideoDecode) Eval(ctx activity.Context) (done bool, err error) {
	//init the activity with the source path and destination path
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	a.VideoSrcPath = input.VideoSrcPath
	a.FrameDesPath = input.FrameDesPath

	ctx.Logger().Debugf("Video path: %s", input.VideoSrcPath)
	ctx.Logger().Debugf("Frame path: %s",input.FrameDesPath)

	// read the video
	vc,err := gocv.OpenVideoCapture(a.VideoSrcPath)
	fmt.Println("-------------")
	if err != nil{
		log.Fatal("open the video capture failed:",err)
		return false,err
	}
	defer vc.Close()
	if !vc.IsOpened(){
		log.Fatal("vc is not open")
		return false,err
	}

	//insure the frameDesPath exists,otherwise create it
	_,err = os.Stat(a.FrameDesPath)
	if err != nil && os.IsNotExist(err){
		err = os.MkdirAll(a.FrameDesPath,os.ModePerm)
		if err != nil{
			log.Println("FrameDesPath doesn't exist,create it...")
			fmt.Println(err)
		}
	}

	//get the fileName
	videoNameWithSuffix := path.Base(a.VideoSrcPath)
	videoSuffix := path.Ext(a.VideoSrcPath)
	videoName := strings.TrimSuffix(videoNameWithSuffix,videoSuffix)

	//set the number off hte video
	frameCount := vc.Get(gocv.VideoCaptureFrameCount)
	fps := vc.Get(gocv.VideoCaptureFPS)
	fmt.Println(gocv.VideoCaptureFPS)
	fmt.Println(fps)
	duration := frameCount / fps
	//fmt.Printf("VideoCaptureFPS : %lf , fps : %lf",gocv.VideoCaptureFPS,fps)

	//get the frame
	frame := gocv.NewMat()
	defer frame.Close()
	timePerFrame := 1    // timePerFrame is the cost time per frame (frame/s)
	frameIsCreate := false
	fmt.Println(duration)
	for i:= 0; i< int(duration);i+=timePerFrame{
		vc.Set(gocv.VideoCapturePosFrames,(float64(i)/duration)*frameCount)
		vc.Read(&frame)
		if frame.Empty(){
			break
		}
		framePath := fmt.Sprintf("%s-%d.jpg",videoName,i)
		fmt.Println(framePath)
		gocv.IMWrite(filepath.Join(input.FrameDesPath,framePath),frame)
		frameIsCreate = true
	}
	//frames := 2
	//vc.Set(gocv.VideoCapturePosFrames,float64(frames))
	//frame := gocv.NewMat()
	//defer frame.Close()
	//
	//vc.Read(&frame)

	//if frame.Empty(){
	//	log.Fatal("frame is empty")
	//}
	//gocv.IMWrite(filepath.Join(input.FrameDesPath,"1.jpg"),frame)
	//gocv.IMWrite(input.FileDesPath,img)
	if frameIsCreate == false{
		log.Fatal("No frame is created...")
		return false,nil
	}else{
		log.Printf("There are %d frames",int(duration)/timePerFrame)
	}

	output := &Output{
		Result: "OK!",
	}
	err = ctx.SetOutputObject(output)
	if err != nil{
		return true,nil
	}

	return true, nil
}
