
# Panda Bot

## Features
- Simple and fun commands

- Implements several free api for silly commands.

- `!imbored` will ping the messenger with a fun and unique activity idea for them

- `!convert <url>` turns a link to an image into ascii art and posts it on pastebin using python scripts

- `!coins <crypto name here>` will give you the current data( price, change within last 30m, 1 day, week )

- `!coins remindme <crypto name here> <amount of time> <unit of time>` will give you an update a after the set amount of the unit of time.

- See the rest with a simple `!help` command in a server hosting the discord bot.

## Hosting
### Prerequisites
You will need ot install the following:
- [Python 3](https://www.python.org/downloads/) (with [pillow](https://pillow.readthedocs.io/en/stable/installation.html#basic-installation)) - This is for the `!convert` command. If you don't install it the command will crash the program.


### Setup
First thing you need to do is run the application once, then close it, this will generate a default config.json file in the same directory. This contains your various tokens and bot prefix

The json should look like this:
```json
{
    "DiscordBotToken":"",
    "PastebinToken":"",
    "BotPrefix":""
}
```

| Name | Description |
|-|-|
| "DiscordBotToken" | Change this to the bot token you generate from: https://discord.com/developers/applications |
| "PastebinToken" | Change this to the token from: https://pastebin.com/doc_api (requires account)|
| "BotPrefix" | This is the prefix on commands, change it to what you like. Normally bots use `!<command>`, in that case change it to `!` |

### Fire it up!
The next thing to do is to start it up and it will appear to be online on all servers your bot token is invited to. If you need help inviting it to your server check on the discord dev portal, or get in contact with me on discord @ `panda#4464`

Now the bot is ready to use, go ahead and try the command `<BotPrefix>help`
This will give you the included commands. 