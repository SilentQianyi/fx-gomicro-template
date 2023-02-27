package enum

type (
	HelloWorldStatus int32
)

const (
	HelloWorldStatusNone    HelloWorldStatus = 0
	HelloWorldStatusDoing   HelloWorldStatus = 1
	HelloWorldStatusDone    HelloWorldStatus = 2
	HelloWorldStatusUnknown HelloWorldStatus = 100
)
