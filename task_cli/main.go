package main

import (
	"path/filepath"

	"github.com/huytn/learn-go/task/db"
	"github.com/huytn/learn-go/task_cli/cmd"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	ir err := nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}