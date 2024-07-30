package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
	"net/http"
	"shop/models"
	pb "shop/proto/goods"
	"strconv"
)

var (
	service = "goods"
	version = "latest"
)

type V2Controller struct{}

// func (con V2Controller) AddGoods(c *gin.Context) {
//
//		// Create client
//		client := pb.NewGoodsService(service, models.MicroClient)
//
//		// Call service
//		var rsp, err = client.AddGoods(context.Background(), &pb.AddGoodsReq{
//			Parame: &pb.GoodsModel{
//				Title:   "增加数据",
//				Price:   5.5,
//				Content: "好吃的东西",
//			},
//		})
//		if err != nil {
//			logger.Fatal(err)
//		}
//
//		logger.Info(rsp)
//
//		c.JSON(200, gin.H{
//			"message": rsp.Msg,
//			"success": rsp.Success,
//		})
//	}
func (con V2Controller) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	srv := pb.NewGoodsService(service, models.MicroClient)
	rsp, err := srv.GetGoods(context.Background(), &pb.GetGoodsRequest{Page: int64(page), PageSize: 10})
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(123)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rsp,
		"长度":      len(rsp.GoodsList),
	})
}
func (con V2Controller) Plist(c *gin.Context) {
	c.String(200, "我是一个api接口-Plist")
}
