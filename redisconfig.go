package Config


/* 
		export from redisconfig.json by tool.
*/
type	TRedisconfigBase	 struct {
	Id	int32`json:"id"`
	Dbindex	int32`json:"DBIndex"`
	Connaddr	string`json:"ConnAddr"`
	Shareconnaddr	string`json:"ShareConnAddr"`
	Passwd	string`json:"Passwd"`
	Pprofaddr	string`json:"PProfAddr"`
	Sharedbindex	int32`json:"ShareDBIndex"`
}


type	TRedisconfig	 struct {  data []* TRedisconfigBase 
}

type	tArrRedisconfig	[]*TRedisconfigBase

var (
	GRedisconfig	*TRedisconfig=&TRedisconfig{}
)

func init(){
	akLog.FmtPrintln("load	redisconfig.json")
}

func loadRedisconfig() {
	var(
		path string
	)
	if len(SvrPath) == 0 {
		path = getserverpath()
	} 
	path = filepath.Join(SvrPath, "redisconfig.json") 
	Config.ParseJson2Cache(	GRedisconfig,&tArrRedisconfig{}, path)
}

func (this *TRedisconfig) ComfireAct(data interface{}) (errlist []string) { errlist = []string{} 
	cfg := data.(*tArrRedisconfig) 
	for _, item := range *cfg {

	}
	return
}

func (this *TRedisconfig) DataRWAct(data interface{}) (errlist []string) {
	cfg := data.(*tArrRedisconfig) 
	this.data = []*TRedisconfigBase{}
	for _, item := range *cfg {

	}
	return
}

func (this *TRedisconfig) Get(idx int) *TRedisconfigBase{
	if idx >= len(this.data) {
		return nil
	}
	return this.data[idx]
}