// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package main

import (
	"github.com/sirupsen/logrus"
	"hannesdw/timesheet/server/cmd"
	"hannesdw/timesheet/server/logger"
)

func main() {
	logger.Log = logrus.New()
	cmd.Execute()
}
