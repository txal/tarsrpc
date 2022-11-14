package servant

import (
	"errors"
	"flag"
	"net/http"
	"os"
	"path"
	"time"

	"code.com/tars/goframework/tars/servant/warn"

	"github.com/go-redis/redis"

	"code.com/tars/goframework/jce/servant/taf"
	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
	"code.com/tars/goframework/kissgo/httpmetrics"
	"code.com/tars/goframework/tars/tarsserver"
	"code.com/tars/goframework/tars/util/conf"
)

const (
	// Turn this off if you do not want the framework calls flag.Parse() during init.
	useflag = true
)

var ConfigPath = ""
var (
	tarsAndPbSvrs map[string]*tarsserver.TarsServer     = make(map[string]*tarsserver.TarsServer)
	httpSvrs      map[string]*http.Server               = make(map[string]*http.Server)
	servantConfig map[string]*tarsserver.TarsServerConf = make(map[string]*tarsserver.TarsServerConf)
	shutdown      chan bool                             = make(chan bool, 1)

	objRunList []string

	configFile string
)

func initConfig() {
	if useflag {
		_configFile := (flag.String("config", ConfigPath, "init config path"))
		if os.Getenv("GOOS") != "darwin" && os.Getenv("test") != "true" {
			flag.Parse()
		}
		configFile = *_configFile
	} else {
		// Let the framework reads a pre-defined config file.
		configFile = "tars-config.conf"
	}
	if len(configFile) == 0 {
		appzaplog.SetLogLevel("debug")
		return
	}
	appzaplog.Debug("configFile parsed", zap.String("configFile", configFile))
	c := conf.NewConf(configFile)
	// load config
	if err := loadConfig(c); err != nil {
		appzaplog.Error("LoadConfig err", zap.Error(err))
		os.Exit(-1)
	}
}

func init() {
	appzaplog.InitAppLog()
	initConfig()
	appzaplog.Sync()
	srvConfig := GetServerConfig()
	if srvConfig != nil {
		logPath := srvConfig.LogPath + "/" + srvConfig.App + "/" + srvConfig.Server + "/" + srvConfig.Server + ".log"
		appzaplog.InitAppLog(appzaplog.ProcessName(path.Base(os.Args[0])+"_rd"), appzaplog.LogPath(logPath))
	}
	comm := startFrameWorkComm()
	if err := initFrameWorkClients(comm); err != nil {
		appzaplog.Error("initFrameWork failed", zap.Error(err))
	}

	if svrCfg != nil && svrCfg.App != "" && svrCfg.Server != "" && svrCfg.metricsenable == "yes" {
		httpmetrics.EnableMetrics(svrCfg.App, svrCfg.Server)
	}
}

func initFrameWorkClients(c *Communicator) error {
	//go initStatF(c)
	if cc := GetClientConfig(); cc != nil {
		if err := initStatF(c, cc.stat); err != nil {
			appzaplog.Error("initStatF failed", zap.Error(err))
		}

		if err := initRedis(cc.redisHost, cc.redisPassword, cc.redisDb); err != nil {
			appzaplog.Error("initRedis failed", zap.Error(err))
		}

		if err := initWechat(cc.wechatUrl); err != nil {
			appzaplog.Error("initWechat failed", zap.Error(err))
		}
	}

	if sc := GetServerConfig(); sc != nil {
		// todo should we return here
		if err := initTarConfig(c, sc, 15); err != nil {
			appzaplog.Error("initTarConfig failed", zap.Error(err))
		}

		if err := initNotify(c, sc); err != nil {
			appzaplog.Error("initNotify failed", zap.Error(err))
		}
	}

	return nil
}

func initRedis(host, password string, db int) error {
	if host == "" {
		return errors.New("redis host is empty")
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
	return nil
}

func initWechat(url string) error {
	return warn.InitWechat(url)
}

// addadminservant 添加admin servant
func addAdminServant() error {
	adf := new(taf.AdminF)
	ad := new(Admin)
	return addServant(adf, ad, fulladminobjname)
}

var RedisClient *redis.Client

func Run() {
	appzaplog.Info("Tars Version 1.3.3")
	// add adminF
	if err := addAdminServant(); err != nil {
		appzaplog.Error("addAdminServant failed", zap.Error(err))
		return
	}

	for _, obj := range objRunList {
		if s, ok := httpSvrs[obj]; ok {
			go func(obj string) {
				appzaplog.Info("http server start")
				err := s.ListenAndServe()
				if err != nil {
					appzaplog.Error("server start failed", zap.String("obj", obj), zap.Error(err))
					os.Exit(1)
				}
			}(obj)
			continue
		}

		s := tarsAndPbSvrs[obj]
		if s == nil {
			appzaplog.Debug("Obj not found", zap.String("obj", obj))
			break
		}
		appzaplog.Debug("Run", zap.String("obj", obj))
		go func(obj string) {
			err := s.Serve()
			if err != nil {
				appzaplog.Error("server start failed", zap.String("obj", obj), zap.Error(err))
				os.Exit(1)
			}
		}(obj)
		if obj != "AdminObj" && RedisClient != nil {
			ipPort := s.GetConfig().Address
			if err := RedisClient.Publish(obj, ipPort+"#start").Err(); err != nil {
				appzaplog.Error("pub start fail", zap.Error(err),
					zap.String("key", "gracefulReboot"),
					zap.String("ipPort", ipPort),
					zap.String("obj", obj))
			} else {
				appzaplog.Info("pub start",
					zap.String("key", "gracefulReboot"),
					zap.String("ipPort", ipPort),
					zap.String("obj", obj))
			}
		}
	}
	go reportNotifyInfo("restart")
	mainloop()
}

func mainloop() {
	ha := new(NodeFHelper)
	node := GetServerConfig().Node
	app := GetServerConfig().App
	server := GetServerConfig().Server
	container := GetServerConfig().Container
	ha.SetNodeInfo(startFrameWorkComm(), node, app, server, container)

	go ha.ReportVersion(GetServerConfig().Version)
	go ha.KeepAlive("") //first start
	loop := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-shutdown:
			appzaplog.Debug("tars shutdown")
			reportNotifyInfo("stop")
			return
		case <-loop.C:
			for name, adapter := range svrCfg.Adapters {
				if adapter.Protocol == "not_tars" {
					//TODO not_taf support
					ha.KeepAlive(name)
					continue
				}
				if s, ok := tarsAndPbSvrs[adapter.Obj]; ok {
					if !s.IsZombie(time.Second * 10) {
						ha.KeepAlive(name)
					}
				}
				if _, ok := httpSvrs[adapter.Obj]; ok {
					ha.KeepAlive(name)
					continue
				}
			}
		}
	}
}
