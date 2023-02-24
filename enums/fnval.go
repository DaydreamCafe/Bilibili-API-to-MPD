package enums

const (
	// FLV格式 FLV格式已下线 仅H.264编码 部分老视频存在分段现象 与MP4、DASH格式互斥
	// FNVAL_FLV          = 0

	// MP4格式 仅H.264编码 与FLV、DASH格式互斥
	FNVAL_MP4          = 1

	// DASH格式 与MP4、FLV格式互斥
	FNVAL_DASH         = 16

	// 是否需求HDR视频 需求DASH格式 仅H.265编码 需要qn=125 大会员认证
	FNVAL_HDR          = 64

	// 是否需求4K分辨率 该值与fourk字段协同作用 需要qn=120 大会员认证
	FNVAL_4K           = 128

	// 是否需求杜比音频 需求DASH格式 大会员认证
	FNVAL_DOLBY_AUDIO  = 256

	// 是否需求杜比视界 需求DASH格式 大会员认证
	FNVAL_DOLBY_VISION = 512

	// 是否需求8K分辨率 需求DASH格式 需要qn=127 大会员认证
	FNVAL_8K           = 1024

	// 是否需求AV1编码 需求DASH格式
	FNVAL_AV1          = 2048
)
