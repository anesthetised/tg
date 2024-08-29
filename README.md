# Telegram Bot API

An incomplete hand-crafted Telegram Bot API client implementation.

## Example

```shell
tg:main* λ go build ./cmd/example
tg:main* λ ./example -token "token"
time=2024-08-29T13:33:43.969+03:00 level=INFO msg="user info" id=1122334455 username=TestBot first_name=Test last_name="" is_bot=true
time=2024-08-29T13:33:49.850+03:00 level=INFO msg=message id=8 time=2024-08-29T13:33:50.000+03:00 chat.id=1122334455 chat.type=private chat.title="" sender.id=1122334455 sender.username=testusername sender.first_name=John sender.last_name=Doe sender.lang_code=en text=test
^Ctime=2024-08-29T13:33:52.870+03:00 level=INFO msg=done
```
