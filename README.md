wk
=====

wk is a smart &amp; lightweight web server engine in golang  
wk is webkit for web server  
wk is Wukong, a famouse cartoon character in Chinese epic. [wukong](http://en.wikipedia.org/wiki/Sun_Wukong)   


Roadmap
---

* 0.1 fork gomvc, refactoring. april 2013  
* 0.2 configration framework. april 2013  
* 0.3 web api server. april 2013  
* 0.4 view engine. may 2013  
* 0.5 move to go 1.1. june 2013  
* 0.6 more example. july 2013  
* 0.7 cookie, session 
* 0.8 file  
* 0.9 custome error & 404 page
* 1.0 go  

Requirements
---

go 1.0 +

Usage
---

go get github.com/sdming/pathexp  
go get github.com/sdming/kiss  
go get github.com/sdming/wk  

Document
---

Take time to translate document to english.  



Getting Started
---

	// ./demo/basic/main.go for more detail

	server, err := wk.NewDefaultServer()

	if err != nil {
		fmt.Println("DefaultServer error", err)
		return
	}

	server.RouteTable.Get("/data/top/{count}").To(...)

	server.Start()


Basic Conception
---

HttpContext is wrap of http.Request & http.Response  

HttpProcessor handle request and build HttpResult(maybe)
	
	type HttpProcessor interface {
		Execute(ctx *HttpContext)

		// Register is called once when server init  
		Register(server *HttpServer)
	}


HttpResult know how to write http.Response

	type HttpResult interface {
		Execute(ctx *HttpContext)
	}


The handle lifecyle is 

1. receive request, create HttpContext   
3. run each HttpProcessor  
4. execute HttpResult  


Route
---

Go cann't get parameter name of function by reflect, so it's a littel tricky to create parameters when call function by reflect.  


	// url: /demo/xxx/xxx
	// route to controller
	server.RouteTable.Path("/demo/{action}/{id}").ToController(controller)

	// url: /data/top/10
	// func: DataTopHandle(ctx *wk.HttpContext) (result wk.HttpResult, err error)
	// route to func (*wk.HttpContext) (wk.HttpResult, error)
	server.RouteTable.Get("/data/top/{count}").To(DataTopHandle)

	// url: /data/int/1
	// func: DataByInt(i int) *Data
	// route to a function, convert parameter by index(p0,p1,p2...)
	server.RouteTable.Get("/data/int/{p0}?").ToFunc(model.DataByInt)

	// url: /data/range/1-9
	// func: DataByIntRange(start, end int) []*Data
	// route to a function, convert parameter by index(p0,p1,p2...)
	server.RouteTable.Get("/data/range/{p0}-{p1}").ToFunc(model.DataByIntRange)

	// url: /data/int/1/xml
	// func: DataByInt(i int) *Data
	// return xml
	server.RouteTable.Get("/data/int/{p0}/xml").ToFunc(model.DataByInt).ReturnXml()

	// url: /data/int/1/json
	// func: DataByInt(i int) *Data
	// return json
	server.RouteTable.Get("/data/int/{p0}/json").ToFunc(model.DataByInt).ReturnJson()

	// url: /data/int/1/kson
	// func: DataByInt(i int) *Data
	// return custome formatted data
	server.RouteTable.Get("/data/int/{p0}/kson").ToFunc(model.DataByInt).Return(formatKson)

	// url: /data/name/1
	// func: DataByInt(i int) *Data
	// route to a function, convert parameter by name
	server.RouteTable.Get("/data/name/{id}").ToFunc(model.DataByInt).
		BindByNames("id")

	// url: /data/namerange/1-9
	// func: DataByIntRange(start, end int) []*Data
	// route to a function, convert parameter by name
	server.RouteTable.Get("/data/namerange/{start}-{end}").ToFunc(model.DataByIntRange).
		BindByNames("start", "end")

	// url: /data/namerange/?start=1&end=9
	// func: DataByIntRange(start, end int) []*Data
	// route to a function, convert parameter by name
	server.RouteTable.Get("/data/namerange/").ToFunc(model.DataByIntRange).
		BindByNames("start", "end")

	// url: post /data/post?
	// form:{"str": {"string"}, "uint": {"1024"}, "int": {"32"}, "float": {"1.1"}, "byte": {"64"}}
	// func: DataPost(data Data) string 
	// route http post to function, build struct parameter from form  
	server.RouteTable.Post("/data/post?").ToFunc(model.DataPost).BindToStruct()

	// url: post /data/postptr?
	// form:{"str": {"string"}, "uint": {"1024"}, "int": {"32"}, "float": {"1.1"}, "byte": {"64"}}
	// func DataPostPtr(data *Data) string
	// route http post to function, build struct parameter from form
	server.RouteTable.Post("/data/postptr?").ToFunc(model.DataPostPtr).BindToStruct()

	// url: delete /data/delete/1
	// func: DataDelete(i int) string 
	// route http delete to function
	server.RouteTable.Delete("/data/delete/{p0}").ToFunc(model.DataDelete)

	// url: get /data/set?str=string&uint=1024&int=32&float=3.14&byte=64
	// func: DataSet(s string, u uint64, i int, f float32, b byte) *Data 
	// test diffrent parameter type
	server.RouteTable.Get("/data/set?").ToFunc(model.DataSet).
		BindByNames("str", "uint", "int", "float", "byte")


Controller  
---

Route url like "/demo/{action}" to T, call it's method by named ({action}).     

Current version only support method of type (*HttpContext) (result HttpResult, error)  

	// url: /demo/int/32
	func (c *DemoController) Int(ctx *wk.HttpContext) (result wk.HttpResult, err error) {
		if id, ok := ctx.RouteData.Int("id"); ok {
			return wk.Json(c.getByInt(id)), nil
		}
		return wk.Data(""), nil
	}

Configration 
---

wk can manage configration for you. below is config file example, more example at test/config_test.go

	#app config file demo

	#string
	key_string: demo

	#string
	key_int: 	101

	#bool
	key_bool: 	true

	#float
	key_float:	3.14

	#map
	key_map:	{
		key1:	key1 value
		key2:	key2 value
	}

	#array
	key_array:	[
		item 1		
		item 2
	]

	#struct
	key_struct:	{
		Driver:		mysql			
		Host: 		127.0.0.1
		User:		user
		Password:	password			
	}

	#composite
	key_config:	{	
		Log_Level:	debug
		Listen:		8000

		Roles: [
			{
				Name:	user
				Allow:	[
					/user		
					/order
				]
			} 
			{
				Name:	*				
				Deny: 	[
					/user
					/order
				]
			} 
		]

		Db_Log:	{
			Driver:		mysql			
			Host: 		127.0.0.1
			User:		user
			Password:	password
			Database:	log
		}

		Env:	{
			auth:		http://auth.io
			browser:	ie, chrome, firefox, safari
		}
	}


http result
---

* ContentResult: 	html raw 
* JsonResult: 		application/json
* XmlResult: 		application/xml
* JsonpResult: 		application/jsonp
* ViewResult 		view
* PartialViewResult	partial view
* FileResult 		static file
* FileStreamResult 	stream file
* RedirectResult 	redirect
* NotFoundResult  	404
* ErrorResult 		error
* more

Event
---

Call Fire() to fire an event

	Fire(moudle, name string, source, data interface{}, context *HttpContext) 

Call On() to listen events 

	On(moudle, name string, handler Subscriber) 


View engine
---
TODO:


Example
---

TODO:  
* basic example  
* rest api example  
* customer result  
* customer processor  
* customer formatter 
* return file stream  
* BigPipe example  
* gzip  

ORM
---
Maybe, maybe not. don't have a plan yet. focus on web server.    

Validation
---
No

Css & js bundling
---
Maybe, maybe not. Do we really need it ?

Cache, gzip
---
nginx, haproxy, Varnish can provide awesome service


change log
---
* kick off
* fork gomvc, refactoring
* configration framework

Contributing
---
* github.com/sdming

License
---
Apache License 2.0  


About
----
