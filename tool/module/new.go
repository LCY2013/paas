// Package new generates micro service templates
package new

import (
	"fmt"
	"github.com/LCY2013/paas/tool/tmpl"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/xlab/treeprint"
)

func protoComments() []string {
	return []string{
		"======================explain===========================",
		"The main functions are as follows：",
		"1、Quickly create a common project directory structure。",
		"2、Quickly create a generic API project directory structure。",
		"\n",
		"======================operate===========================",
		"The operation is as follows：",
		"1、make proto ( To modify the docker command under window)",
		"2、exec go mod tidy",
		"3、go run main.go , Check whether it can be started successfully",
		"4、Check whether the registry service exists（Address default：127.0.0.1:8500）",
		"notice：It can also be installed locally (proto ，protoc-gen-go，protoc-gen-micro) run protoc to generate。",
		"\n",
		"========================================================",
	}
}

type config struct {
	// 服务名称
	Alias string
	// 目录地址
	Dir string
	// 在API模式下默认的后端名称
	ApiDefaultServerName string
	// 文件地址
	Files []file
	// 说明
	Comments []string
	// 是否需要创建本地服务名称目录
	CreateAlias bool
}

type file struct {
	//路径
	Path string
	//模板
	Tmpl string
}

func write(c config, file, tmpl string) error {
	fn := template.FuncMap{
		"title": func(s string) string {
			return strings.ReplaceAll(strings.Title(s), "-", "")
		},
		"dehyphen": func(s string) string {
			return strings.ReplaceAll(s, "-", "")
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("f").Funcs(fn).Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}

func create(c config) error {
	// check if dir exists
	if _, err := os.Stat(c.Dir); !os.IsNotExist(err) {
		fmt.Printf("%s: The directory already exists, cannot be created! Please delete and re-create.", c.Dir)
		return fmt.Errorf("%s already exists", c.Dir)
	}

	fmt.Printf("Create initialization project %s\n\n", c.Alias)

	t := treeprint.New()

	// write the files
	for _, file := range c.Files {
		rootDir := c.Dir
		if c.CreateAlias {
			rootDir = c.Alias
		}

		f := filepath.Join(rootDir, file.Path)
		dir := filepath.Dir(f)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0755); err != nil {
				fmt.Println(err)
				return err
			}
		}

		addFileToTree(t, file.Path)
		if err := write(c, f, file.Tmpl); err != nil {
			fmt.Println(err)
			return err
		}
	}

	// print tree
	fmt.Printf(t.String())

	for _, comment := range c.Comments {
		fmt.Printf(comment)
	}

	// just wait
	<-time.After(time.Millisecond * 250)
	fmt.Printf("\n************congratulations! Project initialization succeeded!************\n")
	return nil
}

func addFileToTree(root treeprint.Tree, file string) {
	split := strings.Split(file, "/")
	curr := root
	for i := 0; i < len(split)-1; i++ {
		n := curr.FindByValue(split[i])
		if n != nil {
			curr = n
		} else {
			curr = curr.AddBranch(split[i])
		}
	}
	if curr.FindByValue(split[len(split)-1]) == nil {
		curr.AddNode(split[len(split)-1])
	}
}

func NewServiceProject(ctx *cobra.Command, args []string, createDir bool) error {

	for _, serviceArg := range args {
		serviceSlice := strings.Split(serviceArg, "/")
		serviceName := serviceSlice[len(serviceSlice)-1]
		if len(serviceName) == 0 {
			return fmt.Errorf("%s Service name format error", serviceArg)
		}

		if path.IsAbs(serviceArg) {
			logrus.Info("require relative path as service will be installed in GOPATH")
			return nil
		}

		c := config{
			Alias:    serviceName,
			Dir:      serviceArg,
			Comments: protoComments(),
			Files: []file{
				{"main.go", tmpl.MainSRV},
				//{"generate.go", tmpl.GenerateFile},
				//{"plugin.go", tmpl.Plugin},
				{"handler/" + serviceName + "Handler.go", tmpl.HandlerSRV},
				{"plugin/hystrix/hystrix.go", tmpl.Hystrix},
				{"domain/model/" + serviceName + ".go", tmpl.DomainModel},
				{"domain/repository/" + serviceName + "_repository.go", tmpl.DomainRepository},
				{"domain/service/" + serviceName + "_data_service.go", tmpl.DomainService},
				{"proto/" + serviceName + "/" + serviceName + ".proto", tmpl.ProtoSRV},
				{"Dockerfile", tmpl.DockerSRV},
				{"filebeat.yml", tmpl.Filebeat},
				{"Makefile", tmpl.Makefile},
				{"README.md", tmpl.Readme},
				{".gitignore", tmpl.GitIgnore},
				{"go.mod", tmpl.Module},
			},
			CreateAlias: createDir,
		}
		// create the files

		return create(c)

	}
	return nil
}

func NewApiProject(ctx *cobra.Command, args []string) error {

	for _, serviceArg := range args {
		serviceSlice := strings.Split(serviceArg, "/")
		serviceName := serviceSlice[len(serviceSlice)-1]
		if len(serviceName) == 0 {
			return fmt.Errorf("%s Name cannot be empty", serviceArg)
		}

		if path.IsAbs(serviceArg) {
			logrus.Info("require relative path as service will be installed in GOPATH")
			return nil
		}

		apiDefaultServerName := "XXX"
		//判断在API的状态下默认的服务名称
		if strings.Contains(serviceName, "Api") {
			//替换指定Api的字符为空
			apiDefaultServerName = strings.Replace(serviceName, "Api", "", 1)
		}

		c := config{
			Alias:                serviceName,
			Dir:                  serviceArg,
			ApiDefaultServerName: apiDefaultServerName,
			Comments:             protoComments(),
			Files: []file{
				{"main.go", tmpl.MainAPI},
				//{"generate.go", tmpl.GenerateFile},
				//{"plugin.go", tmpl.Plugin},
				{"handler/" + serviceName + "Handler.go", tmpl.HandlerAPI},
				{"plugin/hystrix/hystrix.go", tmpl.Hystrix},
				{"proto/" + serviceName + "/" + serviceName + ".proto", tmpl.ProtoAPI},
				{"Dockerfile", tmpl.DockerSRV},
				{"filebeat.yml", tmpl.Filebeat},
				{"Makefile", tmpl.Makefile},
				{"README.md", tmpl.ReadmeApi},
				{".gitignore", tmpl.GitIgnore},
				{"go.mod", tmpl.ApiModule},
			},
		}
		// create the files

		return create(c)

	}
	return nil
}
