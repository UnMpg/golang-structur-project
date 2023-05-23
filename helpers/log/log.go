package log

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var Log = logrus.New()

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func SetupLogger() {
	Log.Println("Setup Logger Start")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.ToSlash("./log/log"),
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}
	multiWriter := io.MultiWriter(lumberjackLogger, os.Stderr)
	formatter := new(logrus.JSONFormatter)

	formatter.TimestampFormat = "02-01-2006 15:04:05"
	Log.SetFormatter(formatter)
	Log.SetOutput(multiWriter)
	Log.Println("Setup Logger Finish")
}

func RequestLoggerActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		writeLogReq(c)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		dur := time.Since(t)
		c.Set("Latency", dur.String())
		writeLogResp(c, blw.body.String())
	}
}

func writeLogReq(c *gin.Context) {
	if printBody(c.Request.Method) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		re := regexp.MustCompile(`\r?\n`)
		var request = re.ReplaceAllString(readBody(rdr1), "")
		if strings.Contains(c.FullPath(), "login") {
			jsonMap := make(map[string]interface{})
			json.Unmarshal([]byte(request), &jsonMap)
			delete(jsonMap, "password")
			req, _ := json.Marshal(jsonMap)
			request = string(req)
		}
		Log.WithFields(logrus.Fields{
			"logType":     "Request",
			"url":         c.Request.URL.Path,
			"method":      c.Request.Method,
			"requestId":   requestid.Get(c),
			"userAgent":   c.Request.UserAgent(),
			"requestBody": request,
		}).Info()
		c.Request.Body = rdr2
	} else {
		if !(c.FullPath() == "/") {
			Log.WithFields(logrus.Fields{
				"logType":   "Request",
				"url":       c.Request.URL.Path,
				"method":    c.Request.Method,
				"requestId": requestid.Get(c),
				"userAgent": c.Request.UserAgent(),
			}).Info()
		}
	}
}

func writeLogResp(c *gin.Context, resp string) {
	latency, _ := c.Get("Latency")
	if !(c.FullPath() == "/") {
		Log.WithFields(logrus.Fields{
			"logType":      "Response",
			"url":          c.Request.URL.Path,
			"method":       c.Request.Method,
			"requestId":    requestid.Get(c),
			"userAgent":    c.Request.UserAgent(),
			"latency":      latency.(string),
			"responseBody": resp,
		}).Info()
	}
}

func printBody(method string) bool {
	if method == "POST" {
		return true
	} else if method == "PUT" {
		return true
	} else if method == "PATCH" {
		return true
	} else {
		return false
	}
}
