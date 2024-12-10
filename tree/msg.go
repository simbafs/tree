package tree

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Cmd[T any](msg T) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

type Msg string

func Msgf(msg string, a ...any) tea.Cmd {
	if len(a) > 0 {
		msg = fmt.Sprintf(msg, a...)
	}
	return func() tea.Msg {
		return Msg(msg)
	}
}

type ErrorMsg struct {
	error
}

func ErrMsg(msg error) tea.Cmd {
	return func() tea.Msg {
		return ErrorMsg{msg}
	}
}

func ErrMsgf(msg string, a ...any) tea.Cmd {
	return func() tea.Msg {
		return ErrorMsg{fmt.Errorf(msg, a...)}
	}
}
