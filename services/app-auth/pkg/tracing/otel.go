package tracing

//
//func SetupTracer(ctx context.Context, cfg *config.Config) (*trace.TracerProvider, func(), error) {
//	var tp *trace.TracerProvider
//	var shutdownFunc func()
//	var err error
//
//	if cfg.UseJaeger {
//		tp, shutdownFunc, err = setupJaegerTracer(cfg)
//		if err != nil {
//			return nil, nil, err
//		}
//	} else {
//		tp = setupOTLPTracer(ctx, cfg)
//		shutdownFunc = func() {}
//	}
//
//	return tp, shutdownFunc, nil
//}
//
//func setupJaegerTracer(cfg *config.Config) (*trace.TracerProvider, func(), error) {
//	// Create a Jaeger configuration.
//	cfgJaeger := jaegerconfig.Configuration{
//		ServiceName: cfg.ServiceName, // Replace with your service name.
//		// Add other Jaeger configuration options as needed.
//	}
//
//	// Initialize the Jaeger tracer.
//	tracer, closer, err := cfgJaeger.NewTracer(
//		jaegerconfig.Logger(jaeger.StdLogger),
//	)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	// Create a trace provider using the Jaeger tracer.
//	tp := trace.NewTracerProvider(
//		trace.WithTracer(tracer),
//		trace.WithResource(resource.NewWithAttributes(
//			semconv.ServiceNameKey.String(cfg.ServiceName),
//			attribute.String("environment", cfg.Env),
//		)),
//	)
//
//	// Define a shutdown function.
//	shutdownFunc := func() {
//		_ = closer.Close()
//	}
//
//	return tp, shutdownFunc, nil
//}
//
//func setupOTLPTracer(ctx context.Context, cfg *config.Config) *trace.TracerProvider {
//	// Create the OTLP exporter.
//	exp, err := otlptracehttp.New(
//		ctx,
//		otlptracehttp.WithEndpoint(cfg.JaegerCollectorHost),
//		otlptracehttp.WithInsecure(),
//	)
//	if err != nil {
//		log.Fatalf("failed to create OTLP exporter: %s", err.Error())
//	}
//
//	// Create a trace provider using the OTLP exporter.
//	tp := trace.NewTracerProvider(
//		trace.WithBatcher(exp),
//		trace.WithResource(resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceName("your-app-here"),
//			attribute.String("environment", cfg.Env),
//		)),
//	)
//
//	return tp
//}
