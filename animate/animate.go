package animate

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Msg struct {
	Delay time.Duration
	jobs  []tea.Msg
}

func New(delay time.Duration) Msg {
	return Msg{
		Delay: delay,
	}
}

func (m *Msg) Push(jobs ...tea.Msg) {
	for _, job := range jobs {
		if job != nil {
			m.jobs = append(m.jobs, job)
		}
	}
}

func (m *Msg) Pop() tea.Msg {
	if len(m.jobs) == 0 {
		return nil
	}

	job := m.jobs[len(m.jobs)-1]
	m.jobs = m.jobs[:len(m.jobs)-1]
	return job
}

func (m *Msg) Shift() tea.Msg {
	if len(m.jobs) == 0 {
		return nil
	}

	job := m.jobs[0]
	m.jobs = m.jobs[1:]
	return job
}

func (m Msg) Cmd() tea.Cmd {
	if len(m.jobs) == 0 {
		return nil
	}
	return tea.Tick(m.Delay, func(time.Time) tea.Msg {
		return m
	})
}
