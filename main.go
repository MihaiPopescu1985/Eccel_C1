package main

import (
	"net/http"
	"time"

	"example.com/c1/c1device"
	"example.com/c1/model"
	"example.com/c1/util"
	"example.com/c1/web/controller"
	"github.com/gorilla/mux"
)

const serverPort string = ":8181"

func main() {

	util.InitLogger()

	err := model.Db.Init("")
	if err != nil {
		util.Log.Println(err)
	}
	go func() {
		for {
			if err := model.Db.Connect(); err != nil {
				util.Log.Println(err)
			}
			time.Sleep(time.Minute)
		}
	}()

	dev1 := c1device.C1Device{
		IP:        "192.168.0.91",
		WsChannel: make(chan []byte),
	}
	go func() {
		for {
			err := dev1.UseDevice()
			if err != nil {
				util.Log.Println(err)
				time.Sleep(time.Minute)
				err = nil
			}
		}
	}()
	dev2 := c1device.C1Device{
		IP:        "192.168.0.92",
		WsChannel: make(chan []byte),
	}
	go func() {
		for {
			err := dev2.UseDevice()
			if err != nil {
				util.Log.Println(err)
				time.Sleep(time.Minute)
				err = nil
			}
		}
	}()
	dev3 := c1device.C1Device{
		IP:        "192.168.0.91",
		WsChannel: make(chan []byte),
	}
	go func() {
		for {
			err := dev3.UseDevice()
			if err != nil {
				util.Log.Println(err)
				time.Sleep(time.Minute)
				err = nil
			}
		}
	}()
	dev4 := c1device.C1Device{
		IP:        "192.168.0.92",
		WsChannel: make(chan []byte),
	}
	go func() {
		for {
			err := dev4.UseDevice()
			if err != nil {
				util.Log.Println(err)
				time.Sleep(time.Minute)
				err = nil
			}
		}
	}()

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomePageHandler)
	router.HandleFunc("/log-out", controller.LogOutHandler)
	router.HandleFunc("/error", controller.ErrorPageHandler)

	router.NotFoundHandler = controller.PageNotFoundHandler{}
	router.Use(controller.AuthMiddleware)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/view/"))))
	util.Log.Println(http.ListenAndServe(serverPort, router))
}
