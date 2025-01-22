package clone

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type Repo struct {
	RepoName string `survey:"name"`
	DirName  string `survey:"dir"`
}

var (
	// repoURL string
	branch string
)

var CmdClone = &cobra.Command{
	Use:     "clone",
	Example: "nunu clone https://gitee.com/jianlu8023/nunu.git",
	Short:   "clone a repo.",
	Long:    `clone a repo.`,
	Run:     run,
}

func init() {
	// CmdClone.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")
	CmdClone.Flags().StringVarP(&branch, "branch", "b", branch, "repo branch")

}

func NewRepo() *Repo {
	return &Repo{}
}

func run(cmd *cobra.Command, args []string) {
	p := NewRepo()
	if len(args) == 0 {
		err := survey.AskOne(&survey.Input{
			Message: "What is your repo url?",
			Help:    "repo url.",
			Suggest: nil,
		}, p.RepoName, survey.WithValidator(survey.Required))
		if err != nil {
			return
		}
	} else {
		p.RepoName = args[0]
	}

	// 将dirName 进行设置
	parsedURL, err := url.Parse(p.RepoName)
	if err != nil {
		return
	}
	repoName := path.Base(parsedURL.Path)

	if strings.HasSuffix(repoName, ".git") {
		repoName = strings.TrimSuffix(repoName, ".git")
	}
	p.DirName = repoName

	// clone repo
	yes, err := p.cloneTemplate()
	if err != nil || !yes {
		return
	}
	p.rmGit()
	fmt.Println("clone success")
}

func (p *Repo) cloneTemplate() (bool, error) {
	stat, _ := os.Stat(p.DirName)
	if stat != nil {
		var overwrite = false

		prompt := &survey.Confirm{
			Message: fmt.Sprintf("Folder %s already exists, do you want to overwrite it?", p.DirName),
			Help:    "Remove old repo and clone new repo.",
		}
		err := survey.AskOne(prompt, &overwrite)
		if err != nil {
			return false, err
		}
		if !overwrite {
			return false, nil
		}
		err = os.RemoveAll(p.DirName)
		if err != nil {
			fmt.Println("remove old repo error: ", err)
			return false, err
		}
	}

	fmt.Printf("git clone %s\n", p.RepoName)
	var cmd *exec.Cmd

	if len(branch) == 0 {
		cmd = exec.Command("git", "clone", p.RepoName, p.DirName)
	} else {
		cmd = exec.Command("git", "clone", "-b", branch, p.RepoName, p.DirName)
	}
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("git clone %s error: %s\n", p.RepoName, err)
		return false, err
	}
	return true, nil
}

func (p *Repo) rmGit() {
	os.RemoveAll(p.DirName + "/.git")
}
