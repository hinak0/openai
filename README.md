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
wget https://github.com/hinak0/openai/releases/download/v1.0.0/robot
chmod 555 ./robot
```
4. 修改配置文件
```bash
vim config.yaml
```
#### 配置文件
```yaml
http:
  # 监听地址，一般不用改
  addr: 0.0.0.0
  # 必填: 端口号
  port: 80
  # 可选: 代理地址。需要你有本地或远程代理软件，举例: socks5://127.0.0.1:7890 或 7891
  proxy:

session:
  # 是否开启上下文
  enable: true
  # 存储后端，目前只支持redis
  type: redis
  addr: example.com:6379
  passowrd:
  database: 0
  # 记录对话数量，越多越消耗token
  track: 5

openai:
  # 必填: KEY。 文档: https://platform.openai.com/account/api-keys
  key: xxx
  # 可选: 参数调节
  params:
    # openai的接口地址，放出来是因为有些人做了反向代理，要注意这有安全问题，谨慎使用
    api: https://api.openai.com/v1/chat/completions
    # 暂时请使用 gpt-3.5-turbo
    model: gpt-3.5-turbo
    # 提示。 可以理解为对其身份设定。 文档: https://platform.openai.com/docs/guides/chat/introduction
    # 每个问题都会携带，注意，它也占用token消耗。
    prompt:
    # 影响 问题+回复的长度。  gpt-3.5模型最大4096， 非1个汉字1token
    maxTokens: 1024
    # 温度。 0-2 。较高的值将使输出更加随机，而较低的值将使输出更加集中和确定。
    temperature: 0.8
  # 限制用户问题最大长度。这个以字计算，非token.
  maxQuestionLength: 200

wechat:
  # 必填(公众号服务). 与公众号设置保持一致
  token: xxx
  # 影响滚动返回结果 (5s-13s)
  timeout: 7
  # 用户关注时主动发送的消息
  subscribeMsg: "关注自动发送的消息"
  # 关键字自动回复，key是关键字,值是对应回复，支持正则
  keyword:
    关键字: 自动回复内容
    # 例如
    粉丝群|^反馈: 反馈请联系hinak0@qq.com
```
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
