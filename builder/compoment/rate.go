package compoment

import (
	"form-builder/builder/types"
)

type RateBuilder struct {
	types.FormItem
}

func Rate(field string) *RateBuilder {
	return &RateBuilder{
		types.FormItem{
			Type:  "rate",
			Field: field,
			Props: make(map[string]interface{}),
		},
	}
}

// SetTitle
func (b *RateBuilder) SetTitle(title string) *RateBuilder {
	b.Title = title
	return b
}

func (b *RateBuilder) SetMax(max int) *RateBuilder {
	b.Props["max"] = max
	return b
}

func (b *RateBuilder) SetProp(prop map[string]interface{}) *RateBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *RateBuilder) SetComputed(config map[string]types.ComputedConfig) *RateBuilder {
	b.Computed = config
	return b
}

func (b *RateBuilder) SetHidden(t ...interface{}) *RateBuilder {
	if len(t) == 0 {
		return b
	}

	switch v := t[0].(type) {
	case bool:
		b.Hidden = v // 直接处理布尔值
	case types.ComputedConfig:
		b.Computed = map[string]types.ComputedConfig{
			"hidden": v,
		} // 处理计算配置
	case func() bool:
		b.Hidden = v() // 支持动态计算
	default:
		panic("unsupported parameter type") // 或返回错误
	}
	return b
}

func (b *RateBuilder) Build() types.FormItem {
	return b.FormItem
}
