package method

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bestruirui/mihomo-check/utils"
)

func SaveToLocal(yamlData []byte, key string) error {
	Path := utils.GetExecutablePath()

	fileName := fmt.Sprintf("%s.yaml", key)

	if _, err := os.Stat(filepath.Join(Path, "output")); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(Path, "output"), 0755)
	}

	path := filepath.Join(Path, "output", fileName)

	os.WriteFile(path, yamlData, 0644)

	return nil
}
