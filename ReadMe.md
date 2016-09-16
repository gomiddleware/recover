# GoMiddleware : Recover #

Middleware that recovers from panics, outputs the error to the log and writes a StatusCode of 500.

This middleware is inspired by the example given on
[Middlewares in Go: Best practices and examples](https://www.nicolasmerouze.com/middlewares-golang-best-practices-examples/)
by [Nicolas Merouze](https://www.nicolasmerouze.com/).

As with ALL [GoMiddleware](https://github.com/gomiddleware/) we adhere to the standard form of Go's middleware, which
means that it doesn't require any extra params over and above what is contained in the standard library. See the main
website [GoMiddleware](https://gomiddleware.github.io/) for more details.

## Synopsis ##

```go


```

## Other Recovery Middlewares ##

Within the [GoMiddleware](https://github.com/gomiddleware/) project, there are other 'recover' middlewares which may
perform the more specific job you're looking for. This 'recover' middleware is very simple and just outputs any
panic/error to the log, whereas other are more configurable.

* ToDo : one based on https://github.com/urfave/negroni/blob/master/recovery.go

## License ##

MIT.

(Ends)
