package compoment

import (
	"form-builder/builder/types"
)

type SelectBuilder struct {
	types.FormItem
}

func Select(field string, defaultValue interface{}) *SelectBuilder {
	return &SelectBuilder{
		types.FormItem{
			Type:  "select",
			Field: field,
			Props: map[string]interface{}{"defaultValue": defaultValue},
		},
	}
}

// SetTitle
func (b *SelectBuilder) SetTitle(title string) *SelectBuilder {
	b.Title = title
	return b
}

func (b *SelectBuilder) SetProp(prop map[string]interface{}) *SelectBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *SelectBuilder) SetOptions(options []types.Option) *SelectBuilder {
	b.Options = options
	return b
}

func (b *SelectBuilder) SetComputed(config map[string]types.ComputedConfig) *SelectBuilder {
	b.Computed = config
	return b
}

func (b *SelectBuilder) SetHidden(t ...interface{}) *SelectBuilder {
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

func (b *SelectBuilder) Build() types.FormItem {
	return b.FormItem
}
