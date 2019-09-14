package hosts


func Open() {

}

//func getHostFilePath() string {
//
//}

func AppendHostRecord(records []HostRecord) {
	//TODO 获取host文件
	hostsFilePath := GetHostsPath()


}

type HostRecord struct {
	Ip string
	Domain []string
}


