package monitoring

import(
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	LoginAttempts = promauto.NewCounter(prometheus.CounterOpts{
		Name: "mediapp_login_intentos",
		Help: "Numero total de intentos de inicio de sesión",
	})

	LoginErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "login_errores_total",
		Help: "Numero total de errores en el inicio de sesión",
	})
)



