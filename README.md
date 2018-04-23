# Slack notification sending resource

Sends messages to [Discord](https://discordapp.com/) channels.

## Resource Type Configuration

```yaml
resource_types:
- name: discord-resource
  type: docker-image
  source:
    repository: syncaide/discord-resource
```