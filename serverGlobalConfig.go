package Config


/* 
		export from serverGlobalConfig.json by tool.
*/
type	TServerglobalconfigBase	 struct {
	Id	int32`json:"id"`
	Value	string`json:"Value"`
}


type	TServerglobalconfig	 struct {  data []* TServerglobalconfigBase 
}

type	tArrServerglobalconfig	[]*TServerglobalconfigBase

var (
	GServerglobalconfig	*TServerglobalconfig=&TServerglobalconfig{}
)

func init(){
	akLog.FmtPrintln("load	serverGlobalConfig.json")
}

func loadServerglobalconfig() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "serverGlobalConfig.json") 
	Config.ParseJson2Cache(	GServerglobalconfig,&tArrServerglobalconfig{}, path)
}

func (this *TServerglobalconfig) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrServerglobalconfig) 
	for _, item := range *cfg {

	}
	return
}

func (this *TServerglobalconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrServerglobalconfig) 
	this.data = []*TServerglobalconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TServerglobalconfig) Get(idx int) *TServerglobalconfigBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}