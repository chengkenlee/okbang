// @program:     gatebitcoin
// @file:        run.go
// @author:      chengkenlee
// @create:      2023-01-05 21:50
// @description:
package util

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const ENCKEY = "**************"

var (
	Config *viper.Viper
	Logger *zap.Logger
)
