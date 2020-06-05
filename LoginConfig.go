package Config


/* 
		export from LoginConfig.json by tool.
*/
type	TLoginconfigBase	 struct {
	Id	int32`json:"id"`
	No	int32`json:"No"`
	Listenaddr	string`json:"ListenAddr"`
	Pprofaddr	string`json:"PProfAddr"`
	Zone	string`json:"Zone"`
}


type	TLoginconfig	 struct {  data []* TLoginconfigBase 
}

type	tArrLoginconfig	[]*TLoginconfigBase

var (
	GLoginconfig	*TLoginconfig=&TLoginconfig{}
)

func init(){
	akLog.FmtPrintln("load	LoginConfig.json")
}

func loadLoginconfig() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "LoginConfig.json") 
	Config.ParseJson2Cache(	GLoginconfig,&tArrLoginconfig{}, path)
}

func (this *TLoginconfig) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrLoginconfig) 
	for _, item := range *cfg {

	}
	return
}

func (this *TLoginconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrLoginconfig) 
	this.data = []*TLoginconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TLoginconfig) Get(idx int) *TLoginconfigBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}