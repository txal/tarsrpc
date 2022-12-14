/*

 */
package servant

import (
	"errors"
	"code.com/tars/goframework/kissgo/appzaplog/zap"
	"net/http"
	"code.com/tars/goframework/kissgo/appzaplog"
	"code.com/tars/goframework/tars/tarsserver"
)

var (
	ServantConfNotFoundTErr = errors.New("server config not found")
	ServantExist = errors.New("goSvrs already exist")
)

func GetServantConfig(objname string) (cfg *tarsserver.TarsServerConf, err error) {
	var fullobjname string
	fullobjname, err = fullObjName(objname)
	if err != nil {
		return
	}
	var ok bool
	cfg, ok = servantConfig[fullobjname]
	if !ok {
		err = errors.New("not found")
		return
	}
	return
}

// AddServant only used for tars header,idl plugin use tarsheaderpbbody2go
func AddServant(v Dispatcher, f interface{}, objname string) error{
	fullobjname,err := fullObjName(objname)
	if err != nil {
		return err
	}
	return addServant(v,f,fullobjname)
}

func addServant(v Dispatcher, f interface{}, fullobjname string) error {
	objRunList = append(objRunList, fullobjname)
	cfg, ok := servantConfig[fullobjname]
	if !ok {
		appzaplog.Debug("AddServant servant obj name not found ", zap.String("fullobjname", fullobjname))
		return ServantConfNotFoundTErr
	}
	appzaplog.Debug("AddServant add tarsConfig", zap.Any("cfg", cfg))
	jp := NewJceProtocol(v, f)
	s := tarsserver.NewTarsServer(jp, cfg)
	tarsAndPbSvrs[fullobjname] = s
	return nil
}

// AddPbServant full pb support, pb header + pb body, idl plugin use tars2go
//func AddPbServant(v PbDispatcher, f interface{}, objname string) error {
//	fullobjname,err := fullObjName(objname)
//	if err != nil {
//		return err
//	}
//
//	if _,exist := tarsAndPbSvrs[fullobjname];exist{
//		return ServantExist
//	}
//
//	cfg, ok := servantConfig[fullobjname]
//	if !ok {
//		appzaplog.Warn("AddPbServant servant obj name not found ", zap.String("fullobjname", fullobjname))
//		return ServantConfNotFoundTErr
//	}
//	appzaplog.Debug("AddPbServant add tarsConfig", zap.Any("cfg", cfg))
//
//	objRunList = append(objRunList, fullobjname)
//	jp := NewPbProtocol(v, f)
//	s := tarsserver.NewTarsServer(jp, cfg)
//	tarsAndPbSvrs[fullobjname] = s
//	return nil
//}

// AddHttpServant ????????????http???server
func AddHttpServant(mux http.Handler, objname string) error{
	fullobjname,err := fullObjName(objname)
	if err != nil {
		return err
	}
	cfg, ok := servantConfig[fullobjname]
	if !ok {
		appzaplog.Debug("servant obj name not found ", zap.String("fullobjname", fullobjname))
		return ServantConfNotFoundTErr
	}
	appzaplog.Debug("add http server", zap.String("fullobjname", fullobjname), zap.Any("cfg", cfg))
	objRunList = append(objRunList, fullobjname)
	s := &http.Server{Addr: cfg.Address, Handler: mux}
	httpSvrs[fullobjname] = s
	return nil
}