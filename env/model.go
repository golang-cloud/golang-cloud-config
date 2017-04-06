package env

import (
	"strings"
)

//PropertySource 配置参数
type PropertySource struct {
	Name   string            `json:"name"`
	Source map[string]string `json:"source"`
}

//NewPropertySource 构建属性对象
func NewPropertySource(name string, source map[string]string) *PropertySource {
	return &PropertySource{Name: name, Source: source}
}

//ToString 返回字符串描述
func (t *PropertySource) ToString() string {
	return "PropertySource [name=" + t.Name + "]"
}

//Environment 配置
type Environment struct {
	Name            string            `json:"name"`
	Profiles        []string          `json:"profiles"`
	Label           string            `json:"label"`
	PropertySources []*PropertySource `json:"propertySources"`
	Version         string            `json:"version"`
	State           string            `json:"state"`
}

//Add 配置
func (t *Environment) Add(propertySource *PropertySource) {
	t.PropertySources = append(t.PropertySources, propertySource)
}

//AddAll 配置
func (t *Environment) AddAll(propertySources []*PropertySource) {
	t.PropertySources = append(t.PropertySources, propertySources...)
}

//ToString 返回字符串描述
func (t *Environment) ToString() string {
	return "Environment [name=" + t.Name + ", profiles=" + strings.Join(t.Profiles, ",") +
		", label=" + t.Label + //", propertySources=" + t.PropertySources +
		", version=" + t.Version +
		", state=" + t.State + "]"
}
