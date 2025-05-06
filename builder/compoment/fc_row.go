package compoment

import (
	"form-builder/builder/types"
)

type FcRowBuilder struct {
	types.FormItem
}

func FcRow() *FcRowBuilder {
	return &FcRowBuilder{
		types.FormItem{
			Type: "fcRow",
		},
	}
}

func (b *FcRowBuilder) AddChild(children ...types.FormItem) *FcRowBuilder {
	b.Children = append(b.Children, children...)
	return b
}

func (b *FcRowBuilder) Build() types.FormItem {
	return b.FormItem
}

type ColBuilder struct {
	types.FormItem
}

func Col() *ColBuilder {
	return &ColBuilder{
		types.FormItem{
			Type: "col",
		},
	}
}

func (b *ColBuilder) SetSpan(span int) *ColBuilder {
	b.Span = span
	return b
}

func (b *ColBuilder) SetComputed(config map[string]types.ComputedConfig) *ColBuilder {
	b.Computed = config
	return b
}

func (b *ColBuilder) SetHidden(t ...interface{}) *ColBuilder {
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

func (b *ColBuilder) AddChild(items ...types.FormItem) *ColBuilder {
	b.Children = append(b.Children, items...)
	return b
}

func (b *ColBuilder) Build() types.FormItem {
	return b.FormItem
}
