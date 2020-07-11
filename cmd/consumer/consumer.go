package main

var (
	brokers = []string{"10.115.5.127:9092", "10.115.5.128:9092", "10.115.5.129:9092"}
	topic   = "android_crash_topic"
	groupId = "android_crash_topic"
)

func main() {
	//dialer := &kafka.Dialer{
	//	Timeout:   2 * time.Second,
	//	DualStack: true,
	//	SASLMechanism: plain.Mechanism{
	//		Username: "testceshi",
	//		Password: "frr6lz2ksb09qx59ghorycq6cq82rm",
	//	},
	//}
	//for i := 0; i < 10; i++ {
	//	if err := producer(brokers, topic, dialer, []byte{}, []byte(str)); err != nil {
	//		fmt.Printf("err:%+v", err)
	//	}
	//}

	//consumer(brokers, topic, groupId, dialer)
}