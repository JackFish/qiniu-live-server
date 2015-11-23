package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/qiniu/log"
	"live/config"
)

func InitOrm(cfg *config.OrmConfig) (err error) {
	log.Info("init orm")
	regErr := orm.RegisterDataBase("default", cfg.DriverName, cfg.DataSource,
		cfg.MaxIdleConn, cfg.MaxOpenConn)
	if regErr != nil {
		err = regErr
		return
	}

	orm.Debug = cfg.DebugMode

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Session))
	orm.RegisterModel(new(LiveStream))
	orm.RegisterModel(new(LiveVideo))

	return
}
