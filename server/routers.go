package server

import (
	"net/http"

	"strings"

	"github.com/labstack/echo"
)

//Hander ww
type Hander struct {
	repository EnvironmentRepository
	//stripDocument
}

//NewHander ww
func NewHander(url, basedir string) *Hander {
	return &Hander{repository: &GGit{url: url, basedir: basedir}}
}

//Labelled @RequestMapping("/:name/:profiles/:label")
func (ec *Hander) Labelled(c echo.Context) error {

	name := c.Param("name")
	profiles := c.Param("profiles")
	label := c.Param("label")

	if strings.Contains(label, "(_)") {
		// "(_)" is uncommon in a git branch name, but "/" cannot be matched
		label = strings.Replace(label, "(_)", "/", -1)
	}
	env := ec.repository.FindOne(name, profiles, label)

	return c.JSON(http.StatusOK, env)
}
