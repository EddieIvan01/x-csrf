package main

import (
    "github.com/gin-gonic/gin"
    "github.com/eddieivan01/x-csrf"
)

func main(){
    r := gin.Default()

    // custom the CSRF token config as you like
    csrf.TokenLength = 32
	csrf.TokenKey = "the-key-name"
	csrf.TokenCookie = "the-cookie-name"
	csrf.DefaultExpire = 3600 * 6
    csrf.RandomSec = false
    
    // every site will set cookie
    r.Use(xcsrf.SetCSRFToken())

    r.GET("/", func(c *gin.Context){
        c.String(200, "ok")
    })
    
    // secret sites, which need to defend CSRF attack
    g1 := r.Group("/secret")
    g1.Use(xcsrf.XCSRF("header"))
    g1.GET("/", func(c *gin.Context){
        c.String(200, "csrf ok")
    })
    r.Run(":5000")
}
