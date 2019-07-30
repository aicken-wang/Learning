## VSCode golang插件安装
在 GOPATH 目录下的src目录下新建 golang.org 再进入golang.org目录新建x 目录
cd GOPATH/src/golang.org/x
使用git 下载插件源码 
```git
git clone https://github.com/golang/tools.git tools
git clone https://github.com/golang/lint.git
```
- 通过VSCode 快捷键Ctrl/Command+Shift+P
- 打开命令输入框 输入` Go:Install/Update Tools`
- 当看到success及安装成功 GOPATH下会生可执行的插件文件