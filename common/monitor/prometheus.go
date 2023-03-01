package monitor

import (
	"github.com/LCY2013/paas/common/async"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())

	// 启动web服务
	async.GO(func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			logrus.WithField("prometheus", "启动失败").Fatal(err)
		}

		logrus.Info("监控启动，端口为：" + strconv.Itoa(port))
	})
}
