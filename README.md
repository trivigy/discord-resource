# Discord notification resource

Sends messages to [Discord](https://discordapp.com/) channels.

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

Behavior
--------

### `check`: Inactive.

### `in`: Inactive.

### `out`: Sends message to Discord.

Send message to discord, with the configured parameters.

#### Parameters

- `channel` (_required_) Specified which channel to post the message in. The bot must have permissions authorized to `send_messages` for the specific channel. The channel id needs to be provided as a snowflake number in quotes (_string_)
- `message` (_required_) Any text wanted to ultimately appear on the page. 

## Usage Example
```yaml
jobs:
- name: discord-send
  plan:
  - put: discord
    params:
      channel: "((channel_id))"
      message: |
        Hello World!
        This message is from your friendly neighborhood wobot.
```