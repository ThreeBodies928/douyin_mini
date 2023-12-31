package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/securemist/douyin-mini/model/resp"
	"net/http"
)

type UserListResponse struct {
	resp.Response
	UserList []resp.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	//token := c.Query("token")
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, resp.Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, resp.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: resp.Response{
			StatusCode: 0,
		},
		UserList: []resp.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: resp.Response{
			StatusCode: 0,
		},
		UserList: []resp.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: resp.Response{
			StatusCode: 0,
		},
		UserList: []resp.User{DemoUser},
	})
}
