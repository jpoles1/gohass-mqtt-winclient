copy /y "dist\gohass-mqtt-winclient.exe" "%userprofile%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\"
echo n | copy /-y "example.env" "%userprofile%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\gohass-mqtt-winclient.env"
explorer.exe "%userprofile%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup"
PAUSE