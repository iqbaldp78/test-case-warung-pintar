package main

import (
	// "fmt"
	// "io/ioutil"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"test-warungpintar/controllers"
)

func main() {
	// Set up a http setver.
	r := gin.Default()
	// Set blueprint request
	setBlueprintApp(r)
	setErrorHandler(r)
	r.Run(":80")
}

/*setBlueprintApp used for handling request*/
func setBlueprintApp(router *gin.Engine) {
	router.GET("/ping/server", controllers.PingServer) // this endpoint used for generate response from server
	router.GET("/ping/client", controllers.PingClient) // this endpoint used as client and request message to server
	router.GET("/ping/ws", controllers.PingWs)         // this endpoint used for handling websocket
	router.GET("/index", func(c *gin.Context) {        // this endpoint used for generate idex.html to trigger websocket
		indexTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/ping/ws")
	})

}

/*setErrorHandler used for handling 404 request*/
func setErrorHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})
}

//indexTemplate used for generate template index.html
var indexTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
