package sample

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func GetActivityData()(*activity.Metadata){
	if activityMetadata == nil{
		jsonfile := "descriptor.json"
		jsonMetadataBytes,err := ioutil.ReadFile(jsonfile)
		if err != nil{
			panic("No json file found in " + jsonfile)
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestCreate(t *testing.T){
	fmt.Println("-----------TestCreate begin-------------")
	act := NewActivity(GetActivityData())
	if act == nil{
		t.Error("Create activity fail!")
		t.Fail()
		return
	}
}

func TestEvalDownload(t *testing.T){
	fmt.Println("-----------TestDownload begin-------------")
	act := NewActivity(GetActivityData())
	tc := test.NewTestActivityContext(GetActivityData())

	tc.SetInput("action","download")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	tc.SetInput("s3BucketName", "")
	tc.SetInput("s3Location", "/test/lp.png")
	tc.SetInput("localLocation", "/home/ubuntu/project/goland/src/flogo_activity_amazons3/img/lp.jpg")
	act.Eval(tc)

	result := tc.GetOutput("result")
	fmt.Println(result)
	fmt.Println("-----------TestDownload end-------------")
}

func TestEvalUpload(t *testing.T){
	fmt.Println("-----------TestUpload begin-------------")
	act := NewActivity(GetActivityData())
	tc := test.NewTestActivityContext(GetActivityData())

	tc.SetInput("action","upload")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	tc.SetInput("s3BucketName", "")
	tc.SetInput("s3Location", "/test/lp.png")
	tc.SetInput("localLocation", "/home/ubuntu/project/goland/src/flogo_activity_amazons3/img/zs-0001_result.jpg")
	act.Eval(tc)

	result := tc.GetOutput("result")
	fmt.Println(result)
	fmt.Println("-----------TestUpload end-------------")
}

func TestEvalCopy(t *testing.T){
	fmt.Println("-----------TestCopy begin-------------")
	act := NewActivity(GetActivityData())
	tc := test.NewTestActivityContext(GetActivityData())

	tc.SetInput("action","copy")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	tc.SetInput("s3BucketName", "")
	tc.SetInput("s3Location", "/test/lp.png")
	tc.SetInput("s3NewLocation", "/test/lp_1.png")
	act.Eval(tc)

	result := tc.GetOutput("result")
	fmt.Println(result)
	fmt.Println("-----------TestUpload end-------------")
}

func TestEvalDelete(t *testing.T){
	fmt.Println("-----------TestDelete begin-------------")
	act := NewActivity(GetActivityData())
	tc := test.NewTestActivityContext(GetActivityData())

	tc.SetInput("action","delete")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	tc.SetInput("s3BucketName", "")
	tc.SetInput("s3Location", "/test/lp_1.png")
	act.Eval(tc)

	result := tc.GetOutput("result")
	fmt.Println(result)
	fmt.Println("-----------TestDelete end-------------")
}

