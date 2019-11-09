package Config


/* 
		export from animal.json by tool.
*/
type	TAnimal	 struct {
	Id	int32`json:"id"`
	Xnq	string`json:"xnq"`
	Cat	string`json:"cat"`
	Chuken	int32`json:"chuken"`
	Pig	string`json:"pig"`
}


type	TAnimalConfig	 struct {
}

func (this *TAnimalConfig) ComfireAct(data interface{}) (errlist []string) {

	return
}

func (this *TAnimalConfig) DataRWAct(data interface{}) (errlist []string) {

	return
}

