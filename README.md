# chatgpt-api-proxy



curl https://yourdomains/v1/chat/completions \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer sk-nidemiyuezsbdzsbdzsbdzsbd' \
  -d '{
  "model": "gpt-3.5-turbo",
  "messages": [{"role": "user", "content": "谢谢,翻译成英语"},{"role":"assistant","content":"Thanks"},{"role": "user", "content": "日语"}]
}'