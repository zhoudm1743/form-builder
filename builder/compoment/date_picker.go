package compoment

import (
	"form-builder/builder/types"
)

type DatePickerBuilder struct {
	types.FormItem
}

func DatePicker(field string) *DatePickerBuilder {
	return &DatePickerBuilder{
		types.FormItem{
			Type:  "datePicker",
			Field: field,
			Props: make(map[string]interface{}),
		},
	}
}

func (b *DatePickerBuilder) SetTitle(title string) *DatePickerBuilder {
	b.Title = title
	return b
}

func (b *DatePickerBuilder) SetFormat(format string) *DatePickerBuilder {
	b.Props["format"] = format
	return b
}

func (b *DatePickerBuilder) SetProp(prop map[string]interface{}) *DatePickerBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *DatePickerBuilder) SetComputed(config map[string]types.ComputedConfig) *DatePickerBuilder {
	b.Computed = config
	return b
}

func (b *DatePickerBuilder) SetHidden(t ...interface{}) *DatePickerBuilder {
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

func (b *DatePickerBuilder) Build() types.FormItem {
	return b.FormItem
}
