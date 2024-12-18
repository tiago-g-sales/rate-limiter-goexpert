package web

import (
	"encoding/json"
	"net"
	"net/http"

	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/model"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/internal/service"
	"github.com/tiago-g-sales/rate-limiter-goexpert/client/pkg"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)



type Webserver struct {
	TemplateData *TemplateData
}

// NewServer creates a new server instance
func NewServer(templateData *TemplateData) *Webserver {
	return &Webserver{
		TemplateData: templateData,
	}
}

// createServer creates a new server instance with go chi router
func (we *Webserver) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(pkg.RateLimiter)
	router.Use(
		middleware.SetHeader("Content-Type", "application/json"),
	)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	// promhttp
	//router.Handle("/metrics", promhttp.Handler())
	router.Get("/", we.HandleRequest)
	return router
}

type TemplateData struct {
	Title              string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}

const(
	API_KEY = "API_KEY"
	INVALID_IP_ADRESS = "invalid Ip Adress"
	NOTFOUND_ZIP_COD = "can not find zipcode"
	LEN_ZIP_CODE = 8
)

func (h *Webserver) HandleRequest(w http.ResponseWriter, r *http.Request) {

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := h.TemplateData.OTELTracer.Start(ctx, "SPAN_INICIAL "+h.TemplateData.RequestNameOTEL)
	spanInicial.End()

	ctx, span := h.TemplateData.OTELTracer.Start(ctx, "Chamada externa "+h.TemplateData.RequestNameOTEL)
	defer span.End()


	parameter := service.FormatParameter(ctx)
	
	ip,_, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, INVALID_IP_ADRESS, http.StatusUnprocessableEntity)
		return
	}	
	parameter.Ip = ip
	parameter.ApiKey = r.Header.Get(API_KEY) 


	result := service.GetParameter(ctx, *parameter)
	if result != nil{
		service.ValidateRateLimiter(ctx, *result)
	}else {
		service.InserirParametros(ctx, *parameter)
	}


	if h.TemplateData.ExternalCallURL != "" {
		var req *http.Request
		var err error

		var dto model.Parameter
		err = json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, INVALID_IP_ADRESS, http.StatusUnprocessableEntity)
			return
		}






		if h.TemplateData.ExternalCallMethod == "GET" {
			req, err = http.NewRequestWithContext(ctx, "GET", h.TemplateData.ExternalCallURL, nil)
		} else if h.TemplateData.ExternalCallMethod == "POST" {
			req, err = http.NewRequestWithContext(ctx, "POST", h.TemplateData.ExternalCallURL, nil)
		} else {
			http.Error(w, "Invalid ExternalCallMethod", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		
		

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(temp)

	

	

	}


}
