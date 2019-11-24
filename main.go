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
		log.Fatalf("gcc install failed", gccInfo)
	}

	dependPackage, err := exec.Command("yum", "install", "openssl-devel", "bizp2-devel", "expat-devel", "gdbm-devel", "readline-devel", "sqlite-devel", "libffi-devel", "-y").Output()
	if err != nil {
		log.Fatalf("depend package like openssl-devel install failed", dependPackage)
	}
	pythonDownInfo, err := exec.Command("wget", "http://npm.taobao.org/mirrors/python/"+version+"/Python-"+version+".tgz").Output()
	if err != nil {
		log.Fatalf("Download python package failed", pythonDownInfo)
	}

	makeDirInfo, err := exec.Command("mkdir", "/usr/local/python3").Output()
	if err != nil {
		log.Fatalf("make python3 dir failed", makeDirInfo)
	}

	tarInfo, err := exec.Command("tar", "-zxvf", "Python-"+version+".tgz", "&&", "cd", "Python-"+version).Output()
	if err != nil {
		log.Fatalf("tar -zxvf failed", tarInfo)
	}
	mkconfigInfo, err := exec.Command("./configure", "--prefix=/usr/local/python3").Output()
	if err != nil {
		log.Fatalf("configure failed", mkconfigInfo)
	}

	makeInfo, err := exec.Command("make", "&&", "make", "install").Output()
	if err != nil {
		log.Fatalf("make install python failed", makeInfo)
	}
	lnPythonInfo, err := exec.Command("ln", "-s", "/usr/local/python3/bin/python3", "/usr/bin/python3").Output()
	if err != nil {
		log.Fatalf("ln python failed", lnPythonInfo)
	}
	lnPipInfo, err := exec.Command("ln", "-s", "/usr/local/python3/bin/pip3", "/usr/bin/pip3").Output()
	if err != nil {
		log.Fatalf("ln pip failed", lnPipInfo)
	}
	eperlInfo, err := exec.Command("sudo", "yum", "-y", "install", "epel-release", "python-pip").Output()
	if err != nil {
		log.Fatalf("eperlInfo install failed", eperlInfo)
	}

	testPython, err := exec.Command("python3", "-V").Output()
	if err != nil {
		log.Fatalf("python-"+version+"install failed", testPython)
	} else {
		log.Info("python-" + version + "install success")
	}
}

func main() {
	EnvInit()
}
