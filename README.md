### 基于GPT-3.5的公众号自动回复机器人
### 一、介绍
通过`openai`与`wechat`官方的api搭建公众号自动回复的项目．
### 二、Feature
- [x] 解决微信被动回复限制问题。(设定超时时间，滚动返回)
- [x] 支持用户语音输入。（要主动开启，设置与开发->接口权限->接收语音识别结果。已关注用户可能24小时内生效，可重新关注尝试）
- [x] 设置代理
- [x] prompt 提示、max_tokens、temperature 参数调节
- [x] 正则关键词回复(微信在开启服务器配置后，自带的关键字回复就不能用了)。
- [x] 上下文(可以配置记录上下文对话数量上限)

### 三、配置&部署
#### 前置准备
1. 获取`API_KEY`。[OpenAI](https://beta.openai.com/account/api-keys)
2. 获取微信公众号`令牌Token`：[微信公众平台](https://mp.weixin.qq.com/)->基本配置->服务器配置->令牌(Token)(令牌是自定义的，保持一致即可)
3. 下载Release中的二进制执行文件`robot`
```bash
mkdir openai
wget https://github.com/hinak0/openai/releases/latest/download/robot
chmod 555 ./robot
```
4. 修改配置文件
```bash
vim config.yaml
```
#### 配置文件
[配置文件实例&&文档](./config.yaml.example)
#### 启动测试
```bash
./robot
```
#### 常驻服务

使用systemd实现常驻启动
```bash
vim /etc/systemd/system/robot.service
```
##### systemd脚本
```ini
[Unit]
Description=chatgpt with wechat

[Service]
# the user
User=root
Type=simple
CapabilityBoundingSet=CAP_NET_BIND_SERVICE CAP_NET_ADMIN
AmbientCapabilities=CAP_NET_BIND_SERVICE CAP_NET_ADMIN
# 执行文件和配置文件路径
WorkingDirectory=/root/openai/
ExecStart=/root/openai/robot
# 挂掉自动重启
Restart=yes
# 重启间隔
RestartSec=20s

[Install]
WantedBy=multi-user.target
```
##### 设置开机自动启动&&启动服务
```bash
sudo systemctl start robot
sudo systemctl enable robot
```
服务器地址(URL)填写 `http://服务器IP或者域名/`，设置明文方式传输，提交后，点击「启用」。
### 四、二次开发
1. 克隆项目
```bash
git clone https://github.com/hinak0/openai.git
```
2. 安装依赖
```bash
go mod tidy
```
3. 编译
```bash
go build -o robot
# 运行缺少依赖
export CGO_ENABLED=0
go build -o robot main.go
```

### Q&A
有问题请提issue,会看的．
