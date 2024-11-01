package controllers

import (
	"net/http"

	"github.com/midedickson/instashop/utils"
)

func (c *Controller) Hello(w http.ResponseWriter, r *http.Request) {
	utils.Dispatch200(w, "hello, you have reached instashop api", nil)
}
