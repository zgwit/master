package manifest

type Maintainer struct {
	Name  string `json:"name,omitempty"` //名称
	Email string `json:"email,omitempty"`
	Url   string `json:"url,omitempty"`
}

type Manifest struct {
	Id          string   `json:"id"`                    //ID
	Name        string   `json:"name"`                  //名称
	Icon        string   `json:"icon,omitempty"`        //图标
	Version     string   `json:"version,omitempty"`     //SEMVER
	Description string   `json:"description,omitempty"` //说明
	Keywords    []string `json:"keywords,omitempty"`    //关键字

	/* 运行信息 */
	Type         string            `json:"type,omitempty"`         //类型：exec 可执行文件，docker 镜像
	Image        string            `json:"image,omitempty"`        //容器镜像：nginx@1.25.3
	Main         string            `json:"main,omitempty"`         //入口：程序文件
	Tags         []string          `json:"tags,omitempty"`         //标签
	Os           []string          `json:"os,omitempty"`           //操作系统支持：linux windows darwin
	Arch         []string          `json:"arch,omitempty"`         //CPU架构：x64 ia32 aarch64
	Dependencies map[string]string `json:"dependencies,omitempty"` //应用和版本

	/* 附加信息 */
	Home        string       `json:"home,omitempty"`
	Bugs        string       `json:"bugs,omitempty"`        //Bug
	License     string       `json:"license,omitempty"`     //软件协议：GPL MIT Apache 。。。
	Maintainers []Maintainer `json:"maintainers,omitempty"` //维护者
}
