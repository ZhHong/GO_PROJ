package utils

import "time"
import "os"
import "fmt"
import "io"

func LogInfo(values string){
	fmt.Println("LOG_INFO_"+Now()+":" + values)
	writeLog(values,"info")
}

func LogError(values string){
	writeLog(values,"error")
}

func LogWarning(values string){
	writeLog(values,"warning")
}

func LogSys(values string){

}

func writeLog(values string ,log_type string) error{
	// LOG_2016_06_17.ERROR
	// 2016-06-17 :: log_string
	//timestr := time.Now()
	//outfile = "."
	//file,err := os.Create(outfile)
	//1 find log file if exists
	//2 if not exists create file
	//3 write append log
	filepath := "../logs/"
	infile := filepath + Today()+ "_"+log_type +".log"
	fileExist := checkFileIsExist(infile)

	var file *os.File
	var err error

	if fileExist{
		file,err = os.OpenFile(infile,os.O_APPEND,0666)
		fmt.Println("file exsits")
	}else{
		file,err = os.Create(infile)
		fmt.Println("file not exists")
	}
	checkError(err)
	n,err :=io.WriteString(file,values+"\n")
	fmt.Println("write len ====%d",n)
	checkError(err)
	file.Sync()
	defer file.Close()
	return nil
}

func checkFileIsExist(filename string)(bool){
	var exist = true
	if _,err := os.Stat(filename);os.IsNotExist(err){
		exist = false
	}
	return exist;
}

func checkError(e error){
	if e != nil {
		panic(e)
	}
}

func Today() string{
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp,0)
	//must be 2006-01-02
	return tm.Format("2006-01-02")
}

func Now() string{
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp,0)
	//must be 2006-01-02 15:04:05
	return tm.Format("2006-01-02 15:04:05")
}