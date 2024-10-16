package control

import (
	"fmt"

	zero "github.com/wdvxdr1123/ZeroBot"
)

// Matcher 是 ZeroBot 匹配和处理事件的最小单元
type Matcher zero.Matcher

// SetBlock 设置是否阻断后面的 Matcher 触发
func (m *Matcher) SetBlock(block bool) *Matcher {
	_ = (*zero.Matcher)(m).SetBlock(block)
	return m
}

// Handle 直接处理事件
func (m *Matcher) Handle(handler zero.Handler) {
	_ = (*zero.Matcher)(m).Handle(handler)
}

// AutoSetName 设置 Matcher 名称
func (m *Matcher) AutoSetName(e *Engine) *Matcher {
	name := fmt.Sprintf("%s.%d", e.service, e.en.Count())
	m.Name = name
	return m
}

// SetName 设置 Matcher 名称
func (m *Matcher) SetName(name string) *Matcher {
	m.Name = name
	return m
}
