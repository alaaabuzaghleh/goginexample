package services

import (
	"context"
	"net/http"
	"time"

	"api.dardasha.com/db"
	"api.dardasha.com/models"
	"api.dardasha.com/requests"
	"api.dardasha.com/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var groupCollection = db.GetCollection(db.DB, "groups")
var validate = validator.New()

func CreateGroup(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var groupRequest requests.Group
	var groupModel models.Group
	var groupResponse responses.Group
	defer cancel()
	err := c.ShouldBind(&groupRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(&groupRequest)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": validationErr.Error()})
		return
	}
	groupModel = models.Group{
		Id:       primitive.NewObjectID(),
		Name:     groupRequest.Name,
		ArName:   groupRequest.ArName,
		Desc:     groupRequest.Desc,
		ArDesc:   groupRequest.ArDesc,
		SmallPic: groupRequest.SmallPic.Filename,
		LargePic: groupRequest.LargePic.Filename,
	}
	result, err := groupCollection.InsertOne(ctx, groupModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	groupResponse = responses.Group{
		Id:       result.InsertedID.(primitive.ObjectID).Hex(),
		Name:     groupModel.Name,
		ArName:   groupModel.ArName,
		Desc:     groupModel.Desc,
		ArDesc:   groupModel.ArDesc,
		SmallPic: groupModel.SmallPic,
		LargePic: groupModel.LargePic,
	}
	c.JSON(http.StatusCreated, &groupResponse)
}
