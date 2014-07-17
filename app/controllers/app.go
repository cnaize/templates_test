package controllers

import (
	"bitbucket.org/kardianos/osext"
	"bytes"
	"fmt"
	"github.com/revel/revel"
	"os"
	"path/filepath"
	"text/template"
)

type App struct {
	*revel.Controller
}

//==============================================================================
func (c App) Index() revel.Result {
	fmt.Println("Begin")

	root_path, err := osext.Executable()
	if err != nil {
		return c.RenderError(err)
	}
	template_path := filepath.Join(filepath.Dir(root_path),
		"templates_test", "app", "views", "templates", "test")
	text_path := template_path + ".txt"

	fmt.Printf("File path: %v\n", text_path)
	fmt.Printf("Is file exists: %v\n", exist(text_path))

	textTmpl, err := template.ParseFiles(text_path)
	if err != nil {
		return c.RenderError(err)
	}

	var text bytes.Buffer
	args := make(map[string]interface{})
	args["message"] = "Hello there!"

	if err := textTmpl.ExecuteTemplate(&text, "test"+".txt", args); err != nil {
		return c.RenderError(err)
	}

	fmt.Println("End")
	fmt.Println("Works!")

	return c.Render()
}

//==============================================================================
func exist(file_path string) bool {
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		return false
	}

	return true
}

//==============================================================================
