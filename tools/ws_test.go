package tools

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var c *gin.Context

func mockHTTP() {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	cGin, r := gin.CreateTestContext(resp)
	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})

	r.GET("/test", func(c *gin.Context) {
		_, found := c.Get("profile")
		log.Println(found)
		// found is true
		c.Status(200)
	})
	cGin.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	c = cGin
}

func TestNewInitWs(t *testing.T) {
	mockHTTP()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want *InitWs
	}{
		{"TestNewInitWs", args{c.Writer, c.Request}, NewInitWs(c.Writer, c.Request)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInitWs(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInitWs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitWs_ServeWs(t *testing.T) {

	tests := []struct {
		name string
		ws   *InitWs
	}{
		{"TestInitWs_ServeWs", NewInitWs(c.Writer, c.Request)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ws.ServeWs()
		})
	}
}

func Test_writer(t *testing.T) {
	type args struct {
		ws      *websocket.Conn
		lastMod time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer(tt.args.ws, tt.args.lastMod)
		})
	}
}

func TestReadFileIfModified(t *testing.T) {
	type args struct {
		lastMod time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ReadFileIfModified(tt.args.lastMod)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileIfModified() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileIfModified() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ReadFileIfModified() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_reader(t *testing.T) {
	type args struct {
		ws *websocket.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader(tt.args.ws)
		})
	}
}
