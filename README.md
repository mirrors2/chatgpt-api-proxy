# chatgpt-api-proxy

<a title="Docker Image CI" target="_blank" href="https://github.com/mirrors2/chatgpt-api-proxy/actions"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/mirrors2/chatgpt-api-proxy/ci.yaml?label=Actions&logo=github&style=flat-square"></a>
<a title="Docker Pulls" target="_blank" href="https://hub.docker.com/r/mirrors2/chatgpt-api-proxy"><img src="https://img.shields.io/docker/pulls/mirrors2/chatgpt-api-proxy.svg?logo=docker&label=docker&style=flat-square"></a>

chatgpt api 代理
## 快速开始
```
docker run -d -p 80:80 --name chatgpt-api-proxy mirrors2/chatgpt-api-proxy
```

## docker-compose

```
version: '3.7'
services: 
  chatgpt-api-proxy:
    image: mirrors2/chatgpt-api-proxy
    container_name: chatgpt-api-proxy 
    restart: unless-stopped
    ports: 80:80
```
# 测试
```
curl http://your.domains/v1/chat/completions \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer sk-yourrrgjhSJjx4bsSJjYsT3BlbkFJyMQjAH3sBcUzvGYFyGcl' \

  -d '{
  "model": "gpt-3.5-turbo",
  "messages": [{"role": "user", "content": "我爱你,翻译成英语"}]
}'
```

# License
[MIT](./LICENSE) License.