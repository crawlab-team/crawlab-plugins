package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func EnvInit() {

	log.Info("python 一键安装开始")
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "init --env=python --version=3.6.5",
			Action: (&Command{}).PythonEnvInit,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "env", Usage: "--env"},
				cli.StringFlag{Name: "version", Usage: "--version"},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("command error :" + err.Error())
	}

}

type Command struct {
}

func (this *Command) PythonEnvInit(cli *cli.Context) {
	var cmd *exec.Cmd
	env := cli.String("env")
	version := cli.String("version")
	fmt.Println(env, version)

	cmd = exec.Command("ping", "-c", "1", "-w", "1", "www.baidu.com")

	pingInfo, err := cmd.Output()
	if err != nil {
		log.Fatalf("网络检测失败程序退出，请重新检测网络环境", string(pingInfo))
	}
	log.Info("网络连接正常")

	yumInfo, err := exec.Command("yum", "list").Output()
	if err != nil {
		log.Fatalf("yum 不可用，请先手动配置yum安装!", string(yumInfo))
	}

	gccInfo, err := exec.Command("yum", "install", "gcc", "-y").Output()
	if err != nil {
		log.Fatalf("gcc install failed", string(gccInfo))
	}

	dependPackage, err := exec.Command("yum", "install", "openssl-devel", "bizp2-devel", "expat-devel", "gdbm-devel", "readline-devel", "sqlite-devel", "libffi-devel", "-y").Output()
	if err != nil {
		log.Fatalf("depend package like openssl-devel install failed", string(dependPackage))
	}
	pythonDownInfo, err := exec.Command("wget", "http://npm.taobao.org/mirrors/python/"+version+"/Python-"+version+".tgz").Output()
	if err != nil {
		log.Fatalf("Download python package failed", string(pythonDownInfo))
	}

	tarInfo, err := exec.Command("tar", "-zxvf", "Python-"+version+".tgz").Output()
	if err != nil {
		log.Fatalf("tar -zxvf failed", string(tarInfo))
	}
	mvInfo, err := exec.Command("bash", "-c", "mv "+"Python-"+version+" /usr/local/").Output()
	if err != nil {
		log.Fatalf("mv Python+version to /usr/local/ failed", string(mvInfo))
	} else {
		log.Info("#########")
	}
	//cdInfo, err := exec.Command("cd", "/usr/local/"+"Python-"+version).Output()
	//if err != nil {
	//	log.Fatalf("cd python dir failed", string(cdInfo))
	//}
	err = os.Chdir("/usr/local/" + "Python-" + version)
	if err != nil {
		log.Fatalf("cd Python dir failed", err)
	} else {
		log.Info("cd /usr/local/Python-" + version)
	}

	mkconfigInfo, err := exec.Command("/usr/local/Python-" + version + "/configure").Output()
	if err != nil {
		log.Fatalf("configure failed", string(mkconfigInfo))
	} else {
		log.Info("configure success")
	}

	makeInfo, err := exec.Command("bash", "-c", "make && make install").Output()
	if err != nil {
		log.Fatalf("make install python failed", string(makeInfo))
	} else {
		log.Info("make install success")
	}

	eperlInfo, err := exec.Command("sudo", "yum", "-y", "install", "epel-release", "python-pip").Output()
	if err != nil {
		log.Fatalf("eperlInfo install failed", string(eperlInfo))
	}
	cpInfo, err := exec.Command("cp", "python", "/bin/python3").Output()
	if err != nil {
		log.Fatalf("cp python /bin/python3 failed", string(cpInfo))
	}

	testPython, err := exec.Command("python3", "-V").Output()
	if err != nil {
		log.Fatalf("python-"+version+"install failed", string(testPython))
	} else {
		log.Info("python-" + version + "install success")
	}
}

func main() {
	EnvInit()
}
