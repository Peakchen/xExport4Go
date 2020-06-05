package Config


/* 
		export from gameConfig.json by tool.
*/
type	TGameconfigBase	 struct {
	Id	int32`json:"id"`
	No	int32`json:"No"`
	Listenaddr	string`json:"ListenAddr"`
	Pprofaddr	string`json:"PProfAddr"`
	Zone	string`json:"Zone"`
}


type	TGameconfig	 struct {  data []* TGameconfigBase 
}

type	tArrGameconfig	[]*TGameconfigBase

var (
	GGameconfig	*TGameconfig=&TGameconfig{}
)

func init(){
	akLog.FmtPrintln("load	gameConfig.json")
}

func loadGameconfig() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "gameConfig.json") 
	Config.ParseJson2Cache(	GGameconfig,&tArrGameconfig{}, path)
}

func (this *TGameconfig) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrGameconfig) 
	for _, item := range *cfg {

	}
	return
}

func (this *TGameconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrGameconfig) 
	this.data = []*TGameconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TGameconfig) Get(idx int) *TGameconfigBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}