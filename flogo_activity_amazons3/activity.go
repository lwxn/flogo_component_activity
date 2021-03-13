package sample

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const(
	inAction             = "action"
	inAwsAccessKeyID     = "awsAccessKeyID"
	inAwsSecretAccessKey = "awsSecretAccessKey"
	inAwsRegion          = "awsRegion"
	inS3BucketName       = "s3BucketName"
	inLocalLocation      = "localLocation"
	inS3Location         = "s3Location"
	inS3NewLocation      = "s3NewLocation"
	otResult             = "result"
)

var log = logger.GetLogger("amazon_s3 activity")

// Activity is an sample Activity that can be used as a base to create a custom activity
type MyActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity{
	act := &MyActivity{metadata: metadata}
	return act
}

// Metadata returns the activity's metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *MyActivity) Eval(ctx activity.Context) (done bool, err error) {

	//get the action
	fmt.Println("---------------Eval begin----------------")
	fmt.Println(inAction)
	action := ctx.GetInput(inAction).(string)
	awsRegion := ctx.GetInput(inAwsRegion).(string)
	s3BucketName := ctx.GetInput(inS3BucketName).(string)
	localLocation := ctx.GetInput(inLocalLocation).(string)
	s3Location := ctx.GetInput(inS3Location).(string)
	s3NewLocation := ctx.GetInput(inS3NewLocation).(string)

	//get the Id and key
	awsAccessKeyID := ctx.GetInput(inAwsAccessKeyID).(string)
	awsSecretAccessKey := ctx.GetInput(inAwsSecretAccessKey).(string)

	// create a session with credentials
	awsSession := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			awsAccessKeyID,awsSecretAccessKey,""),
		Region: aws.String(awsRegion),
	}))

	var action_err error
	if action == "download"{
		action_err = downloadFileFromS3(awsSession,localLocation,s3BucketName,s3Location)
	}else if action == "upload"{
		action_err = uploadFileToS3(awsSession,localLocation,s3BucketName,s3Location)
	}else if action == "copy"{
		action_err = copyFileToS3(awsSession,s3BucketName,s3Location,s3NewLocation)
	}else if action == "delete"{
		action_err = deleteFileFromS3(awsSession,s3BucketName,s3Location)
	}

	if action_err != nil{
		ctx.SetOutput(otResult,action_err.Error())
		return true,nil
	}


	ctx.SetOutput(otResult,"ok")

	return true, nil
}

func downloadFileFromS3(awsSession *session.Session,localLocation string,s3BucketName string,s3Location string)error{
	downloader := s3manager.NewDownloader(awsSession)

	// create a new temp local file
	fmt.Println("download file [" + s3Location + "] from s3")
	fileSplit := (strings.Split(s3Location,"/"))
	fileName := fileSplit[len(fileSplit)-1]
	println(fileName)

	localDirectory := localLocation[:len(localLocation)-len(fileName)]
	_,err := os.Stat(localDirectory)
	if os.IsNotExist(err){
		err = os.Mkdir(localDirectory,os.ModePerm)
		if err != nil{
			fmt.Println(err)
		}
	}

	f,err := os.Create(localLocation)
	if err != nil{
		return err
	}

	fmt.Println("get the object begin!",fileName)
	_,err = downloader.Download(f,&s3.GetObjectInput{
		Bucket : aws.String(s3BucketName),
		Key : aws.String(s3Location),
	})
	if err != nil{
		return err
	}
	fmt.Println("download file from s3 successfully!")
	return nil
}

func deleteFileFromS3(awsSession *session.Session,s3BucketName string,s3Location string)error{
	s3Session := s3.New(awsSession)

	object := &s3.DeleteObjectInput{
		Bucket: aws.String(s3BucketName),
		Key: aws.String(s3Location),
	}

	_,err := s3Session.DeleteObject(object)
	if err != nil{
		return err
	}
	return nil
}

func uploadFileToS3(awsSession *session.Session,localLocation string,s3BucketName string,s3Location string)error{
	uploader := s3manager.NewUploader(awsSession)

	f, err := os.Open(localLocation)
	if err != nil{
		return err
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3BucketName),
		Key: aws.String(s3Location),
		Body: f,
	})
	if err != nil{
		return err
	}
	return nil
}

func copyFileToS3(awsSession *session.Session,s3BucketName string,s3Location string,s3NewLocation string)error{
	s3Session := s3.New(awsSession)

	sourceFile := filepath.Clean(fmt.Sprintf("/%s/%s",s3BucketName,s3Location))
	fmt.Println(sourceFile)
	objectCopy := &s3.CopyObjectInput{
		Bucket: aws.String(s3BucketName),
		CopySource: aws.String(url.PathEscape(sourceFile)),
		Key: aws.String(s3NewLocation),
	}

	_,err := s3Session.CopyObject(objectCopy)
	if err != nil{
		return err
	}
	return nil
}

