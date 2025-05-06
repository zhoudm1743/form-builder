package main

import (
	"encoding/json"
	"form-builder/builder"
	"form-builder/builder/compoment"
	"form-builder/builder/types"
	"os"
)

func main() {
	// 初始化表单构建器
	fb := builder.NewFormBuilder()

	// 创建两列布局
	row := compoment.FcRow().
		AddChild(
			compoment.Col().Span(12).AddChild(
				compoment.Input("username", nil).SetTitle("用户账号"),
				compoment.Select("gender", 1).SetTitle("性别").SetOptions([]types.SelectOption{
					{Label: "男", Value: 1},
					{Label: "女", Value: 2},
				}),
			),
		).
		AddChild(
			compoment.Col().SetSpan(12).AddChild(
				compoment.ElTreeSelect("department").
					SetTitle("所属部门").
					SetData([]map[string]interface{}{
						{
							"label": "技术部",
							"value": 1,
							"children": []map[string]interface{}{
								{"label": "前端组", "value": 11},
								{"label": "后端组", "value": 12},
							},
						},
					}).SetComputed(map[string]interface{}{
				"hidden": map[string]interface{}{
					"mode": "AND",
					"group": []map[string]interface{}{
						{"field": "gender", "condition": "==", "value": 1},
					},
					"invert": true,
				},
			}),
		)

	// 添加非表单组件
	fb.AddComponents(
		compoment.Divider().SetTitle("分隔线"),
		compoment.Image().SetSrc("/static/logo.png").SetStyle(map[string]interface{}{
			"width":  "200px",
			"height": "auto",
			"margin": "20px auto",
		}),
	)

	// 构建表单配置
	config := fb.Build()

	// 生成JSON文件
	file, _ := json.MarshalIndent(config, "", "  ")
	_ = os.WriteFile("form-item.json", file, 0644)
}
