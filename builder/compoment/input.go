package compoment

import (
	"form-builder/builder/types"
)

type InputBuilder struct {
	types.FormItem
}

func Input(field string, defaultValue interface{}) *InputBuilder {
	return &InputBuilder{
		types.FormItem{
			Type:  "input",
			Field: field,
			Props: map[string]interface{}{"defaultValue": defaultValue},
		},
	}
}

// SetTitle
func (b *InputBuilder) SetTitle(title string) *InputBuilder {
	b.Title = title
	return b
}

func (b *InputBuilder) SetName(name string) *InputBuilder {
	b.Name = name
	return b
}

func (b *InputBuilder) SetProp(prop map[string]interface{}) *InputBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *InputBuilder) SetComputed(config map[string]types.ComputedConfig) *InputBuilder {
	b.Computed = config
	return b
}

func (b *InputBuilder) SetHidden(t ...interface{}) *InputBuilder {
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

func (b *InputBuilder) Build() types.FormItem {
	return b.FormItem
}
