package main

type AttrSettings struct {
	Top     []string          `yaml:"top,flow" json:"top,flow"`
	SortBy  string            `yaml:"sortBy" json:"sortBy"`   // time | Name | value desc
	Ignores []string          `yaml:"ignores" json:"ignores"` // 这里面的属性将不被显示
	ShowAs  map[string]string `yaml:"showAs" json:"showAs"`   // 展示形式，即st show时格式
	Setter  map[string]string `yaml:"setter" json:"setter"`   // st写入时执行这个
}

type TextTemplateItem = []interface{} // 实际上是 [](string | int) 类型
type TextTemplateItemList []TextTemplateItem

type TextTemplateWithWeight = map[string][]TextTemplateItem
type TextTemplateWithWeightDict = map[string]TextTemplateWithWeight

// TextTemplateHelpItem 辅助信息，用于UI中，大部分自动生成
type TextTemplateHelpItem = struct {
	Filename        []string           `json:"filename"` // 文件名
	Origin          []TextTemplateItem `json:"origin"`   // 初始文本
	Vars            []string           `json:"vars"`     // 可用变量
	Commands        []string           `json:"commands"` // 所属指令
	Modified        bool               `json:"modified"` // 跟初始相比是否有修改
	SubType         string             `json:"subType"`
	ExtraText       string             `json:"extraText"`       // 额外解说
	ExampleCommands []string           `json:"exampleCommands"` // 案例命令
	NotBuiltin      bool               `json:"notBuiltin"`      // 非内置
}

type TextTemplateHelpGroup = map[string]*TextTemplateHelpItem
type TextTemplateWithHelpDict = map[string]TextTemplateHelpGroup

type NameTemplateItem struct {
	Template string `yaml:"template" json:"template"`
	HelpText string `yaml:"helpText" json:"helpText"`
}

type CharacterTemplate struct {
	KeyName      string                      `yaml:"keyName" json:"keyName"`           // 模板名字
	FullName     string                      `yaml:"fullName" json:"fullName"`         // 全名
	Authors      []string                    `yaml:"authors" json:"authors"`           // 作者
	Version      string                      `yaml:"version" json:"version"`           // 版本
	UpdatedTime  string                      `yaml:"updatedTime" json:"updatedTime"`   // 更新日期
	TemplateVer  string                      `yaml:"templateVer" json:"templateVer"`   // 模板版本
	NameTemplate map[string]NameTemplateItem `yaml:"nameTemplate" json:"nameTemplate"` // 名片模板
	AttrSettings AttrSettings                `yaml:"attrSettings" json:"attrSettings"` // 默认展示顺序

	Defaults         map[string]int64    `yaml:"defaults" json:"defaults"`                 // 默认值
	DefaultsComputed map[string]string   `yaml:"defaultsComputed" json:"defaultsComputed"` // 计算类型
	Alias            map[string][]string `yaml:"alias" json:"alias"`                       // 别名/同义词

	TextMap         *TextTemplateWithWeightDict `yaml:"textMap" json:"textMap"` // UI文本
	TextMapHelpInfo *TextTemplateWithHelpDict   `yaml:"TextMapHelpInfo" json:"textMapHelpInfo"`

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
