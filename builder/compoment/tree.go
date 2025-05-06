package compoment

import (
	"form-builder/builder/types"
)

type TreeBuilder struct {
	types.FormItem
}

func Tree(field string, defaultValue interface{}) *TreeBuilder {
	return &TreeBuilder{
		types.FormItem{
			Type:  "tree",
			Field: field,
			Props: map[string]interface{}{
				"data":         []map[string]interface{}{},
				"showCheckbox": true,
			},
		},
	}
}

// SetTitle
func (b *TreeBuilder) SetTitle(title string) *TreeBuilder {
	b.Title = title
	return b
}

func (b *TreeBuilder) SetData(data []map[string]interface{}) *TreeBuilder {
	b.Props["data"] = data
	return b
}

func (b *TreeBuilder) SetProp(prop map[string]interface{}) *TreeBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *TreeBuilder) SetComputed(config map[string]types.ComputedConfig) *TreeBuilder {
	b.Computed = config
	return b
}

func (b *TreeBuilder) SetHidden(t ...interface{}) *TreeBuilder {
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

func (b *TreeBuilder) Build() types.FormItem {
	return b.FormItem
}
