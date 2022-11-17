package boot

import (
	// 初始化配置文件
	_ "DJ-Blog/pkg/config"
	// 初始化日志
	_ "DJ-Blog/pkg/logger"
	// 初始化配置数据
	_ "DJ-Blog/internal/conf"
	// 初始化验证包
	_ "DJ-Blog/pkg/validate"
)

// Initialize 用来触发其它的 init函数并完成部分函数的初始化
func Initialize() {
}
