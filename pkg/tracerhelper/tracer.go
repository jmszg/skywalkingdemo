package tracerhelper

import (
	"skywalkingdemo/pkg/tracerhelper/util"
	"sync"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
)

var tracer *go2sky.Tracer
var gcm util.GoroutineContextManager
var once sync.Once

// 开始链路跟踪，创建一个trace
func StartTracer(serviceAddr string, serviceName string) error {
	// 创建grpc reporter用于给oap发送trace数据
	rp, err := reporter.NewGRPCReporter(serviceAddr, reporter.WithCheckInterval(time.Second))
	if err != nil {
		return err
	}
	// 创建一个tracer
	tracer, err = go2sky.NewTracer(serviceName, go2sky.WithReporter(rp))

	// 初始化localstore，只初始化一次
	once.Do(func() {
		gcm = util.GoroutineContextManager{}
	})

	return nil
}

func GetTracer() *go2sky.Tracer {
	return tracer
}

func GetGcm() *util.GoroutineContextManager {
	return &gcm
}
