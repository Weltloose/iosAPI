package views

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/Weltloose/testGo/dal/mongodb"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username, password, ", username, password)
	if mongodb.ValidateUser(username, password) {
		c.JSON(200, gin.H{
			"ok": true,
		})
	} else {
		c.JSON(200, gin.H{
			"ok": false,
		})
	}
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username, password, ", username, password)
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("get: ", c.Request.FormValue("username"))
	fmt.Println("data: ", string(data))
	if !mongodb.UserNameExist(username) && mongodb.AddUser(username, password) {
		c.JSON(200, gin.H{
			"ok": true,
		})
	} else {
		c.JSON(200, gin.H{
			"ok": false,
		})
	}
}

func CreateGroup(c *gin.Context) {
	username := c.PostForm("username")
	groupname := c.PostForm("groupname")
	fmt.Println("username groupname ", username, groupname)
	id := mongodb.CreateGroup(username, groupname)
	c.JSON(200, gin.H{
		"groupID": id,
	})
}

func JoinGroup(c *gin.Context) {
	username := c.PostForm("username")
	groupIDStr := c.PostForm("groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		fmt.Println("groupID error ", err)
		c.JSON(200, gin.H{
			"ok": false,
		})
	}
	success := mongodb.JoinGroup(username, groupID)
	c.JSON(200, gin.H{
		"ok": success,
	})
}

func LeaveGroup(c *gin.Context) {
	username := c.PostForm("username")
	groupIDStr := c.PostForm("groupID")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		fmt.Println("groupID error ", err)
		c.JSON(200, gin.H{
			"ok": false,
		})
	}
	success := mongodb.LeaveGroup(username, groupID)
	c.JSON(200, gin.H{
		"ok": success,
	})
}

func GetGroupList(c *gin.Context) {
	username := c.Query("username")
	groupList := mongodb.GetGroupList(username)
	groupName := mongodb.GetGroupName(groupList)
	c.JSON(200, gin.H{
		"groupList":  groupList,
		"groupnames": groupName,
	})
}

func GetEventList(c *gin.Context) {
	username := c.Query("username")
	fmt.Println("form: ", c.Request.FormValue("username"))
	fmt.Println("postform: ", username)
	eventList := mongodb.GetEventList(username)
	c.JSON(200, gin.H{
		"events": eventList,
	})
}

func AddEvent(c *gin.Context) {
	addEvent := struct {
		GroupID  int
		FromTime string
		ToTime   string
		Theme    string
		Desc     string
	}{}
	addEvent.GroupID, _ = strconv.Atoi(c.PostForm("groupID"))
	addEvent.FromTime = c.PostForm("timeFrom")
	addEvent.ToTime = c.PostForm("timeTo")
	addEvent.Theme = c.PostForm("theme")
	addEvent.Desc = c.PostForm("desc")
	mongodb.AddEvent(addEvent.GroupID, addEvent.FromTime, addEvent.ToTime, addEvent.Theme, addEvent.Desc)
}

func EditEvent(c *gin.Context) {
	editEvent := struct {
		GroupID  int
		EventID  int
		FromTime string
		ToTime   string
		Theme    string
		Desc     string
	}{}
	editEvent.GroupID, _ = strconv.Atoi(c.PostForm("groupID"))
	editEvent.EventID, _ = strconv.Atoi(c.PostForm("eventID"))
	editEvent.FromTime = c.PostForm("timeFrom")
	editEvent.ToTime = c.PostForm("timeTo")
	editEvent.Theme = c.PostForm("theme")
	editEvent.Desc = c.PostForm("desc")

	mongodb.EditEvent(editEvent.GroupID, editEvent.EventID, editEvent.FromTime, editEvent.ToTime, editEvent.Theme, editEvent.Desc)

}

func DeleteEvent(c *gin.Context) {
	groupID, _ := strconv.Atoi(c.PostForm("groupID"))
	eventID, _ := strconv.Atoi(c.PostForm("eventID"))
	mongodb.DeleteEvent(groupID, eventID)
}
