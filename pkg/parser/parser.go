package parser

import (
	"github.com/rafaelmf3/parse/internal"
	file "github.com/rafaelmf3/parse/pkg/file"
)

func ParseGames(fileNamePath string) {
	file := file.ReadFile(fileNamePath)
	defer file.Close()
	games := internal.ScanGames(file)
	internal.CreateJsonGameFile(games)
}
