# chatgpt-api-proxy 

<a title="Docker Image CI" target="_blank" href="https://github.com/mirrors2/chatgpt-api-proxy/actions"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/mirrors2/chatgpt-api-proxy/ci.yaml?label=Actions&logo=github&style=flat-square"></a>
<a title="Docker Pulls" target="_blank" href="https://hub.docker.com/r/mirrors2/chatgpt-api-proxy"><img src="https://img.shields.io/docker/pulls/mirrors2/chatgpt-api-proxy.svg?logo=docker&label=docker&style=flat-square"></a>

# (后续没有新功能,推荐使用[opencatd-open](https://github.com/mirrors2/opencatd-open))

chatgpt api 代理,已验证OpenCat,AssisChat,AMA(问天),chathub

可配置好OPENAI_API_KEY分享代理地址给他人用.
## 快速开始
```
docker run -d -p 80:80 --name chatgpt-api-proxy mirrors2/chatgpt-api-proxy

可选 -e OPENAI_API_KEY={nide_api_key}
```

## docker-compose

```
version: '3.7'
services: 
  chatgpt-api-proxy:
    image: mirrors2/chatgpt-api-proxy
    container_name: chatgpt-api-proxy 
    restart: unless-stopped
    ports:
      - 80:80
    # environment:
    # 自定义apikey，可分享给别人用
    #   - OPENAI_API_KEY={openai_api_key}

```
or

```
wget https://github.com/mirrors2/chatgpt-api-proxy/raw/main/docker/docker-compose.yml
```
# 测试
```
curl https://chatgpt.gopher.ink/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d '{
  "model": "gpt-3.5-turbo",
  "messages": [{"role": "user", "content": "我爱你,翻译成英语"}]
}'
```

# License
[MIT](./LICENSE) License.