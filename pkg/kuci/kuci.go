package kuci

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Controller struct
type Controller struct {
}

// NewController function
func NewController() *Controller {
	c := &Controller{}
	return c
}

func (c *Controller) Start() {
	for {
		gitURL := "git@github.com:zzh8829/kuci.git"
		tagString := "zihao/play:kuci-${GIT_SHA_SHORT}"

		sha, err := shellCommand(fmt.Sprintf("git ls-remote %v HEAD | head -c7", gitURL))
		if err != nil {
			log.Printf("%v", err)
			log.Fatal("Rip")
		}
		mapper := func(key string) string {
			switch key {
			case "GIT_SHA_SHORT":
				return string(sha)
			}
			return ""
		}
		imageTag := os.Expand(tagString, mapper)
		if !dockerImageExists(imageTag) {
			log.Infof("Image %v no exist\n", imageTag)
			err := doCI(gitURL, imageTag)
			if err != nil {
				log.Errorf("%v", err)
			}
		}
		return
		time.Sleep(10 * time.Second)
	}
}

func dockerImageExists(tag string) bool {
	splits := strings.Split(tag, ":")
	image, tag := splits[0], splits[1]
	resp, err := http.Get(fmt.Sprintf("https://index.docker.io/v1/repositories/%v/tags/%v", image, tag))
	if err == nil && resp.StatusCode == 200 {
		return true
	}
	return false
}

func shellCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	reader := io.MultiReader(stdout, stderr)
	in := bufio.NewScanner(reader)

	outputs := ""
	for in.Scan() {
		outputs += in.Text()
		log.Printf(in.Text())
	}

	if err := in.Err(); err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return outputs, nil
}

func doCI(gitURL string, imageTag string) error {
	dir, err := ioutil.TempDir("", "kuci")
	if err != nil {
		return err
	}
	log.Printf(dir)
	// defer os.RemoveAll(dir)
	gitDir := path.Join(dir, "repo")

	_, err = shellCommand(fmt.Sprintf("git clone %v %v", gitURL, gitDir))
	if err != nil {
		return err
	}

	os.Chdir(gitURL)

	_, err = shellCommand(fmt.Sprintf("docker build -t %v .", imageTag))
	if err != nil {
		return err
	}

	_, err = shellCommand(fmt.Sprintf("docker push %v", imageTag))
	if err != nil {
		return err
	}

	_, err = shellCommand(fmt.Sprintf("kubectl set image deployment kuci %v", imageTag))
	if err != nil {
		return err
	}

	return nil
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/kuci")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
