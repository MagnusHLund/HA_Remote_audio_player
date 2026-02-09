# Home assistant remote audio player

## Configuration

### Env

Create an .env file, in the root directory of the project.
The file should contain the following mandatory fields:

```
MQTT_HOST=
MQTT_PORT=
```

And can include the following optional fields:

```
MQTT_USERNAME=
MQTT_PASSWORD=
MQTT_CLIENT_ID=ha-remote-audio-player
MQTT_TOPIC_PREFIX=ha/remote-audio
MQTT_DISCOVERY=true
MQTT_DEVICE_ID=ha-remote-audio-player
MQTT_DEVICE_NAME=HA Remote Audio Player
```

You only need to include the optional fields, if you wish to change them.
Their default values are the same as in the snippet above.

### audio.json

From the root of the project, navigate to the configs/ directory.
Create a file called "audio.json" within that directory.

Now you can configure audio files to be played, using the following JSON structure:

```
[
  { "name": "doorbell", "path": "file:///home/user/sounds/doorbell.mp3" },
  { "name": "chime", "path": "https://example.com/chime.mp3" }
]
```

The name will be used in the mqtt payload, from home assistant. <br>
The files themselves can be hosted remotely using http or retrieved from the host computer.

You can place audio clips in the gitignored directory called audio/
