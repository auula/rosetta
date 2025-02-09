package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/auula/wiredkv/clog"
	"github.com/auula/wiredkv/types"
	"github.com/gorilla/mux"
)

const version = "wiredb/0.1.1"

var (
	root         *mux.Router
	authPassword string
	allowIpList  []string
	allowMethod  = []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete}
)

// http://192.168.101.225:2468/{types}/{key}
// POST 创建 http://192.168.101.225:2468/zset/user-01-score
// PUT  更新 http://192.168.101.225:2468/zset/user-01-score
// GET  获取 http://192.168.101.225:2468/table/user-01-shop-cart

func init() {
	root = mux.NewRouter()
	root.HandleFunc("/", statusController).Methods(allowMethod...)
	root.HandleFunc("/tables/{key}", tablesController).Methods(allowMethod...)
	root.Use(authMiddleware)
}

type ResponseBody struct {
	Code    int         `json:"code"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
}

func statusController(w http.ResponseWriter, r *http.Request) {}

func tablesController(w http.ResponseWriter, r *http.Request) {
	tables := []interface{}{
		types.Tables{},
		types.Tables{},
	}

	switch r.Method {
	case http.MethodGet:
		seg, err := storage.FetchSegment(mux.Vars(r)["key"])
		if err != nil {

		}
		table, err := seg.ToTables()
		if err != nil {

		}
		okResponse(w, http.StatusOK, table, "request processed successfully!")
	case http.MethodPut:
		okResponse(w, http.StatusOK, tables, "request processed successfully!")
	case http.MethodPost:
		okResponse(w, http.StatusOK, tables, "request processed successfully!")
	case http.MethodDelete:
		okResponse(w, http.StatusOK, tables, "request processed successfully!")
	default:
		methodNotAllowedResponse(w, "HTTP Protocol Method Not Allowed!")
	}
}

func okResponse(w http.ResponseWriter, code int, result interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", version)
	w.WriteHeader(code)

	resp := ResponseBody{
		Code:    code,
		Result:  result,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		clog.Error(err)
	}
}

func unauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", version)
	w.WriteHeader(http.StatusUnauthorized)

	resp := ResponseBody{
		Code:    http.StatusUnauthorized,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		clog.Error(err)
	}
}

func methodNotAllowedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", version)
	w.WriteHeader(http.StatusUnauthorized)

	resp := ResponseBody{
		Code:    http.StatusMethodNotAllowed,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		clog.Error(err)
	}
}

// 中间件函数，进行 Basic Auth 鉴权
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 使用标准 Auth 头
		token := r.Header.Get("Auth")
		clog.Debugf("HTTP request header authorization: %v", r.Header)

		// 获取客户端 IP 地址
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}

		// 检查 IP 白名单
		isAllowedIP := false
		if len(allowIpList) > 0 {
			for _, allowedIP := range allowIpList {
				if strings.Split(ip, ":")[0] == allowedIP {
					isAllowedIP = true
					break
				}
			}
		}
		if !isAllowedIP {
			clog.Warnf("Unauthorized IP address: %s", ip)
			unauthorizedResponse(w, fmt.Sprintf("Your IP %s is not allowed!", ip))
			return
		}

		if token != authPassword {
			clog.Warnf("Unauthorized access attempt from client %s", ip)
			unauthorizedResponse(w, "Access not authorised!")
			return
		}

		clog.Infof("Client %s authorized successfully", ip)
		next.ServeHTTP(w, r)

	})
}
