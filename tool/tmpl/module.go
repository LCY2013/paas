package tmpl

var (
	Module = `module {{.Dir}}

go 1.20

require (
	github.com/LCY2013/paas/common v0.0.0-20230301062216-94bcaddc2b39
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-micro/plugins/v4/registry/consul v1.2.0
	github.com/go-micro/plugins/v4/wrapper/select/roundrobin v1.2.0
	github.com/go-micro/plugins/v4/wrapper/trace/opentracing v1.2.0
	go-micro.dev/v4 v4.9.0
	github.com/jinzhu/gorm v1.9.16
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.28.1
	k8s.io/api v0.22.4 //其它版本会报错
	k8s.io/client-go v0.22.4 //其它版本会报错
	gorm.io/driver/mysql v1.4.7
)
`
	ApiModule = `module {{.Dir}}

go 1.20

require (
	github.com/LCY2013/paas/common v0.0.0-20230301061309-671b1e41245e
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-micro/plugins/v4/registry/consul v1.2.0
	github.com/go-micro/plugins/v4/wrapper/select/roundrobin v1.2.0
	github.com/go-micro/plugins/v4/wrapper/trace/opentracing v1.2.0
	go-micro.dev/v4 v4.9.0
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.28.1
	k8s.io/api v0.22.4 //其它版本会报错
	k8s.io/client-go v0.22.4 //其它版本会报错
)
`
)
