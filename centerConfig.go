package Config

import "path/filepath"

/*
	export from centerConfig.json by tool.
*/
type TCenterconfigBase struct {
	Id         int32  `json:"id"`
	No         int32  `json:"No"`
	Listenaddr string `json:"ListenAddr"`
	Pprofaddr  string `json:"PProfAddr"`
	Zone       string `json:"Zone"`
}

type TCenterconfig struct {
	data []*TCenterconfigBase
}

type tArrCenterconfig []*TCenterconfigBase

var (
	GCenterconfig *TCenterconfig = &TCenterconfig{}
)

func init() {
	akLog.FmtPrintln("load	centerConfig.json")
}

func loadCenterconfig() {
	var (
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	}
	path = filepath.Join(SvrPath, "centerConfig.json")
	Config.ParseJson2Cache(GCenterconfig, &tArrCenterconfig{}, path)
}

func (this *TCenterconfig) ComfireAct(data interface{}) (errlist []string) {
	errlist = []string{}
	cfg := data.(*tArrCenterconfig)
	for _, item := range *cfg {

	}
	return
}

func (this *TCenterconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrCenterconfig)
	this.data = []*TCenterconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TCenterconfig) Get(idx int) *TCenterconfigBase {
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}
