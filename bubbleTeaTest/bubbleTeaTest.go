package bubbleTeaTest

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

// Model 存储应用状态的模型
type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

// initialModel 应用程序的初始状态(初始模型)
func initialModel() model {
	return model{
		// To-do list 是一个便利店清单
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		selected: make(map[int]struct{}),
	}
}

// Init 初始化方法 返回一个可以执行一些初始I/O的命令 此处不需要做任何IO，直接返回nil
func (m model) Init() tea.Cmd {
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

// View UI处理 查看当前model并返回 string（UI内容）
func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// 选项遍历
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func TeaTest() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
