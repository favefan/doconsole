package util

import (
	"encoding/json"
	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// StructAdaptiveCopyByJSON uses JSON to copy the elements in the source struct to the corresponding elements in the object struct
func StructAdaptiveCopyByJSON(source interface{}, object interface{}) error {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(sourceJSON, object); err != nil {
		return err
	}
	return nil
}

func StructBindDataValid(c *gin.Context, s interface{}) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	err := c.BindJSON(&s)
	ok, _ := valid.Valid(&s)

	if err != nil || !ok {
		// app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
}
