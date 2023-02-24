package bilibiliapitompd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/DaydreamCafe/Bilibili-API-To-MPD/enums"
	"github.com/DaydreamCafe/Bilibili-API-To-MPD/types"
	"github.com/zencoder/go-dash/v3/mpd"
)

const (
	NOT_CONSTRUCTED = iota
	REQUEST_WITH_AVID
	REQUEST_WITH_BVID
)

type BilibiliVideo struct {
	Avid        int64
	Bvid        string
	Cid         int64
	Sessdata    string
	RequestType int
}

func NewVideoWithAvid(avid int64, cid int64, sessdata string) *BilibiliVideo {
	return &BilibiliVideo{
		Avid:        avid,
		Cid:         cid,
		Sessdata:    sessdata,
		RequestType: REQUEST_WITH_AVID,
	}
}

func NewVideoWithBvid(bvid string, cid int64, sessdata string) *BilibiliVideo {
	return &BilibiliVideo{
		Bvid:        bvid,
		Cid:         cid,
		Sessdata:    sessdata,
		RequestType: REQUEST_WITH_BVID,
	}
}

func (v *BilibiliVideo) ConvertToMPD() (string, error) {
	// 构造请求链接
	var reqURL string
	switch v.RequestType {
	case REQUEST_WITH_AVID:
		reqURL = fmt.Sprintf(
			"https://api.bilibili.com/x/player/playurl?avid=%d&cid=%d&qn=%d&fnval=%d",
			v.Avid,
			v.Cid,
			enums.QN_8K,
			enums.FNVAL_DASH|enums.FNVAL_4K|enums.FNVAL_8K|enums.FNVAL_DOLBY_AUDIO|enums.FNVAL_DOLBY_VISION|enums.FNVAL_HDR,
		)
	case REQUEST_WITH_BVID:
		reqURL = fmt.Sprintf(
			"https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&qn=%d&fnval=%d",
			v.Bvid,
			v.Cid,
			enums.QN_8K,
			enums.FNVAL_DASH|enums.FNVAL_4K|enums.FNVAL_8K|enums.FNVAL_DOLBY_AUDIO|enums.FNVAL_DOLBY_VISION|enums.FNVAL_HDR,
		)
		fmt.Println(reqURL)
	default:
		return "", errors.New("not constructed struct BilibiliVideo")
	}

	// 调用API请求结果
	request, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("Cookie", fmt.Sprintf("SESSDATA=%s", v.Sessdata))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析为视频流结构体
	var videoStream types.VideoStreamResponse
	err = json.NewDecoder(resp.Body).Decode(&videoStream)
	if err != nil {
		return "", err
	}

	// 构造MPD结构体
	m := mpd.NewMPD(
		mpd.DASH_PROFILE_ONDEMAND,
		fmt.Sprintf("PT%dS", videoStream.Data.Dash.Duration),
		fmt.Sprintf("PT%fS", videoStream.Data.Dash.MinBufferTime),
	)

	videoAS, err := m.AddNewAdaptationSetVideo(
		mpd.DASH_MIME_TYPE_VIDEO_MP4,
		"progressive",
		true,
		1,
	)
	if err != nil {
		return "", err
	}

	for _, stream := range videoStream.Data.Dash.Video {
		videoRep, _ := videoAS.AddNewRepresentationVideo(
			stream.Bandwidth,
			stream.Codecs,
			fmt.Sprint(stream.ID),
			stream.FrameRate,
			stream.Width,
			stream.Height,
		)
		videoRep.SetNewBaseURL(stream.BaseURL)
		videoRep.AddNewSegmentBase(stream.SegmentBase.IndexRange, stream.SegmentBase.Initialization)
	}

	audioAS, err := m.AddNewAdaptationSetAudio(
		mpd.DASH_MIME_TYPE_AUDIO_MP4,
		true,
		1,
		"und",
	)
	if err != nil {
		return "", err
	}

	for _, stream := range videoStream.Data.Dash.Audio {
		audioRep, _ := audioAS.AddNewRepresentationAudio(
			44100,
			stream.Bandwidth,
			stream.Codecs,
			fmt.Sprint(stream.ID),
		)
		audioRep.SetNewBaseURL(stream.BaseURL)
		audioRep.AddNewSegmentBase(stream.SegmentBase.IndexRange, stream.SegmentBase.Initialization)
	}

	// 将MPD结构体转为字符串
	mpdStr, err := m.WriteToString()
	if err != nil {
		return "", err
	}

	return mpdStr, nil
}

// 写入MPD文件
func (v *BilibiliVideo) ConvertToMPDFile(path string) error {
	mpdStr, err := v.ConvertToMPD()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(mpdStr)
	if err != nil {
		return err
	}
	
	return nil
}