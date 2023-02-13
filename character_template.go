package main

type AttributeOrderOthers struct {
	SortBy string `yaml:"sortBy"` // time | Name | value desc
}

type ShowSettings struct {
	Top     []string          `yaml:"top,flow"`
	SortBy  string            `yaml:"sortBy"`  // time | Name | value desc
	Ignores []string          `yaml:"ignores"` // 这里面的属性将不被显示
	ShowAs  map[string]string `yaml:"show-as"` // 展示形式，即st show时格式
}

type TextTemplateItem = []interface{} // 实际上是 [](string | int) 类型
type TextTemplateItemList []TextTemplateItem

type TextTemplateWithWeight = map[string][]TextTemplateItem
type TextTemplateWithWeightDict = map[string]TextTemplateWithWeight

type CharacterTemplate struct {
	KeyName           string       `yaml:"key-name"`                // 模板名字
	FullName          string       `yaml:"full-name"`               // 全名
	Authors           []string     `yaml:"authors"`                 // 作者
	Version           string       `yaml:"version"`                 // 版本
	UpdatedTime       string       `yaml:"updatedTime"`             // 更新日期
	NameTemplateKeys  []string     `yaml:"name-template-keys,flow"` // .sn xxx
	NameTemplateValue string       `yaml:"name-template-value"`     // sn名片模板
	ShowSettings      ShowSettings `yaml:"show-settings"`           // 默认展示顺序

	Defaults         map[string]int64    `yaml:"defaults"`          // 默认值
	DefaultsComputed map[string]string   `yaml:"defaults-computed"` // 计算类型
	Alias            map[string][]string `yaml:"alias"`             // 别名/同义词

	Texts *TextTemplateWithWeightDict `yaml:"texts"` // UI文本

	//BasedOn           string                 `yaml:"based-on"`           // 基于规则
}

//type DefaultValue struct {
//	TypeId VMValueType `json:"typeId"`
//	Value  interface{} `json:"value"`
//}

//type CharacterTemplate struct {
//	// 属性
//	// 技能
//	// 默认值 AttrDefault
//	// 同义词 synonyms
//	AttrList      []string
//	DefaultValues map[string]DefaultValue
//}
