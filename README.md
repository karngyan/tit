![make me a new titter](./titui/static/icon.png)


# tit

tit daemon

- write tests
- automate build & install

## plist

- replace `{{path_for_titd_executable}}` accordingly.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist SYSTEM "file://localhost/System/Library/DTDs/PropertyList.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>titd</string>
		<key>ServiceDescription</key>
		<string>Twitter local scheduler tit daemon</string>
		<key>ProgramArguments</key>
		<array>
			<string>{{path_for_titd_executable}}</string>
		</array>
		<key>RunAtLoad</key>
		<true/>
	</dict>
</plist>

```