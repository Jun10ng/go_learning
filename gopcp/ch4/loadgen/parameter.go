package loadgen

import (
	"time"
)

// ParamSet 代表了载荷发生器参数的集合
type ParamSet struct {
	// 响应超时时间 单位：ns
	timeoutNS time.Duration
	// 每秒载荷量 Loads per second
	lqs uint32
	// 负载持续时间 单位：ns
	durationNS time.Duration
	// 调用结果通道
	resultCh chan *lib.CallResult
}

// CallResult 用于表示调用结果的结构
type CallResult struct {
	ID int64
	// 原生请求
	Req RawReq
	// 原生响应
	Resp RawResp
	// 原生响应代码
	Code RetCode
	// 结果成因的简述
	Msg string
	// 耗时
	Elapse time.Duration
}

// RawReq 表示原生请求的结构
type RawReq struct {
	ID  int64
	Req []byte
}

// RawResp 用来表示原生响应的结构
type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}
