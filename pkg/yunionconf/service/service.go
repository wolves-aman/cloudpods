package service

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"yunion.io/x/log"

	"yunion.io/x/onecloud/pkg/cloudcommon"
	"yunion.io/x/onecloud/pkg/cloudcommon/consts"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/yunionconf"
	"yunion.io/x/onecloud/pkg/yunionconf/models"
	"yunion.io/x/onecloud/pkg/yunionconf/options"
)

func StartService() {

	opts := &options.Options
	commonOpts := &options.Options.CommonOptions
	dbOpts := &options.Options.DBOptions
	cloudcommon.ParseOptions(opts, os.Args, "yunionconf.conf", "yunionconf")
	cloudcommon.InitAuth(commonOpts, func() {
		log.Infof("Auth complete!!")
	})

	if opts.GlobalVirtualResourceNamespace {
		consts.EnableGlobalVirtualResourceNamespace()
	}

	cloudcommon.InitDB(dbOpts)

	app := cloudcommon.InitApp(commonOpts, true)
	yunionconf.InitHandlers(app)

	if db.CheckSync(opts.AutoSyncTable) {
		err := models.InitDB()
		if err == nil {
			cloudcommon.ServeForeverWithCleanup(app, commonOpts, func() {
				cloudcommon.CloseDB()
			})
		} else {
			log.Errorf("InitDB fail: %s", err)
		}
	}
}
