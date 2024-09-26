package utils

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"unicode"
)

const ContextUserKey string = "user_id"

var BadRequestError = errors.New("Bad Request")

// Для того, чтобы узнать, есть ли лишние поля, нужно сравнить интерфейс, который просто декодировали полностью (data) и значение v, структуры
// Если true -- то наверху обрабатываем как Bad
func HasUnknownFields(data map[string]interface{}, val interface{}) bool {
	v := reflect.ValueOf(val)

	fields := make(map[string]struct{})

	// Get all struct fields
	for i := 0; i < v.Type().NumField(); i++ {
		field := v.Type().Field(i)
		fieldName := field.Tag.Get("json")
		fields[fieldName] = struct{}{}
	}

	//// Check for any keys in the request that are not in the struct fields
	for key := range data {
		if _, ok := fields[key]; !ok {
			return true
		}
	}

	return false
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)                   // Convert string to runes to handle multi-byte characters
	runes[0] = unicode.ToUpper(runes[0]) // Capitalize the first rune
	return string(runes)
}

func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal(fmt.Sprintf("%s must be set in env", key))
	}
	return val
}

func IsProd() bool {
	mode := os.Getenv("mode")

	return mode == "production"
}

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading %s file", path))
	}
}

func GetIPAddress(r *http.Request) string {
	// If behind reverse proxy
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return strings.Split(ip, ":")[0] // Remove port if included
}

func GetBearerToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		bearer = "Bearer "

	}
	return bearer
}

func IpFromRequest(r *http.Request) string {
	// Try to get the IP from the X-Forwarded-For header (used by proxies)
	// This will be a list of IPs, where the first one is the original client IP
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// If there are multiple IPs, take the first one
		ip = strings.Split(ip, ",")[0]
		return strings.TrimSpace(ip)
	}

	// Try to get the IP from the X-Real-IP header (another common header for proxies)
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return strings.TrimSpace(ip)
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}
