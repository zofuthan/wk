// Copyright 2012 by sdm. All rights reserved.
// license that can be found in the LICENSE file.

package wk

const (
	HttpVerbsGet     = "GET"
	HttpVerbsPost    = "POST"
	HttpVerbsPut     = "PUT"
	HttpVerbsDelete  = "DELETE"
	HttpVerbsHead    = "HEAD"
	HttpVerbsTrace   = "TRACE"
	HttpVerbsConnect = "CONNECT"
	HttpVerbsOptions = "OPTIONS"
)

const (
	HeaderAccept          = "Accept"
	HeaderAcceptCharset   = "Accept-Charset"
	HeaderAcceptEncoding  = "Accept-Encoding"
	HeaderCacheControl    = "Cache-Control"
	HeaderContentEncoding = "Content-Encoding"
	HeaderContentLength   = "Content-Length"
	HeaderContentType     = "Content-Type"
	HeaderDate            = "Date"
	HeaderEtag            = "Etag"
	HeaderExpires         = "Expires"
	HeaderLastModified    = "Last-Modified"
	HeaderLocation        = "Location"
	HeaderPragma          = "Pragma"
	HeaderServer          = "Server"
	HeaderSetCookie       = "Set-Cookie"
	HeaderUserAgent       = "User-Agent"
)

const (
	ContentTypeStream     = "application/octet-stream"
	ContentTypeJson       = "application/json"
	ContentTypeJsonp      = "application/jsonp"
	ContentTypeJavascript = "application/javascript"
	ContentTypeHTML       = "text/html"
	ContentTypeXml        = "text/xml"
	ContentTypeCss        = "text/css"
	ContentTypePlain      = "text/plain"
	ContentTypeGif        = "image/gif"
	ContentTypeIcon       = "image/x-icon"
	ContentTypeJpeg       = "image/jpeg"
	ContentTypePng        = "image/png"
)

//application/x-www-form-urlencoded
//multipart/form-data

const (
	msgServerTimeout     = "server timeout"
	msgServerInternalErr = "server internal error"
	msgNotFound          = "404 page not found"
	msgNoResult          = "no result"
	msgNoAction          = "can not find action"
	StatusNotModified
)

const (
	codeServerInternaError = 500
)

const (
	LogError = iota
	LogInfo
	LogDebug
)

const (
	_root           = "/"
	_any            = "*"
	_route          = "_route"
	_static         = "_static"
	_render         = "_render"
	_action         = "action"
	_notFoundAction = "noaction"
	_defaultAction  = "default"
	_serverName     = "go web server "
)

const (
	_wkWebServer             = "_webserver"
	_eventStartRequest       = "start_request"
	_eventEndRequest         = "end_request"
	_eventStartExecute       = "start_execute"
	_eventEndExecute         = "end_execute"
	_eventStartResultExecute = "start_result"
	_eventEndResultExecute   = "end_result"
	_eventStartAction        = "start_action"
	_eventEndAction          = "end_action"
)

const (
	_defaultSize = 61
)