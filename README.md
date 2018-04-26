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

Behavior
--------

### `check`: Inactive.

### `in`: Inactive.

### `out`: Sends message to Discord.

Send message to discord, with the configured parameters.

#### Parameters: `params`

- `channel` (_required_) Specifies which channel to post the message in. The bot must have permissions authorized to send_messages for the specific channel. The channel id needs to be provided as a snowflake number in quotes (_string_)
- `color` (_required_) Indicates what color the embed should be marked with. If no color is specified black will be used. (Not sure how to pass hex values via concourse so just look up the hex value and calculate what integer it is.) 
- `title` (_required_) Any text wanted to ultimately appear on the page as the title of the message. 
- `message` (_required_) The text that will be inside of the body of the message. 

## Usage Example
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