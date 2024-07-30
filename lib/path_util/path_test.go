package path_util

import (
	"shop/lib/path_util/logger"
	"testing"
)

func TestEb5gcPath(t *testing.T) {
	logger.PathLog.Infoln(Goeb5gcPath("shop/conf/app.ini"))
}
