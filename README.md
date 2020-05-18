# Discord notification resource

Sends messages to [Discord](https://discordapp.com/) channels.

> Please note that I do not have automated deployment to docker hub yet and 
will set it up in the future. For now I do it regularly when I update the repo.

## Resource Type Configuration

```yaml
resource_types:
- name: discord-resource
  type: docker-image
  source:
    repository: trivigy/discord-resource
```

## Resource Configuration Example
```yaml
- name: discord
  type: discord-resource
  check_every: 999999h
  source:
    token: ((token))
```

## Behavior

### `check`: Not implemented.

### `in`: Not implemented.

### `out`: Send message to a Discord channel.

Send message to a discord channel with the configured parameters. Parameters can be passed in using the [params](https://concourse-ci.org/jobs.html#schema.step.put-step.params) key on the `put` step or passed in via files.

#### Parameters: `params`

**One of either the `_file` or non-`_file` parameters are required.**

The `_file` parameters take precedence over whatever is set in the `params` key of the `put` step.

- `channel` (_string_): Specifies which channel to post the message in. The bot must have permissions authorized to send_messages for the specific channel. The channel id needs to be provided as a snowflake number in quotes (_string_)
- `channel_file` (_string_): Specifies which channel to post the message in. The bot must have permissions authorized to send_messages for the specific channel. The channel id needs to be provided as a snowflake number in quotes (_string_)
- `color` (_int_): Indicates what color the embed should be marked with. If no color is specified black will be used. (Not sure how to pass hex values via concourse so just look up the hex value and calculate what integer it is.)
- `color_file` (_string_): Indicates what color the embed should be marked with. If no color is specified black will be used. (Not sure how to pass hex values via concourse so just look up the hex value and calculate what integer it is.)
- `title` (_string)_: Any text wanted to ultimately appear on the page as the title of the message.
- `title_file` (_string_): Path to file containing the text wanted to ultimately appear on the page as the title of the message.
- `message` (_string_): The text that will be inside the body of the message.
- `message_file` (_string_): Path to file containing the text that will be inside the body of the message.

## Usage Examples
```yaml
jobs:
- name: discord-send
  plan:
  - put: discord
    params:
      channel: "((channel_id))"
      color: 6076508
      title: Hello World!
      message: |
        Success
```

Using the `_file` params. There's a `task` step before the `put` which generates the `discord-message` output.

```yaml
jobs:
- name: discord-send
  plan:
  - task: generate-discord-message
    file: tasks/generate-discord-message.sh
  - put: discord
    params:
      channel_file: "discord-message/channel"
      color_file: "discord-message/color"
      title_file: "discord-message/title"
      message_file: "discord-message/message"
```
