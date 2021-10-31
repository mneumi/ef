package color

import "fmt"

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textMagenta
	textCyan
	textWhite
)

func Black(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textBlack)
}

func Red(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textRed)
}

func Green(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textGreen)
}

func Yellow(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textYellow)
}

func Blue(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textBlue)
}

func Magenta(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textMagenta)
}

func Cyan(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textCyan)
}

func White(format string, args ...interface{}) string {
	return setColor(fmt.Sprintf(format, args...), 0, 0, textWhite)
}

func setColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
