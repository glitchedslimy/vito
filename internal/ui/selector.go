package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// Mensajes para el sistema de eventos
type tickMsg time.Time
type repoInjectedMsg error

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	progress progress.Model
	choice   string
	quitting bool
	loading  bool
	percent  float64
}

func (m model) Init() tea.Cmd {
	return nil
}

// tickCmd crea un intervalo para actualizar la barra
func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// downloadSimulatedCmd simula la descarga (Aquí pondrás tu lógica de git clone)
func downloadSimulatedCmd() tea.Cmd {
	return func() tea.Msg {
		// Simula tiempo de red
		time.Sleep(time.Second * 3)
		return repoInjectedMsg(nil)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			m.quitting = true
			return m, tea.Quit
		}
		if msg.String() == "enter" && !m.loading {
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.title
				m.loading = true
				// Iniciamos la descarga y el tick de la barra
				return m, tea.Batch(downloadSimulatedCmd(), tickCmd())
			}
		}

	case tea.WindowSizeMsg:
		h, v := DocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.progress.Width = msg.Width - (h * 4)

	case tickMsg:
		if m.percent < 0.95 { // Llegamos al 95% y esperamos al proceso real
			m.percent += 0.02
			return m, tickCmd()
		}
		return m, nil

	case repoInjectedMsg:
		m.percent = 1.0
		return m, tea.Quit // Cerramos cuando termina
	}

	var cmd tea.Cmd
	if !m.loading {
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return "\n Operation Canceled.\n"
	}

	if m.loading {
		title := CommandStyle.Render(fmt.Sprintf("Injecting %s...", m.choice))
		return "\n\n" +
			DocStyle.Render(fmt.Sprintf("%s\n\n%s", title, m.progress.ViewAs(m.percent))) +
			"\n\n" + DescStyle.Render(" VITO is preparing the code dosis...")
	}

	return DocStyle.Render(m.list.View())
}

func StartTemplateSelector() {
	items := []list.Item{
		item{title: "Go Backend", desc: "Hexagonal Architecture + Gin"},
		item{title: "Frontend React", desc: "Next.js + Tailwind + TS"},
		item{title: "CLI Pill", desc: "Create new CLI pills."},
	}

	// Configuración de la barra de progreso
	pg := progress.New(
		progress.WithDefaultGradient(),
	)
	pg.FullColor = string(BrandColor)

	m := model{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		progress: pg,
	}
	m.list.Title = "VITO: Select a template"

	// Usamos WithAltScreen para que la UI sea inmersiva
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}

	// Al terminar, imprimimos el resultado final fuera del AltScreen
	if m.choice != "" && !m.quitting {
		fmt.Printf("\n✨ %s inyectado con éxito. ¡A picar código!\n", CommandStyle.Render(m.choice))
	}
}
