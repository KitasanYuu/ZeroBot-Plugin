package main

import (
	"fmt"
	"test1/ID_Search"
	"test1/handle"
	
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"time"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	"github.com/FloatTech/zbputils/file"
	"github.com/FloatTech/zbputils/img"
	"github.com/FloatTech/zbputils/img/text"
	"github.com/FloatTech/zbputils/img/writer"
	"github.com/FloatTech/zbputils/web"
	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)
engine.OnRegex(`^>查id\s?(.{1,25})$`, getdb).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			keyword := ctx.State["regex_matched"].([]string)[1]}

func main() {
	Users := "https://robertsspaceindustries.com/citizens/" + keyword

	h := handle.UserHandle{}
	Search := ID_Search.NewSearch()
	request, err := ID_Search.NewRequest("GET", Users, ID_Search.UserAgent, &h, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	Search.Request = request
	Search.Visit()
}
