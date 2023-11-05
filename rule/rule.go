package rule

import (
	"fmt"
	"strings"
)

type Rule struct {
	Condition string // 规则条件
	Action    string // 规则触发的操作
}

// 规则引擎，检测流量是否匹配规则
func ruleEngine(rule Rule, networkData string) bool {
	// 在这里，你可以实现规则的匹配逻辑
	// 这是一个简单示例，只检查字符串是否包含规则条件
	return strings.Contains(networkData, rule.Condition)
}

func rule() {
	// 示例规则
	rule1 := Rule{
		Condition: "malware",
		Action:    "Block the traffic",
	}

	rule2 := Rule{
		Condition: "high bandwidth usage",
		Action:    "Alert the administrator",
	}

	// 模拟网络流量数据
	networkData := "This traffic contains malware and has high bandwidth usage."

	// 检查规则是否匹配并执行操作
	if ruleEngine(rule1, networkData) {
		fmt.Println("Rule 1 匹配. Action: 封锁交通.")
	}

	if ruleEngine(rule2, networkData) {
		fmt.Println("Rule 2 matched. Action: 提醒管理员.")
	}
}
