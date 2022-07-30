# gowaybarplug

Go framework for custom [Waybar](https://github.com/Alexays/Waybar) plugins

# Usage

First build a plugin that reports some interesting status:

```go
package main

import (
    "time"

    waybar "github.com/lack/gowaybarplug"
)

main() {
    updater := waybar.NewUpdater()
    for {
        status := waybar.Status{
            Text: "Some text",
            Toolbar: "Other text",
        }
        // Obviously do something more interesting than just static text in the status...
        updater.Status <- &status
        time.Sleep(15 * time.Second)
    }
}
```

Then add it to your ~/.config/waybar/config:

```json
{
    // ... Other waybar config
    "custom/mything": {
        "format": "{} {icon}",
        "return-type": "json",
        "exec": "/path/to/my/new/plugin"
        // etc
    }
}
```
