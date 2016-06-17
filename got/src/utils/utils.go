package utils

import "time"

func LogInfo(){

}

func LogError(){

}

func LogWarning(){

}

func writeLog(values []string ,log_type string) error{
	// LOG_2016_06_17.ERROR
	// 2016-06-17 :: log_string
	//timestr := time.Now()
	//outfile = "."
	//file,err := os.Create(outfile)
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