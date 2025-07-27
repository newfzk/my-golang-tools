package k8s

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

var iniChoices = []string{"命名空间", "服务", "中间件"}

// commandFunc 命令执行函数类型
type commandFunc func()

// command 命令结构体 关联显示文本和执行函数
type command struct {
	name string
	exec commandFunc
}

// model
type model struct {
	// commands 命令选项列表
	commands []command
	// cursor 当前光标位置 用于选择要执行的命令
	cursor int
}

// initialModel 初始化函数 返回初始的模型
func initialModel(commands []command) model {
	return model{
		commands: commands,
	}
}

// Init 初始化方法 如果有初始就需要执行的命令 写在此
func (m model) Init() tea.Cmd {
	// 不需要立即执行命令
	return nil
}

// Update 更新方法 当事情发生时 调用 Update 方法
// 查看发生了什么并返回更新的Model作为响应
// 还可以返回一个Cmd让更多事情发生
// 本例中 当用户按向下箭头时 Update 注意到按了向下的箭头 并相应的移动光标
// "事情发生" 以消息 Msg 的形式出现 可以是任何类型
// 消息Msg 是发生某些IO的结果 例如 按键、计时器或来自服务器的响应
//
//	通常使用类型开关来确定收到的Msg类型 也可以使用类型断言
//	本例中以`tea.KeyMsg`为例 接收按键消息 供Update使用
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.commands)-1 {
				m.cursor++
			}
		case "enter", " ":
			// 执行命令
			m.commands[m.cursor].exec()
		}
	}
	return m, nil
}

// View UI处理 查看当前model并返回 string（UI内容）
func (m model) View() string {
	// The header
	s := "要做哪类操作?\n\n"

	// 选项遍历
	for i, choice := range m.commands {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\n按 q 键退出.\n"

	// Send the UI for rendering
	return s
}

type statusMsg string

func testCmd() tea.Msg {
	return statusMsg("test Msg")
}

func TeaRun() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("[TeaRun] 命令执行报错: %v", err)
	}
}
