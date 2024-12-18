package pkg

import (

	"fmt"
	"net/http"
)


func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add rate limiter logic here
		
		//ctx := context.Background()

		//database :=  database.DatabaseImpl{
		//	Client: config.ConfigRedis(),
		//}
		
		//database.ConsultarParametros(ctx)
		//database.InserirParametros(ctx)

		
		fmt.Println("Teste rate limiter")
		
		next.ServeHTTP(w, r)
	})
}