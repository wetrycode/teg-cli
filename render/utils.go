package render

import (
	"encoding/json"
	"regexp"
	"strings"
)

func StructToMap(src interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	val := make(map[string]interface{})
	err = json.Unmarshal(bytes, &val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func NamedToCamelCase(name string, export bool) string {
	// 使用正则表达式匹配非字母和数字的字符作为分隔符
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	// 分割字符串
	words := reg.Split(name, -1)

	// 遍历单词，将首字母转换为大写
	for i, word := range words {
		if len(word) > 0 {
			// 首字母大写
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	if !export {
		words[0] = strings.ToLower(words[0])
	}
	// 连接单词
	return strings.Join(words, "")
}
