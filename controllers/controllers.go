package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"test-warungpintar/message"
	"test-warungpintar/tools"
)

var (
	indexTempl = template.Must(template.New("").Parse(indexHTML))
)

const indexHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>WebSocket Example</title>
    </head>
    <body>
        <pre id="fileData">{{.Data}}</pre>
        <script type="text/javascript">
            (function() {
                var data = document.getElementById("fileData");
                var conn = new WebSocket("ws://{{.Host}}/ws?lastMod={{.LastMod}}");
                conn.onclose = function(evt) {
                    data.textContent = 'Connection closed';
                }
                conn.onmessage = function(evt) {
                    console.log('file updated');
                    data.textContent = evt.data;
                }
            })();
        </script>
    </body>
</html>
`

//Sample is controllers for request to server
func Sample(c *gin.Context) {
	urlParam := struct {
		Msg string `form:"message" binding:"required"`
	}{}
	//check url query param same with blueprint
	if err := c.ShouldBindQuery(&urlParam); err != nil {
		msg := message.New(4, err)
		c.AbortWithStatusJSON(msg.StatusCode, msg)
		return
	}

	//write response to txt
	if err := tools.WriteFile(urlParam.Msg); err != nil {
		msg := message.New(4, err)
		c.AbortWithStatusJSON(msg.StatusCode, msg)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": urlParam.Msg})
}

//ShowResponse is controllers for show previous response
func ShowResponse(c *gin.Context) {
	data, err := tools.ReadFile()
	if err != nil {
		msg := message.New(4, err)
		c.AbortWithStatusJSON(msg.StatusCode, msg)
		return
	}
	responseAll := struct {
		Total int      `json:"total"`
		Data  []string `json:"data"`
	}{
		len(data),
		data,
	}
	c.JSON(http.StatusOK, responseAll)
}

//Index is controllers for show html page and listen websocket
func Index(c *gin.Context) {

	w := c.Writer
	r := c.Request
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p, lastMod, err := tools.ReadFileIfModified(time.Time{})
	if err != nil {
		p = []byte(err.Error())
		lastMod = time.Unix(0, 0)
	}
	var v = struct {
		Host    string
		Data    string
		LastMod string
	}{
		r.Host,
		string(p),
		strconv.FormatInt(lastMod.UnixNano(), 16),
	}
	indexTempl.Execute(w, &v)
}

//RunWs used for handling websocket
func RunWs(c *gin.Context) {
	ws := tools.NewInitWs(c.Writer, c.Request)
	ws.ServeWs()
}
