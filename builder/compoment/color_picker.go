package compoment

import (
	"form-builder/builder/types"
)

type ColorPickerBuilder struct {
	types.FormItem
}

func ColorPicker(field string) *ColorPickerBuilder {
	return &ColorPickerBuilder{
		types.FormItem{
			Type:  "colorPicker",
			Field: field,
			Props: make(map[string]interface{}),
		},
	}
}

func (b *ColorPickerBuilder) SetDefault(color string) *ColorPickerBuilder {
	b.Props["defaultValue"] = color
	return b
}

func (b *ColorPickerBuilder) SetProp(prop map[string]interface{}) *ColorPickerBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *ColorPickerBuilder) SetComputed(config map[string]types.ComputedConfig) *ColorPickerBuilder {
	b.Computed = config
	return b
}

func (b *ColorPickerBuilder) SetHidden(t ...interface{}) *ColorPickerBuilder {
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

// SetTitle
func (b *ColorPickerBuilder) SetTitle(title string) *ColorPickerBuilder {
	b.Title = title
	return b
}

func (b *ColorPickerBuilder) Build() types.FormItem {
	return b.FormItem
}
