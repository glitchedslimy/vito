package templates

import (
	"fmt"
	"os"
	"os/exec"
)

type RemoteTemplate struct {
	URL  string
	Name string
}

func (t RemoteTemplate) Inject(targetDir string) error {
	// 1. Clonar el repositorio (depth 1 para que sea instantáneo)
	fmt.Printf("💊 Inyectando vitamina %s desde %s...\n", t.Name, t.URL)

	cmd := exec.Command("git", "clone", "--depth", "1", t.URL, targetDir)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al descargar el template: %v", err)
	}

	// 2. Limpiar rastros de Git para que sea un boilerplate puro
	gitDir := fmt.Sprintf("%s/.git", targetDir)
	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("error al limpiar metadatos: %v", err)
	}

	return nil
}
