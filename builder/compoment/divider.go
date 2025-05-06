package compoment

import (
	"form-builder/builder/types"
)

type DividerBuilder struct {
	types.FormItem
}

func Divider() *DividerBuilder {
	return &DividerBuilder{
		types.FormItem{
			Type: "elDivider",
		},
	}
}

// SetTitle
func (b *DividerBuilder) SetTitle(title string) *DividerBuilder {
	b.Title = title
	return b
}

func (b *DividerBuilder) AddChild(content string) *DividerBuilder {
	b.Props = map[string]interface{}{
		"content": content,
	}
	return b
}

func (b *DividerBuilder) Build() types.FormItem {
	return b.FormItem
}
