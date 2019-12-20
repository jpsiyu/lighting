package ui

import (
	"github.com/jpsiyu/lighting/ui/demo"
)

func Run() {
	ins := demo.NewImage()
	/*
		images := []string{
			"https://cn.bing.com/th?id=OIP.9j-saEMnwDuVhXq00rIH2AHaFj&pid=Api&rs=1",
		}
	*/
	ins.Run(nil)
}
