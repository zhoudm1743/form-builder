package builder

import (
	"form-builder/builder/types"
)

type FormBuilder struct {
	components []types.FormItem
}

func NewFormBuilder() *FormBuilder {
	return &FormBuilder{
		components: make([]types.FormItem, 0),
	}
}

func (fb *FormBuilder) AddComponents(items ...types.FormItem) {
	if len(fb.components) == 0 {
		fb.components = items
		return
	}
	if cap(fb.components)-len(fb.components) < len(items) {
		newComponents := make([]types.FormItem, len(fb.components), len(fb.components)+len(items)+10)
		copy(newComponents, fb.components)
		fb.components = newComponents
	}
	fb.components = append(fb.components, items...)
}

func (fb *FormBuilder) Build() []types.FormItem {
	return fb.components
}
