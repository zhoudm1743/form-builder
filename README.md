# Form Builder 表单构建器

基于Go语言和[FormCreate](https://github.com/xaboy/form-create-designer)实现的动态表单生成系统，可通过JSON配置快速构建企业级表单。

## 功能特性
- 🏗️ 可视化表单构建
- 📦 丰富的组件库（输入框、下拉选择、树形选择等）
- 📐 响应式布局系统
- ⚙️ 动态计算逻辑配置
- 📄 自动生成JSON配置文件

## 快速开始

### 安装依赖
```bash
go get form-builder/builder
```

### 基础用法
```go
package main

import (
    "form-builder/builder"
    "form-builder/builder/compoment"
    "form-builder/builder/types"
)

func main() {
    fb := builder.NewFormBuilder()
    
    // 创建两列布局
    row := compoment.FcRow()
        .AddChild(
            compoment.Col().Span(12).AddChild(
                compoment.Input("username", nil).SetTitle("用户账号"),
                compoment.Select("gender", 1)
                    .SetTitle("性别")
                    .SetOptions([]types.SelectOption{
                        {Label: "男", Value: 1},
                        {Label: "女", Value: 2},
                    }),
            ),
        )
    
    // 添加动态计算逻辑
    treeSelect := compoment.ElTreeSelect("department")
        .SetTitle("所属部门")
        .SetData([]map[string]interface{}{
            {
                "label": "技术部",
                "children": [
                    {"label": "前端组", "value": 11},
                    {"label": "后端组", "value": 12},
                ],
            },
        })
        .SetComputed(map[string]interface{}{
            "hidden": map[string]interface{}{
                "mode": "AND",
                "group": []map[string]interface{}{
                    {"field": "gender", "condition": "==", "value": 1},
                },
                "invert": true,
            },
        })

    // 生成配置
    config := fb.Build()
    // 生成form-item.json文件...
}
```

## 组件配置说明

### 通用属性
| 属性名     | 类型       | 说明                  | 示例                     |
|------------|------------|---------------------|-------------------------|
| SetTitle   | string     | 组件标题              | SetTitle("用户账号")    |
| SetOptions | []Option   | 下拉选项配置          | 见Select组件示例         |
| SetComputed| map        | 动态计算配置          | 见树形选择示例           |
| SetHidden  | bool/func  | 条件隐藏              | SetHidden(true)         |

### 布局组件
**FcRow** - 行容器：
```go
compoment.FcRow().AddChild(col1, col2)
```

**Col** - 列容器：
```go
compoment.Col().Span(12) // 占12栅格
```

## 高级功能

### 动态计算配置
通过SetComputed实现字段联动：
```go
.SetComputed(map[string]interface{}{
    "hidden": map[string]interface{}{
        "mode": "AND",
        "group": []map[string]interface{}{
            {"field": "gender", "condition": "==", "value": 1},
        },
        "invert": true,
    },
})
```

### 生成配置文件
构建完成后自动生成标准化JSON：
```go
file, _ := json.MarshalIndent(config, "", "  ")
_ = os.WriteFile("form-item.json", file, 0644)
```

## 项目结构
```
form-builder/
├── builder/            # 核心构建模块
│   ├── compoment/      # 表单组件库
│   └── types/          # 类型定义
├── main.go             # 入口文件
└── form-item.json      # 生成的配置文件
```