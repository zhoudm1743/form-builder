package types

type FormItem struct {
	Name     string                    `json:"name,omitempty"`
	Type     string                    `json:"type,omitempty"`
	Field    string                    `json:"field,omitempty"`
	Title    string                    `json:"title,omitempty"`
	Props    map[string]interface{}    `json:"props,omitempty"`
	Options  []Option                  `json:"options,omitempty"`
	Children []FormItem                `json:"children,omitempty"`
	Span     int                       `json:"span,omitempty"`
	Computed map[string]ComputedConfig `json:"computed,omitempty"`
	Hidden   bool                      `json:"hidden,omitempty"`
}

type Option struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Disabled bool        `json:"disabled,omitempty"`
}

type ComputedConfig struct {
	Mode   string      `json:"mode"`
	Group  []Condition `json:"group"`
	Invert bool        `json:"invert"`
}

type Condition struct {
	Field     string      `json:"field"`
	Condition string      `json:"condition"`
	Value     interface{} `json:"value"`
}
