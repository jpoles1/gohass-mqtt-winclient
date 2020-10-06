# gohass-mqtt-winclient

_Note, this repo only works on Windows at this time._

## Installation:

1) Clone this repo
2) Run `install-in-startup.bat` in order to copy the executable and gohass-mqtt-winclient.env file to your startup folder.
3) Edit gohass-mqtt-winclient.env to include your MQTT server's URI, username, and password
4) Start the client, or restart your computer and it should start on its own!
5) Add your computer as a switch to your MQTT `configuration.yaml` file:
```
  - platform: mqtt
    name: Computer
    command_topic: "computer/power"
    optimistic: false
```
6) Add a button to your dashboard (optional):
```
type: button
tap_action:
  action: call-service
  service: mqtt.publish
  service_data:
    payload: "OFF"
    topic: computer/power
entity: switch.computer
```