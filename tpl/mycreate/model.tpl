package {{ .PackageName }}


import (
    "gorm.io/gorm"
    "encoding/json"
)

// {{ .StructName }} {{ .StructNameSnakeCase }} model
type {{ .StructName }} struct {
	gorm.Model
}

// String: 返回json字符串
// @return string: json字符串
func(m {{ .StructName }}) String() string {
    if jsonBytes, err := json.Marshal(m); err != nil {
		return ""
	} else {
		return string(jsonBytes)
	}
}


// TableName 返回表名
// @return string: 表名
func (m {{ .StructName }}) TableName() string {
	return "{{ .StructNameSnakeCase  }}"
}
