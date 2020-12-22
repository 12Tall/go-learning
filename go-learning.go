package go-learning

// 模块初始化函数
func init() {
	log.SetPrefix("[INFO]")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("log")
}