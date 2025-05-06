package compoment

import (
	"form-builder/builder/types"
)

type RadioBuilder struct {
	types.FormItem
}

func Radio(field string, defaultValue interface{}) *RadioBuilder {
	return &RadioBuilder{
		types.FormItem{
			Type:  "radio",
			Field: field,
			Props: map[string]interface{}{"defaultValue": defaultValue},
		},
	}
}

// SetTitle
func (b *RadioBuilder) SetTitle(title string) *RadioBuilder {
	b.Title = title
	return b
}

func (b *RadioBuilder) SetOptions(options []types.Option) *RadioBuilder {
	b.Options = options
	return b
}

func (b *RadioBuilder) SetProp(prop map[string]interface{}) *RadioBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *RadioBuilder) SetComputed(config map[string]types.ComputedConfig) *RadioBuilder {
	b.Computed = config
	return b
}

func (b *RadioBuilder) SetHidden(t ...interface{}) *RadioBuilder {
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

func (b *RadioBuilder) Build() types.FormItem {
	return b.FormItem
}
