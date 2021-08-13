package service

import "log"

func CreateCode(id string) (interface{}, error) {

	log.Println("service in ...")

	return "ok", nil
}
