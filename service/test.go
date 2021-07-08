package service

import "go-web-template/util"

func Hello() util.Result {
	return util.SuccessInfo("Hello,World!")
}
