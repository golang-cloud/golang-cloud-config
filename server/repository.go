package server

import (
	"fmt"
	"os"
	"sync"

	"path/filepath"

	"github.com/golang-cloud/golang-cloud-config/env"
	"github.com/magiconair/properties"
	"gopkg.in/src-d/go-git.v4"
)

//EnvironmentRepository 接口
type EnvironmentRepository interface {
	FindOne(application, profile, label string) *env.Environment
}

func checkIfError(err error) bool {
	if err == nil {
		return false
	}
	fmt.Printf("[INFO]%s\n", fmt.Sprintf(" %s", err))
	return true
}

//GGit s
type GGit struct {
	sync.Mutex
	url     string
	basedir string //
}

func (g *GGit) refresh() {

	g.Lock()

	defer g.Unlock()

	_, err := os.Stat(filepath.Join(g.basedir, ".git"))

	var repo *git.Repository

	if err == nil || os.IsExist(err) {
		repo, err = git.PlainOpen(g.basedir)
	} else {
		repo, err = git.PlainClone(g.basedir, false, &git.CloneOptions{
			URL:               g.url,
			ReferenceName:     "refs/heads/master",
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
	}
	if repo != nil {
		checkIfError(repo.Pull(&git.PullOptions{ReferenceName: "refs/heads/master"}))
	} else {
		checkIfError(err)
	}

}

//FindOne s
func (g *GGit) FindOne(application, profile, label string) *env.Environment {

	g.refresh()

	var files = []string{
		filepath.Join("application-" + profile + ".properties"),
		filepath.Join("application.properties"),
		filepath.Join(label, "application-"+profile+".properties"),
		filepath.Join(label, "application.properties"),
		filepath.Join(label, application+"-"+profile+".properties"),
		filepath.Join(label, application+".properties"),
	}

	sources := []*env.PropertySource{}

	for _, file := range files {
		p, e := properties.LoadFile(filepath.Join(g.basedir, file), properties.UTF8)

		if !checkIfError(e) {
			sources = append(sources, &env.PropertySource{Name: file, Source: p.Map()})
		}
	}
	return &env.Environment{Name: application, Profiles: append([]string{}, profile), PropertySources: sources}
}
