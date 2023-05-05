### 基于GPT-3.5的公众号自动回复机器人
### 持续优化中，喜欢的同学给个🌟关注一下

### 一、介绍
- 说明
  - 这是一个用于**公众号自动回复机器人**的项目。需要你有 OpenAI 账号、公众号、海外服务器或代理。
  - 花费。`OpenAI`账号赠送18$，限期使用。按字算钱，0.002刀 /1000 tokens 。
  - 观点。我觉得公众号不是一个好的使用场景，订阅号入口麻烦，服务号需要资质且风险更大。所以如果你只是玩玩可以部署。
  - 体验。关注公众号`杠点杠`尝试提问，这仅是个人娱乐号，不推送。   别问预测和实事问题，它不会。

### 二、Feature
- [x] 解决微信被动回复限制问题。(设定超时时间，滚动返回)
- [x] 支持用户语音输入。（要主动开启，设置与开发->接口权限->接收语音识别结果。已关注用户可能24小时内生效，可重新关注尝试）
- [x] 设置代理
- [x] prompt 提示、max_tokens、temperature 参数调节
- [x] [敏感词](https://github.com/tomatocuke/sieve)检测及自定义添加。(不清楚这样是否降低风险。代码内置隐藏了一些敏感词，你也可以启动时在根目录添加`keyword.txt`自定义敏感词。  如有敏感词误杀，你可以向我反馈)
- [ ] 上下文。(其实开发也不算难。主要是OpenAI不记录会话，上下文的本质是把之前的QA都作为新的参数传过去，这会叠加消耗token)
- [ ] 用户身份验证。(待开发)
- [ ] 关键词回复。(待开发)

### 三、部署
1. 获取`API_KEY`。[OpenAI](https://beta.openai.com/account/api-keys) （如果访问被拒绝，注意全局代理，打开调试，Application清除LocalStorage后刷新，实测可以）
2. 获取微信公众号`令牌Token`：[微信公众平台](https://mp.weixin.qq.com/)->基本配置->服务器配置->令牌(Token)
3. 下载Release中的二进制执行文件，然后复制本项目中的`config.yaml.example`,在二进制执行文件的同目录中新建文件`config.yaml`，将复制的内容粘贴进去，修改为自己的自定义配置
4. 启动服务
```bash
./robot
# 或者常驻运行
nohup ./robot 2>&1 &
```
5. 服务器地址(URL)填写 `http://服务器IP/`，设置明文方式传输，提交后，点击「启用」。
### 四、二次开发
1. 克隆项目
```bash
git clone https://github.com/hinak0/openai.git
```
1. 安装依赖
```bash
go mod tidy
```
1. 编译
```bash
go build -o robot
# 运行缺少依赖
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o robot main.go
```

### 四、QA
- 日志出现 `openai请求超时` <br>
答：对openai的请求发不过去

- 出现报错  `connection reset by peer` 或 `Post "https://api.openai.com/v1/chat/completions": EOF`  <br>
答：是否使用了代理呢？ 大概率是IP被多人使用的结果，换个IP，但是部署不建议使用代理的方式，不稳定。
