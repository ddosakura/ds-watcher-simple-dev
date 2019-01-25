# simple-dev 开发监控工具

## 组成

+ dssdc 用于监控文件变化
+ dssds 用于接收变化数据

## 下载

+ [客户端](https://github.com/ddosakura/ds-watcher-simple-dev/releases)
+ [服务器端](https://cloud.docker.com/u/ddosakura/repository/docker/ddosakura/simple-dev)

### 容器配置示例

```yaml
  dssds:
    image: ddosakura/simple-dev:v0.0.1-beta.1
    volumes:
      - <URL>/dssds/ws:/data/ws
    environment:
      - DSSDS_DB_HOST=mysql
      - DSSDS_WORKSPACE=./data/ws
```

## 依赖

+ 检测主目录 `github.com/mitchellh/go-homedir`
+ CLI `github.com/spf13/cobra`
+ 配置读取 `github.com/spf13/viper`
    + 建议用 yaml `go get -u gopkg.in/yaml.v2`
+ 文件系统 `github.com/spf13/afero`
+ 静态资源 `github.com/rakyll/statik`
+ 文件监控 `go get -u gopkg.in/fsnotify/fsnotify.v1`
+ Windows 启动浏览器 `github.com/inconshreveable/mousetrap`
+ ~~压缩文件 `go get -u github.com/mholt/archiver/cmd/arc`~~

## 参考

+ 20190111: 部分参考 `https://github.com/dengsgo/fileboy`

## 功能

### Done

+ [x] dssdc 客户端
    + [x] init 初始化配置
    + [x] dev 开发模式（更新监控）
        + [x] WebPage页面映射
        + [x] 入口点
        + [x] 监控UI
        + [x] webapi通知
        + [x] API代理
    + [x] package 项目打包
    + [x] publish 发布项目（WebAPI/Git）
        + [x] Git
        + [x] WebAPI Upload
        + [x] `-a`
    + [x] get 获取&查看项目
        + ~~[ ] (考虑在服务器端也监控文件变化{delay需要久一些})~~ <暂时不做>
    + ~~[ ] update 客户端更新(检测github)~~ <暂时不做>

+ [x] dssds 服务器端
    + [x] 数据可视化(与dssdc使用同一个)
    + [x] 数据收集API
    + [x] 项目上传API
        + ~~[ ] 文件覆盖策略~~ <暂时不做>
        + ~~[ ] 自动刷新页面（远程）~~ <暂时不做>
+ ~~[ ] 项目上传API(PHP ver)~~ <暂时不做>

### 注意！

由于相对路径的处理方式未完全确定，请注意一下内容，避免新旧版兼容问题
+ 尽量使用默认配置文件路径
+ 相对路径尽量不要使用 `../`

### TODO & BUG-FIX

+ [ ] 开发者、项目 监控显示细化
+ [ ] 配置更新自动重启
+ [ ] 关于配置文件路径及内部相对路径的处理（相对于程序路径/配置文件路径）
+ [ ] 路由的处理
+ [ ] 处理sqlite要用cgo的问题
+ [ ] 修正另外生成的changeTime(暂时忘记加哪儿了，影响不是特别大)
+ [ ] 端口占用时自增
+ [ ] 归档时的忽略文件

+ [ ] DB文件存在性检查
+ [ ] 监听文件夹存在性检查
