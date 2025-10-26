# Telegram Bot

open the following url in your browser
> https://web.telegram.org/a/#93372553

search `BotFather`

> type `/newbot` and send
Alright, a new bot. How are we going to call it? Please choose a name for your bot.

> type `aibootcamp` and send

Good. Now let's choose a username for your bot. It must end in `bot`. Like this, for example: TetrisBot or tetris_bot.

>  type `aiagentbootcamp_bot` and send

Done! Congratulations on your new bot. You will find it at [https://t.me/aiagentbootcamp_bot](https://web.telegram.org/a/#8280699199)


> Use this token to access the HTTP API:
`8280699199:AAHgFfH3yOrs8SPqLtD5mxOBCIH8AD6UyWw`


### Make the url using your access token:
> https://api.telegram.org/bot`8280699199:AAHgFfH3yOrs8SPqLtD5mxOBCIH8AD6UyWw`/getUpdates

go to n8n website and send a text message using Telegram node and open up the above url in your browser to see the json response, find `chat` id
in my case my chat id is: `792322732`

```json
{"ok":true,"result":[{"update_id":860163980,"message":{"message_id":1,"from":{"id":792322732,"is_bot":false,"first_name":"Mostain","last_name":"Billah","username":"MostainBillah","language_code":"en"},"chat":{"id":792322732,"first_name":"Mostain","last_name":"Billah","username":"MostainBillah","type":"private"},"date":1760941707,"text":"/start","entities":[{"offset":0,"length":6,"type":"bot_command"}]}},{"update_id":860163981,"message":{"message_id":2,"from":{"id":792322732,"is_bot":false,"first_name":"Mostain","last_name":"Billah","username":"MostainBillah","language_code":"en"},"chat":{"id":792322732,"first_name":"Mostain","last_name":"Billah","username":"MostainBillah","type":"private"},"date":1760941715,"text":"Bismillah"}}]}
```


### Detect the message type using an IF or Switch node
You can check which key exists under message:

| Message Type      | Key present in message object | Example check in n8n                               |
| ----------------- | ----------------------------- | -------------------------------------------------- |
| **Text**          | `message.text`                | `{{ $json["message"]["text"] !== undefined }}`     |
| **Voice**         | `message.voice`               | `{{ $json["message"]["voice"] !== undefined }}`    |
| **Photo/Image**   | `message.photo`               | `{{ $json["message"]["photo"] !== undefined }}`    |
| **Document/File** | `message.document`            | `{{ $json["message"]["document"] !== undefined }}` |
| **Sticker**       | `message.sticker`             | `{{ $json["message"]["sticker"] !== undefined }}`  |
