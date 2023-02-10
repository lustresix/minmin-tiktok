// Code generated by Kitex v0.4.4. DO NOT EDIT.
package videoservice

import (
	server "github.com/cloudwego/kitex/server"
	videoservice "mini-min-tiktok/kitex_gen/videoservice"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler videoservice.VideoService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
