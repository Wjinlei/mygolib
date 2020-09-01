package psutils

import "encoding/json"

// ResStat 资源结构体
type ResStat struct {
	Load1       float64 `json:"load1"`       // 平均每1分钟的负载值
	Load5       float64 `json:"load5"`       // 平均每5分钟的负载值
	Load15      float64 `json:"load15"`      // 平均每15分钟的负载值
	Total       uint64  `json:"total"`       // 总量
	Available   uint64  `json:"available"`   // 可用量
	Used        uint64  `json:"used"`        // 已用量
	UsedPercent float64 `json:"usedPercent"` // 使用百分比
	Title       string  `json:"title"`       // 标题
	Info        string  `json:"info"`        // 信息
	Data        int     `json:"data"`        // 扩展数据1, 一般用于前端需要
	Data2       int     `json:"data2"`       // 保留
	Data3       int     `json:"data3"`       // 保留
	Data4       int     `json:"data4"`       // 保留
	Data5       int     `json:"data5"`       // 保留
}

// 转换为json
func (m ResStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
