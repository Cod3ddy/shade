package lib

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{
	choices [] string
	cursor int
	selected map[int]struct{}
}


func InitialModel(choices []string)model{
	return model{
		choices: choices,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd{
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "ctr+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0{
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1{
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok{
				delete(m.selected, m.cursor)
			}else{
				m.selected[m.cursor] = struct{}{}
			}
		}//end switch 
	}//end switch

	return m, nil
}

func (m model) View() string{

	// header

	s := "Select which Log line to remove use  upArrow, downArrow and ENTER"

	for i, choice := range m.choices{
		cursor := " "
		checked := " "


		if m.cursor == i{
			cursor = ">"
		}
		
		if _, ok := m.selected[i]; ok{
			checked = "x"
		}


		s += fmt.Sprintf("%s [%s] %s", cursor, checked, choice)
		
	}

	s += "\n Press q to quit.\n"
	return s
}