package utils

import (
	"net"
	"net/http"
	"strings"
)

// GetRealIP obtiene la IP real del cliente considerando proxies
func GetRealIP(r *http.Request) string {
	// Primero intentar obtener IP de headers comunes de proxies
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Tomar la primera IP si hay múltiples
		ips := strings.Split(forwarded, ",")
		if len(ips) > 0 {
			ip := strings.TrimSpace(ips[0])
			if net.ParseIP(ip) != nil {
				return ip
			}
		}
	}

	// Intentar con otros headers
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		if net.ParseIP(realIP) != nil {
			return realIP
		}
	}

	// Si no hay headers, usar RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}