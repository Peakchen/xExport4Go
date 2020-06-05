package Config


/* 
		export from InnerGWConfig.json by tool.
*/
type	TInnergwconfigBase	 struct {
	Id	int32`json:"id"`
	Connectaddr	string`json:"ConnectAddr"`
	Listenaddr	string`json:"ListenAddr"`
	Zone	string`json:"Zone"`
	No	int32`json:"No"`
	Pprofaddr	string`json:"PProfAddr"`
}


type	TInnergwconfig	 struct {  data []* TInnergwconfigBase 
}

type	tArrInnergwconfig	[]*TInnergwconfigBase

var (
	GInnergwconfig	*TInnergwconfig=&TInnergwconfig{}
)

func init(){
	akLog.FmtPrintln("load	InnerGWConfig.json")
}

func loadInnergwconfig() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "InnerGWConfig.json") 
	Config.ParseJson2Cache(	GInnergwconfig,&tArrInnergwconfig{}, path)
}

func (this *TInnergwconfig) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrInnergwconfig) 
	for _, item := range *cfg {

	}
	return
}

func (this *TInnergwconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrInnergwconfig) 
	this.data = []*TInnergwconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TInnergwconfig) Get(idx int) *TInnergwconfigBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}