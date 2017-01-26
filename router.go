// router.go
//
// Author::    Chirantan Mitra
// Copyright:: Copyright (c) 2015-2017. All rights reserved
// License::   MIT

package main

import "net/http"

func NewRouter(logger *Logger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/dashy", DashyHandler(logger))
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	return mux
}
