package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Rule struct{}

// 空结构体，只是给接口方法分类，或者用函数也可以
func NewRule() Rule {
	return Rule{}
}

func (r Rule) Hello(c *gin.Context) {
	fmt.Println("hello")
}
