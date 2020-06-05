package Config


/* 
		export from mgoconfig.json by tool.
*/
type	TMgoconfigBase	 struct {
	Pprofaddr	string`json:"PProfAddr"`
	Id	int32`json:"id"`
	Passwd	string`json:"Passwd"`
	Username	string`json:"UserName"`
	Shareusername	string`json:"ShareUserName"`
	Host	string`json:"Host"`
	Sharehost	string`json:"ShareHost"`
	Sharepasswd	string`json:"SharePasswd"`
}


type	TMgoconfig	 struct {  data []* TMgoconfigBase 
}

type	tArrMgoconfig	[]*TMgoconfigBase

var (
	GMgoconfig	*TMgoconfig=&TMgoconfig{}
)

func init(){
	akLog.FmtPrintln("load	mgoconfig.json")
}

func loadMgoconfig() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "mgoconfig.json") 
	Config.ParseJson2Cache(	GMgoconfig,&tArrMgoconfig{}, path)
}

func (this *TMgoconfig) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrMgoconfig) 
	for _, item := range *cfg {

	}
	return
}

func (this *TMgoconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrMgoconfig) 
	this.data = []*TMgoconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TMgoconfig) Get(idx int) *TMgoconfigBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}