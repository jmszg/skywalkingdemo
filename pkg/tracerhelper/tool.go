package tracerhelper

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/SkyAPM/go2sky"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

// 发送get请求方法，并创建一个exitspan
func Get(link string) (response string, err error) {
	// 创建一个client对象
	client := http.Client{Timeout: time.Second * 10}

	// 创建一个http get请求对象
	var reqest *http.Request
	reqest, err = http.NewRequest("GET", link, nil)
	if err != nil {
		return
	}

	// 获取当前进程的tracer
	tracer = GetTracer()
	// perr 的作用就是Tag信息
	// operationName 就是调用名称，意思要明确
	url := reqest.URL
	operationName := url.Scheme + "://" + url.Host + url.Path
	ctx, _ := GetGcm().GetContext()
	// 创建一个exitspan
	span, err := tracer.CreateExitSpan(*ctx, operationName, url.Host, func(key, value string) error {
		reqest.Header.Set(key, value)
		return nil
	})
	if err != nil {
		return
	}
	span.SetComponent(ComponentIDGINHttpServer)
	span.Tag(go2sky.TagHTTPMethod, reqest.Method)
	span.Tag(go2sky.TagURL, link)
	span.SetSpanLayer(agentv3.SpanLayer_Http)

	// client发送请求
	resp, err := client.Do(reqest)
	if err != nil {
		// 标记span的is_error状态
		span.Error(time.Now(), err.Error())
	} else {
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(resp.StatusCode))
	}

	// 结束跟踪
	span.End()

	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
