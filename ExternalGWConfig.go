package Config


/* 
		export from ExternalGWConfig.json by tool.
*/
type	TExternalgwconfigBase	 struct {
	Id	int32`json:"id"`
	Pprofaddr	string`json:"PProfAddr"`
	Listenaddr	string`json:"ListenAddr"`
}


type	TExternalgwconfig	 struct {  data []* TExternalgwconfigBase 
}

type	tArrExternalgwconfig	[]*TExternalgwconfigBase

var (
	GExternalgwconfig	*TExternalgwconfig=&TExternalgwconfig{}
)

func init(){
	akLog.FmtPrintln("load	ExternalGWConfig.json")
}

func loadExternalgwconfig() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "ExternalGWConfig.json") 
	Config.ParseJson2Cache(	GExternalgwconfig,&tArrExternalgwconfig{}, path)
}

func (this *TExternalgwconfig) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrExternalgwconfig) 
	for _, item := range *cfg {

	}
	return
}

func (this *TExternalgwconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrExternalgwconfig) 
	this.data = []*TExternalgwconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TExternalgwconfig) Get(idx int) *TExternalgwconfigBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}