package gateway

import (
	"fmt"
	"github.com/LCY2013/paas/micro-api-gateway/internal/handler"
	"github.com/LCY2013/paas/micro-api-gateway/internal/helper"
	"github.com/LCY2013/paas/micro-api-gateway/internal/namespace"
	apiReg "github.com/LCY2013/paas/micro-api-gateway/internal/registry"
	rrmicro "github.com/LCY2013/paas/micro-api-gateway/internal/resolver/api"
	"github.com/LCY2013/paas/micro-api-gateway/stats"
	grpcClient "github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	ahandler "go-micro.dev/v4/api/handler"
	aapi "go-micro.dev/v4/api/handler/api"
	"go-micro.dev/v4/api/handler/event"
	ahttp "go-micro.dev/v4/api/handler/http"
	arpc "go-micro.dev/v4/api/handler/rpc"
	"go-micro.dev/v4/api/handler/web"
	"go-micro.dev/v4/api/resolver"
	"go-micro.dev/v4/api/resolver/grpc"
	"go-micro.dev/v4/api/resolver/host"
	"go-micro.dev/v4/api/resolver/path"
	"go-micro.dev/v4/api/router"
	regRouter "go-micro.dev/v4/api/router/registry"
	"go-micro.dev/v4/api/server"
	"go-micro.dev/v4/api/server/acme"
	"go-micro.dev/v4/api/server/cors"
	httpapi "go-micro.dev/v4/api/server/http"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"

	//"go-micro.dev/v4/auth"
	//plugin "go-micro.dev/v4/plugins"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"net/http"
	"strings"
)

var (
	Name                  = "go.micro.api"
	Address               = ":8080"
	Handler               = "meta"
	Resolver              = "micro"
	RPCPath               = "/rpc"
	APIPath               = "/"
	ProxyPath             = "/{service:[a-zA-Z0-9]+}"
	Namespace             = "go.micro"
	Type                  = "api"
	HeaderPrefix          = "X-Micro-"
	EnableRPC             = false
	ACMEProvider          = "autocert"
	ACMEChallengeProvider = "cloudflare"
	ACMECA                = acme.LetsEncryptProductionCA
	Registry              = "consul"
	RegistryAddress       = "127.0.0.1:8500"
)

func Run(ctx *cli.Context, srvOpts ...micro.Option) {
	if len(ctx.String("server_name")) > 0 {
		Name = ctx.String("server_name")
	}
	if len(ctx.String("address")) > 0 {
		Address = ctx.String("address")
	}
	if len(ctx.String("handler")) > 0 {
		Handler = ctx.String("handler")
	}
	if len(ctx.String("resolver")) > 0 {
		Resolver = ctx.String("resolver")
	}
	if len(ctx.String("enable_rpc")) > 0 {
		EnableRPC = ctx.Bool("enable_rpc")
	}
	if len(ctx.String("acme_provider")) > 0 {
		ACMEProvider = ctx.String("acme_provider")
	}
	if len(ctx.String("type")) > 0 {
		Type = ctx.String("type")
	}
	if len(ctx.String("namespace")) > 0 {
		// remove the service type from the namespace to allow for
		// backwards compatability
		Namespace = strings.TrimSuffix(ctx.String("namespace"), "."+Type)
	}

	// apiNamespace has the format: "go.micro.api"
	//??????????????????
	apiNamespace := Namespace + "." + Type

	// append name to opts
	//??????????????? srv ???????????????

	srvOpts = append(srvOpts, micro.Name(Name))

	if len(ctx.String("registry_address")) > 0 {
		RegistryAddress = ctx.String("registry_address")
	}

	if len(ctx.String("registry")) > 0 {
		Registry = ctx.String("registry")
	}

	var reOpt []registry.Option
	reOpt = append(reOpt, registry.Addrs(RegistryAddress))
	//registry := consul.NewRegistry(reOpt...)
	registry := apiReg.Register(Registry, reOpt...)

	srvOpts = append(srvOpts, micro.Registry(registry))
	//????????? registory

	// initialise service
	//???????????????
	service := micro.NewService(srvOpts...)

	// Init plugins
	//???????????????
	//for _, p := range Plugins() {
	//	p.Init(ctx)
	//}

	// Init API
	//?????????API
	var opts []server.Option

	//????????????
	if ctx.Bool("enable_acme") {
		//hosts := helper.ACMEHosts(ctx)
		//opts = append(opts, server.EnableACME(true))
		//opts = append(opts, server.ACMEHosts(hosts...))
		//switch ACMEProvider {
		//case "autocert":
		//	opts = append(opts, server.ACMEProvider(autocert.NewProvider()))
		//case "certmagic":
		//	if ACMEChallengeProvider != "cloudflare" {
		//		log.Fatal("The only implemented DNS challenge provider is cloudflare")
		//	}
		//
		//	apiToken := os.Getenv("CF_API_TOKEN")
		//	if len(apiToken) == 0 {
		//		log.Fatal("env variables CF_API_TOKEN and CF_ACCOUNT_ID must be set")
		//	}
		//
		//	storage := certmagic.NewStorage(
		//		memory.NewSync(),
		//		service.Options().Store,
		//	)
		//
		//	config := cloudflare.NewDefaultConfig()
		//	config.AuthToken = apiToken
		//	config.ZoneToken = apiToken
		//	challengeProvider, err := cloudflare.NewDNSProviderConfig(config)
		//	if err != nil {
		//		log.Fatal(err.Error())
		//	}
		//
		//	opts = append(opts,
		//		server.ACMEProvider(
		//			certmagic.NewProvider(
		//				acme.AcceptToS(true),
		//				acme.CA(ACMECA),
		//				acme.Cache(storage),
		//				acme.ChallengeProvider(challengeProvider),
		//				acme.OnDemand(false),
		//			),
		//		),
		//	)
		//default:
		//	log.Fatalf("%s is not a valid ACME provider\n", ACMEProvider)
		//}
	} else if ctx.Bool("enable_tls") {
		config, err := helper.TLSConfig(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		opts = append(opts, server.EnableTLS(true))
		opts = append(opts, server.TLSConfig(config))
	}

	if ctx.Bool("enable_cors") {
		opts = append(opts, server.EnableCORS(true))
		opts = append(opts, server.CORSConfig(&cors.Config{
			//AllowCredentials: true,
			//AllowOrigin:      "*",
			//AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
			//AllowHeaders:     "Accept, Content-Type, Content-Length",
		}))
	}

	// create the router
	//????????????
	var h http.Handler
	r := mux.NewRouter()
	h = r

	//??????????????????
	if ctx.Bool("enable_stats") {
		st := stats.New()
		r.HandleFunc("/stats", st.StatsHandler)
		h = st.ServeHTTP(r)
		st.Start()
		defer st.Stop()
	}

	// return version and list of services
	//?????????????????????
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}
		response := fmt.Sprintf(`{"version": "%s"}`, ctx.App.Version)
		w.Write([]byte(response))
	})

	// strip favicon.ico
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	// register rpc handler
	//?????? rpc
	if EnableRPC {
		log.Infof("Registering RPC Handler at %s", RPCPath)
		r.HandleFunc(RPCPath, handler.RPC)
	}

	// create the namespace resolver
	nsResolver := namespace.NewResolver(Type, Namespace)

	// resolver options
	ropts := []resolver.Option{
		resolver.WithNamespace(nsResolver.ResolveWithType),
		resolver.WithHandler(Handler),
	}

	// default resolver

	rr := rrmicro.NewResolver(ropts...)

	switch Resolver {
	case "host":
		rr = host.NewResolver(ropts...)
	case "path":
		rr = path.NewResolver(ropts...)
	case "grpc":
		service.Init(func(options *micro.Options) {
			options.Client = grpcClient.NewClient()
		})
		rr = grpc.NewResolver(ropts...)
	}

	switch Handler {
	case "rpc":
		log.Infof("Registering API RPC Handler at %s", APIPath)
		rt := regRouter.NewRouter(
			router.WithHandler(arpc.Handler),
			router.WithResolver(rr),
			router.WithRegistry(service.Options().Registry),
		)
		rp := arpc.NewHandler(
			ahandler.WithNamespace(apiNamespace),
			ahandler.WithRouter(rt),
			ahandler.WithClient(service.Client()),
		)
		r.PathPrefix(APIPath).Handler(rp)
	case "api":
		log.Infof("Registering API Request Handler at %s", APIPath)
		rt := regRouter.NewRouter(
			router.WithHandler(aapi.Handler),
			router.WithResolver(rr),
			router.WithRegistry(service.Options().Registry),
		)
		ap := aapi.NewHandler(
			ahandler.WithNamespace(apiNamespace),
			ahandler.WithRouter(rt),
			ahandler.WithClient(service.Client()),
		)
		r.PathPrefix(APIPath).Handler(ap)
	case "event":
		log.Infof("Registering API Event Handler at %s", APIPath)
		rt := regRouter.NewRouter(
			router.WithHandler(event.Handler),
			router.WithResolver(rr),
			router.WithRegistry(service.Options().Registry),
		)
		ev := event.NewHandler(
			ahandler.WithNamespace(apiNamespace),
			ahandler.WithRouter(rt),
			ahandler.WithClient(service.Client()),
		)
		r.PathPrefix(APIPath).Handler(ev)
	case "http", "proxy":
		log.Infof("Registering API HTTP Handler at %s", ProxyPath)
		rt := regRouter.NewRouter(
			router.WithHandler(ahttp.Handler),
			router.WithResolver(rr),
			router.WithRegistry(service.Options().Registry),
		)
		ht := ahttp.NewHandler(
			ahandler.WithNamespace(apiNamespace),
			ahandler.WithRouter(rt),
			ahandler.WithClient(service.Client()),
		)
		r.PathPrefix(ProxyPath).Handler(ht)
	case "web":
		log.Infof("Registering API Web Handler at %s", APIPath)
		rt := regRouter.NewRouter(
			router.WithHandler(web.Handler),
			router.WithResolver(rr),
			router.WithRegistry(service.Options().Registry),
		)
		w := web.NewHandler(
			ahandler.WithNamespace(apiNamespace),
			ahandler.WithRouter(rt),
			ahandler.WithClient(service.Client()),
		)
		r.PathPrefix(APIPath).Handler(w)
	default:
		log.Infof("Registering API Default Handler at %s", APIPath)
		rt := regRouter.NewRouter(
			router.WithResolver(rr),
			router.WithRegistry(service.Options().Registry),
		)
		r.PathPrefix(APIPath).Handler(handler.Meta(service, rt, nsResolver.ResolveWithType))
	}

	// reverse wrap handler
	//plugins := append(Plugins(), plugin.Plugins()...)
	//for i := len(plugins); i > 0; i-- {
	//	h = plugins[i-1].Handler()(h)
	//}

	// create the auth wrapper and the server
	//authWrapper := auth.Wrapper(rr, nsResolver)

	api := httpapi.NewServer(Address)

	err := api.Init(opts...)
	if err != nil {
		log.Fatal(err)
		return
	}
	api.Handle("/", h)

	// Start API
	if err = api.Start(); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err = api.Stop(); err != nil {
		log.Fatal(err)
	}
}

func Commands(options ...micro.Option) []*cli.Command {
	command := &cli.Command{
		Name:  "api",
		Usage: "Run the api gateway",
		Action: func(ctx *cli.Context) error {
			Run(ctx, options...)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Usage:   "Set the api address e.g 0.0.0.0:8080",
				EnvVars: []string{"MICRO_API_ADDRESS"},
			},
			&cli.StringFlag{
				Name:    "registry",
				Usage:   "Set the registry type e.g consul, etcd",
				EnvVars: []string{"MICRO_REGISTRY"},
			},
			&cli.StringFlag{
				Name:    "registry_address",
				Usage:   "Set the registry address e.g 127.0.0.1:8500",
				EnvVars: []string{"MICRO_REGISTRY_ADDRESS"},
			},
			&cli.StringFlag{
				Name:    "handler",
				Usage:   "Specify the request handler to be used for mapping HTTP requests to services; {api, event, http, rpc}",
				EnvVars: []string{"MICRO_API_HANDLER"},
			},
			&cli.StringFlag{
				Name:    "namespace",
				Usage:   "Set the namespace used by the API e.g. com.example",
				EnvVars: []string{"MICRO_API_NAMESPACE"},
			},
			&cli.StringFlag{
				Name:    "type",
				Usage:   "Set the service type used by the API e.g. api",
				EnvVars: []string{"MICRO_API_TYPE"},
			},
			&cli.StringFlag{
				Name:    "resolver",
				Usage:   "Set the hostname resolver used by the API {host, path, grpc}",
				EnvVars: []string{"MICRO_API_RESOLVER"},
			},
			&cli.BoolFlag{
				Name:    "enable_rpc",
				Usage:   "Enable call the backend directly via /rpc",
				EnvVars: []string{"MICRO_API_ENABLE_RPC"},
			},
			&cli.BoolFlag{
				Name:    "enable_cors",
				Usage:   "Enable CORS, allowing the API to be called by frontend applications",
				EnvVars: []string{"MICRO_API_ENABLE_CORS"},
				Value:   true,
			},
		},
	}

	//for _, p := range plugin.Plugin{}() {
	//	if cmds := p.Commands(); len(cmds) > 0 {
	//		command.Subcommands = append(command.Subcommands, cmds...)
	//	}
	//
	//	if flags := p.Flags(); len(flags) > 0 {
	//		command.Flags = append(command.Flags, flags...)
	//	}
	//}

	return []*cli.Command{command}
}
