/*
非常感谢bilibili-API-collect项目提供的API和相关注解
项目地址: https://github.com/SocialSisterYi/bilibili-API-collect
*/
package types

// Dash视频流结构体
type VideoDashVideoStream struct {
	ID           int      `json:"id"`
	BaseURL      string   `json:"baseUrl"`
	BackupURL    []string `json:"backupUrl"`
	Bandwidth    int64    `json:"bandwidth"`
	MimeType     string   `json:"mimeType"`
	Codecs       string   `json:"codecs"`
	Width        int64    `json:"width"`
	Height       int64    `json:"height"`
	FrameRate    string   `json:"frameRate"`
	Sar          string   `json:"sar"`
	StartWithSap int      `json:"startWithSap"`
	SegmentBase  struct {
		Initialization string `json:"Initialization"`
		IndexRange     string `json:"indexRange"`
	} `json:"SegmentBase"`
	Codecid int `json:"codecid"`
}

// Dash音频流结构体
type VideoDashAudioStream struct {
	ID           int      `json:"id"`
	BaseURL      string   `json:"baseUrl"`
	BackupURL    []string `json:"backupUrl"`
	Bandwidth    int64    `json:"bandwidth"`
	MimeType     string   `json:"mimeType"`
	Codecs       string   `json:"codecs"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	FrameRate    string   `json:"frameRate"`
	Sar          string   `json:"sar"`
	StartWithSap int      `json:"startWithSap"`
	SegmentBase  struct {
		Initialization string `json:"Initialization"`
		IndexRange     string `json:"indexRange"`
	} `json:"SegmentBase"`
	Codecid int `json:"codecid"`
}

// Dash杜比音频流结构体
type VideoDashDolbyAudioStream struct {
	ID           int      `json:"id"`
	BaseURL      string   `json:"baseUrl"`
	BackupURL    []string `json:"backupUrl"`
	Bandwidth    int      `json:"bandwidth"`
	MimeType     string   `json:"mimeType"`
	Codecs       string   `json:"codecs"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	FrameRate    string   `json:"frameRate"`
	Sar          string   `json:"sar"`
	StartWithSap int      `json:"startWithSap"`
	SegmentBase  struct {
		Initialization string `json:"Initialization"`
		IndexRange     string `json:"indexRange"`
	} `json:"SegmentBase"`
	Codecid int `json:"codecid"`
}

// Dash杜比流结构体
type VideoDashDolby struct {
	Type  int                         `json:"type"`
	Audio []VideoDashDolbyAudioStream `json:"audio"`
}

// Hi-Res无损音频音频流
type VideoDashFlacAudioStream struct {
	ID           int      `json:"id"`
	BaseURL      string   `json:"baseUrl"`
	BackupURL    []string `json:"backupUrl"`
	Bandwidth    int      `json:"bandwidth"`
	MimeType     string   `json:"mimeType"`
	Codecs       string   `json:"codecs"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	FrameRate    string   `json:"frameRate"`
	Sar          string   `json:"sar"`
	StartWithSap int      `json:"startWithSap"`
	SegmentBase  struct {
		Initialization string `json:"Initialization"`
		IndexRange     string `json:"indexRange"`
	} `json:"SegmentBase"`
	Codecid int `json:"codecid"`
}

// Hi-Res无损音频流结构体
type VideoDashFlac struct {
	Display bool                     `json:"display"`
	Audio   VideoDashFlacAudioStream `json:"audio"`
}

// Dash结构体
type VideoDash struct {
	Duration      int                    `json:"duration"`
	MinBufferTime float64                `json:"minBufferTime"`
	Video         []VideoDashVideoStream `json:"video"`
	Audio         []VideoDashAudioStream `json:"audio"`
	Dolby         VideoDashDolby         `json:"dolby"`
	Flac          VideoDashFlac          `json:"flac"`
}

// Data结构体
type VideoData struct {
	// local？
	From string `json:"from"`
	// suee？
	Result string `json:"result"`
	// 空？
	Message string `json:"message"`
	// 清晰度标识 含义见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/videostream_url.md#qn%E8%A7%86%E9%A2%91%E6%B8%85%E6%99%B0%E5%BA%A6%E6%A0%87%E8%AF%86
	Quality int `json:"quality"`
	// 视频格式 mp4/flv
	Format string `json:"format"`
	// 视频长度 单位为毫秒 不同分辨率 / 格式可能有略微差异
	Timelength int `json:"timelength"`
	// 支持的全部格式 每项用,分隔
	AcceptFormat string `json:"accept_format"`
	// 支持的清晰度列表（文字说明）
	AcceptDescription []string `json:"accept_description"`
	// 支持的清晰度列表（代码） 含义见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/videostream_url.md#qn%E8%A7%86%E9%A2%91%E6%B8%85%E6%99%B0%E5%BA%A6%E6%A0%87%E8%AF%86
	AcceptQuality []int `json:"accept_quality"`
	// 默认选择视频流的编码id 含义见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/videostream_url.md#%E8%A7%86%E9%A2%91%E7%BC%96%E7%A0%81%E4%BB%A3%E7%A0%81
	VideoCodecid int `json:"video_codecid"`
	// start？
	SeekParam string `json:"seek_param"`
	// offset（DASH / FLV） second（MP4）？
	SeekType string `json:"seek_type"`
	// DASH 流信息
	Dash VideoDash `json:"dash"`
	// 支持格式的详细信息
	SupportFormats []struct {
		Quality        int      `json:"quality"`
		Format         string   `json:"format"`
		NewDescription string   `json:"new_description"`
		DisplayDesc    string   `json:"display_desc"`
		Superscript    string   `json:"superscript"`
		Codecs         []string `json:"codecs"`
	} `json:"support_formats"`
	// （？）
	HighFormat interface{} `json:"high_format"`
	// 上次播放进度
	LastPlayTime int `json:"last_play_time"`
	// 上次播放分P的cid
	LastPlayCid int `json:"last_play_cid"`
}

// B站API返回结果结构体
type VideoStreamResponse struct {
	// 返回值  0：成功  -400：请求错误  -404：无视频
	Code int `json:"code"`
	// 错误信息 默认为0
	Message string `json:"message"`
	// 恒为1
	TTL int `json:"ttl"`
	// 数据本体
	Data VideoData `json:"data"`
}
