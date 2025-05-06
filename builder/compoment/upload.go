package compoment

import (
	"form-builder/builder/types"
)

type UploadBuilder struct {
	types.FormItem
}

func Upload(field string) *UploadBuilder {
	return &UploadBuilder{
		types.FormItem{
			Type:  "upload",
			Field: field,
			Props: make(map[string]interface{}),
		},
	}
}

// SetTitle
func (b *UploadBuilder) SetTitle(title string) *UploadBuilder {
	b.Props["title"] = title
	return b
}

func (b *UploadBuilder) SetAction(action string) *UploadBuilder {
	b.Props["action"] = action
	return b
}

func (b *UploadBuilder) SetOnSuccess(callback string) *UploadBuilder {
	b.Props["onSuccess"] = callback
	return b
}

func (b *UploadBuilder) SetProp(prop map[string]interface{}) *UploadBuilder {
	for k, v := range prop {
		b.Props[k] = v
	}
	return b
}

func (b *UploadBuilder) SetComputed(config map[string]types.ComputedConfig) *UploadBuilder {
	b.Computed = config
	return b
}

func (b *UploadBuilder) SetHidden(t ...interface{}) *UploadBuilder {
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

func (b *UploadBuilder) Build() types.FormItem {
	return b.FormItem
}
