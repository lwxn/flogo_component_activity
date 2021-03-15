package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	VideoSrcPath string `md:"videoSrcPath,required"`
	FrameDesPath string `md:"frameDesPath,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["videoSrcPath"])
	r.VideoSrcPath = strVal

	strVal, _ = coerce.ToString(values["frameDesPath"])
	r.FrameDesPath = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"videoSrcPath":r.VideoSrcPath,
		"frameDesPath": r.FrameDesPath,
	}
}

type Output struct {
	Result string `md:"result"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Result"])
	o.Result = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}
