package compoment

import (
	"form-builder/builder/types"
)

type TimePickerBuilder struct {
	types.FormItem
}

func TimePicker(field string) *TimePickerBuilder {
	return &TimePickerBuilder{
		types.FormItem{
			Type:  "timePicker",
			Field: field,
			Props: make(map[string]interface{}),
		},
	}
}

// SetTitle
func (b *TimePickerBuilder) SetTitle(title string) *TimePickerBuilder {
	b.Title = title
	return b
}

func (b *TimePickerBuilder) SetRange(isRange bool) *TimePickerBuilder {
	b.Props["isRange"] = isRange
	return b
}

func (b *TimePickerBuilder) SetProp(prop map[string]interface{}) *TimePickerBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *TimePickerBuilder) SetComputed(config map[string]types.ComputedConfig) *TimePickerBuilder {
	b.Computed = config
	return b
}

func (b *TimePickerBuilder) SetHidden(t ...interface{}) *TimePickerBuilder {
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

func (b *TimePickerBuilder) Build() types.FormItem {
	return b.FormItem
}
