package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func feedHandler(c *gin.Context) {
	//check IsLogin
	p := NewFeedVideoList(c) //todo
	token, ok := c.GetQuery("token")
	if !ok {
		err := p.InvalidToken()
		if err != nil {
			p.FeedVideoListError(err.Error())
		}
		return
	}
	//IsLogin
	err := p.ValidToken(token)
	if err != nil {
		p.FeedVideoListError(err.Error())
	}
}

type FeedVideoList struct {
	*gin.Context
}

func NewFeedVideoList(c *gin.Context) *FeedVideoList {
	return &FeedVideoList{Context: c}
}

// Visitor
func (p *FeedVideoList) InvalidToken() error {
	rawTimestamp := p.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
	}
	videoList, err := video.QueryFeedVideoList(0, latestTime)
	if err != nil {
		return err
	}
	p.FeedVideoListOk(videoList)
	return nil
}

// IsLogin todo 改验证方式
func (p *FeedVideoList) ValidToken(token string) error {
	//解析成功
	//查token
	if claim, ok := middleware.ParseToken(token); ok {
		//token超时
		if time.Now().Unix() > claim.ExpiresAt {
			return errors.New("token超时")
		}
		rawTimestamp := p.Query("latest_time")
		var latestTime time.Time
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err != nil {
			latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
		}
		//调用service层接口
		videoList, err := video.QueryFeedVideoList(claim.UserId, latestTime)
		if err != nil {
			return err
		}
		p.FeedVideoListOk(videoList)
		return nil
	}
	//解析失败
	return errors.New("token不正确")
}
func (p *FeedVideoList) FeedVideoListError(msg string) {
	p.JSON(http.StatusOK, FeedResponse{CommonResponse: models.CommonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *FeedVideoList) FeedVideoListOk(videoList *video.FeedVideoList) {
	p.JSON(http.StatusOK, FeedResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 0,
		},
		FeedVideoList: videoList,
	},
	)
}
