package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/autumnzw/hiweb/webcmd"
	"gitlab.bjictc.com/go/logsd"
	"gitlab.bjictc.com/go/sqlham"
)

func main() {
	logsd.SetLogger(logsd.AdapterConsole)
	args := os.Args
	if len(args) < 2 {
		logsd.Error("no param:gen swag")
		return
	}
	cmdType := args[1]
	switch cmdType {
	case "init":
		if len(args) < 3 {
			logsd.Error("no param:projectName")
			return
		}
		projectName := args[2]
		has, err := webcmd.PathExists(projectName)
		if has && err == nil {
			logsd.Error("dir exists " + projectName)
			return
		}
		os.Mkdir(projectName, 0666)
		os.Mkdir(projectName+"/controllers", 0666)
		os.Create(projectName + "/controllers/hello.go")
		os.WriteFile(projectName+"/controllers/hello.go",
			[]byte(`package controllers
import (
	"net/http"

	"github.com/autumnzw/hiweb"
)

type Hello struct {
	hiweb.Controller
}

// @Description string out
func (h *Hello) H() {
	d := make(map[string]string)
	d["aa"] = "aa"
	h.ServeJSON(http.StatusOK, d)
	return
}`), 0666)

		os.Create(projectName + "/main.go")
		os.WriteFile(projectName+"/main.go",
			[]byte(`package main

import (
	_ "`+projectName+`/controllers"
	"fmt"

	"flag"
	"net/http"

	"github.com/autumnzw/hiweb"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "ip")
	port := flag.String("port", "8111", "port")
	flag.Parse()
	if *port == "" {
		flag.PrintDefaults()
		return
	}
	hiweb.WebConfig.SecretKey = "aaaa"
	fmt.Printf("serve web:" + *ip + ":" + *port)
	e := http.ListenAndServe(*ip+":"+*port, nil)
	if e != nil {
		fmt.Printf("err:%s", e)
	}
}
		`), 0666)
		swag("./"+projectName+"/controllers", projectName)

		cmd := exec.Command("go", "mod", "init", projectName)
		cmd.Dir = "./" + projectName
		err = cmd.Start()
		if err != nil {
			logsd.Error("err:%s", err)
		}

		cmd = exec.Command("go", "mod", "tidy")
		cmd.Dir = "./" + projectName
		err = cmd.Start()
		if err != nil {
			logsd.Error("err:%s", err)
		}
		logsd.Info("cd " + projectName)
		logsd.Info("go run ./main.go")
	case "gen":
		if len(args) < 3 {
			logsd.Error("no param:mysql modelDir")
			return
		}
		mysqlUrl := args[2]
		modelDir := args[3]
		sqlham.InitMysqlDb(mysqlUrl, logsd.GetLoggerSd())
		tables, err := sqlham.GetTables()
		if err != nil {
			logsd.Error("err:%v", err)
			return
		}
		err = sqlham.AutoColumn(tables, modelDir)
		if err != nil {
			logsd.Error("err:%v", err)
			return
		}
	case "swag":
		controllersDir := ""
		projectName := ""
		if len(args) < 3 {
			controllersDir = "./controllers"
			projectName, _ = os.Getwd()
			projectName = filepath.Base(projectName)
		} else {
			projectName = args[2]
			controllersDir = args[3]
		}
		swag(controllersDir, projectName)
	default:
		logsd.Error("not support " + cmdType)
	}
}

func swag(controllersDir, projectName string) {
	err := webcmd.CreateRoute(controllersDir, projectName, "", "", controllersDir)
	if err != nil {
		logsd.Error(err)
	}
}
