# Rancher Logger

This is a formatter implementation to the popular logrus logging library for golang

This formatter adds the caller and optionally also adds the stack trace to log messages printed by the logrus logger.

Here's a simple program showing its usage and output

```
package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	rl "github.com/wlan0/rancher-logger"
)

func main() {
	log.Info("Here's an Info message")
	log.Debug("Here's a Debug message")
	log.Error("Here's a Error message")

	textFormatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(textFormatter)

	log.Info("Here's an Info message")
	log.Debug("Here's a Debug message")
	log.Error("Here's a Error message")

	textFormatter2 := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.StampNano,
	}
	log.SetFormatter(textFormatter2)

	log.Info("Here's an Info message")
	log.Debug("Here's a Debug message")
	log.Error("Here's a Error message")

	rancherFormatter := &rl.RancherFormatter{
		textFormatter,
		false,
	}
	log.SetFormatter(rancherFormatter)

	log.Info("Here's an Info message")
	log.Debug("Here's a Debug message")
	log.Error("Here's a Error message")
}
```

output

```
INFO[0000] Here's an Info message
ERRO[0000] Here's a Error message
INFO[2015-10-08T11:49:55-07:00] Here's an Info message
ERRO[2015-10-08T11:49:55-07:00] Here's a Error message
INFO[Oct  8 11:49:55.406427849] Here's an Info message
ERRO[Oct  8 11:49:55.406432917] Here's a Error message
INFO[2015-10-08T11:49:55-07:00] Here's an Info message                        caller=/path/to/repo/github.com/wlan0/testx/main.go:40
ERRO[2015-10-08T11:49:55-07:00] Here's a Error message                        caller=/path/to/repo/github.com/wlan0/testx/main.go:42
```
