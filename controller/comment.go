package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/resp"
)

type CommentListResponse struct {
	resp.Response
	CommentList []resp.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	resp.Response
	Comment resp.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	//token := c.Query("token")
	//actionType := c.Query("action_type")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	if actionType == "1" {
	//		text := c.Query("comment_text")
	//		c.JSON(http.StatusOK, CommentActionResponse{Response: resp.Response{StatusCode: 0},
	//			Comment: resp.Comment{
	//				Id:         1,
	//				User:       user,
	//				Content:    text,
	//				CreateDate: "05-01",
	//			}})
	//		return
	//	}
	//	c.JSON(http.StatusOK, resp.Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, resp.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	//c.JSON(http.StatusOK, CommentListResponse{
	//	Response:    resp.Response{StatusCode: 0},
	//	CommentList: DemoComments,
	//})
}
