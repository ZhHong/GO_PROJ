package utils

import "time"

func LogInfo(values []string){
	writeLog(values,"info")
}

func LogError(values []string){
	writeLog(values,"error")
}

func LogWarning(values []string){
	writeLog(values,"warning")
}

func writeLog(values []string ,log_type string) error{
	// LOG_2016_06_17.ERROR
	// 2016-06-17 :: log_string
	//timestr := time.Now()
	//outfile = "."
	//file,err := os.Create(outfile)
	//1 find log file if exists
	//2 if not exists create file
	//3 write append log
	return nil
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