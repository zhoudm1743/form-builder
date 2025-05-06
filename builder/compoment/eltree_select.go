package compoment

import (
	"form-builder/builder/types"
)

type ElTreeSelectBuilder struct {
	types.FormItem
}

func ElTreeSelect(field string) *ElTreeSelectBuilder {
	return &ElTreeSelectBuilder{
		types.FormItem{
			Type:  "elTreeSelect",
			Field: field,
			Props: map[string]interface{}{
				"showCheckbox": true,
				"nodeKey":      "value",
			},
		},
	}
}

// SetTitle
func (b *ElTreeSelectBuilder) SetTitle(title string) *ElTreeSelectBuilder {
	b.Title = title
	return b
}

func (b *ElTreeSelectBuilder) SetData(data []map[string]interface{}) *ElTreeSelectBuilder {
	b.Props["data"] = data
	return b
}

func (b *ElTreeSelectBuilder) SetNodeKey(key string) *ElTreeSelectBuilder {
	b.Props["nodeKey"] = key
	return b
}

func (b *ElTreeSelectBuilder) SetProp(prop map[string]interface{}) *ElTreeSelectBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *ElTreeSelectBuilder) SetComputed(config map[string]types.ComputedConfig) *ElTreeSelectBuilder {
	b.Computed = config
	return b
}

func (b *ElTreeSelectBuilder) SetHidden(t ...interface{}) *ElTreeSelectBuilder {
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

func (b *ElTreeSelectBuilder) Build() types.FormItem {
	return b.FormItem
}
