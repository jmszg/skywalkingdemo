package ginagent

import (
	tracerhelper2 "skywalkingdemo/pkg/tracerhelper"
	"strconv"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/gin-gonic/gin"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

// 请求拦截中间件
func Middleware() gin.HandlerFunc {
	tracerobj := tracerhelper2.GetTracer()
	if tracerobj == nil {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		// 在api提供服务的接口中创建entryspan
		operationName := c.FullPath()
		span, ctx, err := tracerobj.CreateEntrySpan(c.Request.Context(), operationName, func(key string) (string, error) {
			return c.Request.Header.Get(key), nil
		})
		if err != nil {
			c.Next()
			return
		}
		// 组件id对应名称：https://github.com/apache/skywalking/blob/master/oap-server/server-starter/src/main/resources/component-libraries.yml
		span.SetComponent(tracerhelper2.ComponentIDGINHttpServer)
		span.Tag(go2sky.TagHTTPMethod, c.Request.Method)
		span.Tag(go2sky.TagURL, c.Request.Host+c.Request.URL.Path)
		span.SetSpanLayer(agentv3.SpanLayer_Http)

		// test log
		span.Log(time.Now(), "test log info")

		// test error 暂时注释
		//span.Error(time.Now(), "test log error")

		//c.Request = c.Request.WithContext(ctx)
		tracerhelper2.GetGcm().SetContext(&ctx)
		defer tracerhelper2.GetGcm().DelContext()

		c.Next()

		if len(c.Errors) > 0 {
			span.Error(time.Now(), c.Errors.String())
		}
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(c.Writer.Status()))
		span.End()
	}
}
