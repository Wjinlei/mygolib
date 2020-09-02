package psutils

import "encoding/json"

// ResStat 资源结构体
type ResStat struct {
	Total       uint64      `json:"total"`       // 总量
	Available   uint64      `json:"available"`   // 可用量
	Used        uint64      `json:"used"`        // 已用量
	UsedPercent float64     `json:"usedPercent"` // 使用百分比
	Title       string      `json:"title"`       // 标题
	Info        string      `json:"info"`        // 信息
	Data        int         `json:"data"`        // 主要数据
	Data2       interface{} `json:"data2"`       // 保留,额外数据
	Data3       interface{} `json:"data3"`       // 保留,额外数据
	Data4       interface{} `json:"data4"`       // 保留,额外数据
}

// 转换为json
func (m ResStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
