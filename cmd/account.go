package cmd

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/codegangsta/cli"

	"github.com/dockercn/docker-bucket/global"
	"github.com/dockercn/docker-bucket/models"
)

var CmdAccount = cli.Command{
	Name:        "account",
	Usage:       "通过命令行管理系统的账户",
	Description: "通过命令行添加、激活、停用 Bucket 中的用户账户，账户停用后该账户下公开的 Repository 依旧可以下载。",
	Action:      runAccount,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "action",
			Value: "",
			Usage: "Action 参数: add/active/unactive/log，添加激活的账户/添加待激活的账户/停用账户/查看用户日志",
		},
		cli.StringFlag{
			Name:  "email",
			Value: "",
			Usage: "账户邮件地址",
		},
		cli.StringFlag{
			Name:  "username",
			Value: "",
			Usage: "账户名",
		},
		cli.StringFlag{
			Name:  "passwd",
			Value: "",
			Usage: "账户初始密码",
		},
		cli.StringFlag{
			Name:  "conf",
			Value: "",
			Usage: "Web 服务的配置文件路径",
		},
	},
}

func runAccount(c *cli.Context) {
	var action, email, username, passwd string
	var err error

	basePath, _ := os.Getwd()

	//如果外部指定了配置文件就不读取 include::Bucket 指定的配置文件
	//读取 Bucket 的单独配置
	if len(c.String("conf")) > 0 {
		if global.BucketConfig, err = config.NewConfig("ini", c.String("conf")); err != nil {
			beego.Error("[Application] 读取配置文件错误: " + err.Error())
		}
	} else {
		if global.BucketConfig, err = config.NewConfig("ini", fmt.Sprintf("%s/%s", basePath, beego.AppConfig.String("include::Bucket"))); err != nil {
			beego.Error("[Application] 读取配置文件错误: " + err.Error())
		}
	}

	if len(c.String("action")) > 0 {
		models.InitDb()
		action = c.String("action")
		switch action {
		case "add":
			if len(c.String("username")) > 0 && len(c.String("email")) > 0 && len(c.String("passwd")) > 0 {
				username = c.String("username")
				email = c.String("email")
				passwd = c.String("passwd")

				user := new(models.User)
				if err := user.Put(username, passwd, email); err != nil {
					fmt.Println(fmt.Sprintf("%s: %s", "添加用户失败", err.Error()))
				} else {
					fmt.Println(fmt.Sprintf("添加 %s 用户成功！", username))
					//TODO 发送注册邮件
				}

			} else {
				fmt.Println("account add 命令必须指定 username/email/passwd 参数")
			}

			break
		case "active":
			break
		case "unactive":
			break
		case "log":
			break
		default:
			fmt.Println("目前只支持 add/active/unactive 三个指令。")
		}
	} else {
		fmt.Println("需要指定操作用户的指令，仅支持 add/active/unactive 三个指令")
	}
}
