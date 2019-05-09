# SlackMessage Parser

This program filters through a dump of slack messages exported from your slack, by a user ID. So you can have a history of everything a person has said in a slack workspace.
Motivation of building this was to record in history what our dear friend Kenny has said ever in our slack workspace.

# Usage

Slack folder structure:
```shell
|---rootDataFolder
|   |--slackchannel
|   |   |--2017-01-01.json
|   |   |__YYYY-MM-DD.json
|   |--anotherslackchannel
|   |   |--YYYY-MM-DD.json
|   |   |__YYYY-MM-DD.json
```
Remove `channels.json`, `integration_logs.json` and `users.json` that are in the root dir of the slack dump.

**Running**
```shell
go run filterslackmessages [full_path_to_slack_dump] [user_id_from_slack_workspace]
```

It will output a text file with the username as the filename eg: `UXXXXXXX.txt`