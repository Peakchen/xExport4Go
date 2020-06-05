package Config


/* 
		export from NetFilter.json by tool.
*/
type	TNetfilterBase	 struct {
	Id	int32`json:"id"`
	Black	string`json:"black"`
	White	string`json:"white"`
}


type	TNetfilter	 struct {  data []* TNetfilterBase 
}

type	tArrNetfilter	[]*TNetfilterBase

var (
	GNetfilter	*TNetfilter=&TNetfilter{}
)

func init(){
	akLog.FmtPrintln("load	NetFilter.json")
}

func loadNetfilter() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "NetFilter.json") 
	Config.ParseJson2Cache(	GNetfilter,&tArrNetfilter{}, path)
}

func (this *TNetfilter) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrNetfilter) 
	for _, item := range *cfg {

	}
	return
}

func (this *TNetfilter) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrNetfilter) 
	this.data = []*TNetfilterBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TNetfilter) Get(idx int) *TNetfilterBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}