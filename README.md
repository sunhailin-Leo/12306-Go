<h1 align="center">Go-12306</h1>
<p align="center">
    <em>Use go-resty to crawl 12306</em>
</p>
<p align="center">
    <a href="https://github.com/sunhailin-Leo">
        <img src="https://img.shields.io/badge/Author-sunhailin--Leo-blue" alt="Author">
    </a>
</p>
<p align="center">
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="License">
    </a>
</p>

## 💯 项目说明

* 项目包管理基于 [govendor](https://github.com/kardianos/govendor) 构建，项目使用了 [go-resty](https://github.com/go-resty/resty) 作为 HTTP 请求框架
* 打包文件在 `pkg` 文件夹中（darwin 对应 Mac OS，linux 对应 Linux 系统，win64 对应 Windows 64位系统）

## 💻 使用说明

**Linux / Mac OS 下使用**
```shell script
# Linux / Mac OS
chmod a+x go_12306
# 查询两地车次信息
./go_12306 schedule <起始站名> <到达站名> <当前日期(日期格式: YYYY-MM-DD)>
# 查询某车次时刻表
./go_12306 info <车次号(例如: G1)> <当前日期(日期格式: YYYY-MM-DD)>
```

**Windows 下使用**
```bash
# Windows 下
# 查询两地车次信息
go_12306.exe schedule <起始站名> <到达站名> <当前日期(日期格式: YYYY-MM-DD)>
# 查询某车次时刻表
go_12306.exe info <车次号(例如: G1)> <当前日期(日期格式: YYYY-MM-DD)>
```

**车次时刻表**
![](https://user-images.githubusercontent.com/17564655/67031455-29a10780-f144-11e9-9180-862d8a368595.png)

**两地车次信息**
![](https://user-images.githubusercontent.com/17564655/67031522-4fc6a780-f144-11e9-86de-9bceb86a4936.png)

## 📖 功能说明

* 目前暂时开发了两个功能:
    * 查询两地车次信息
    * 查询某车次时刻表

* 后续开发功能点:
    * 加入代理配置
    * 争取完善一些命令行交互以及其他 12306 的功能

## 📃 License

MIT [©sunhailin-Leo](https://github.com/sunhailin-Leo)