package enums

const (
	// 240P极速 仅MP4格式支持 仅platform=html5时有效
	QN_240P         = 6
	
	// 360P流畅
	QN_360P         = 16

	// 480P清晰
	QN_480P         = 32

	// 720P高清 WEB端默认值 B站前端需要登录才能选择 但是直接发送请求可以不登录就拿到720P的取流地址 无720P时则为720P60
	QN_720P         = 64

	// 720P60高帧率 登录认证
	QN_720P_60FPS   = 74

	// 1080P高清 TV端与APP端默认值 登录认证
	QN_1080P        = 80

	// 1080P+高码率 大会员认证
	QN_1080P_HFR    = 112

	// 1080P60高帧率 大会员认证
	QN_1080P_60FPS  = 116

	// 4K超清 同QN_4K 需要fnval&128=128且fourk=1 大会员认证
	QN_2160P        = 120

	// 4K超清 同QN_2160P 需要fnval&128=128且fourk=1 大会员认证
	QN_4K           = QN_2160P

	// HDR真彩色 仅支持DASH格式 需要fnval&64=64 大会员认证
	QN_HDR          = 125

	// 杜比视界 仅支持DASH格式 需要fnval&512=512 大会员认证
	QN_DOLBY_VISION = 126

	// 8K超高清 同QN_8K 仅支持DASH格式 需要fnval&1024=1024 大会员认证
	QN_4320P        = 127
	
	// 8K超高清 同QN_4320P 仅支持DASH格式 需要fnval&1024=1024 大会员认证
	QN_8K           = QN_4320P
)
