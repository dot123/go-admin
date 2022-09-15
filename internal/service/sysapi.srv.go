package service

import (
	"bytes"
	"context"
	"github.com/bitly/go-simplejson"
	"github.com/dot123/gin-gorm-admin/internal/errors"
	"github.com/dot123/gin-gorm-admin/internal/models/system"
	"github.com/dot123/gin-gorm-admin/internal/schema"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"io/ioutil"
	"regexp"
	"strings"
)

var SysApiSet = wire.NewSet(wire.Struct(new(SysApiSrv), "*"))

type SysApiSrv struct {
	SysApiRepo *system.SysApiRepo
}

func (s *SysApiSrv) GetPage(ctx context.Context, req *schema.SysApiGetPageReq) (*schema.SysApiGetPageResp, error) {
	list, total, err := s.SysApiRepo.GetPage(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	result := new(schema.SysApiGetPageResp)
	result.Total = total
	copier.Copy(&result.Data, list)

	return result, err
}

func (s *SysApiSrv) Get(ctx context.Context, id uint64) (*schema.SysApi, error) {
	model, err := s.SysApiRepo.Get(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("查看对象不存在或无权查看")
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return nil, errors.NewDefaultResponse("查看对象失败")
	}

	result := new(schema.SysApi)
	copier.Copy(result, model)

	return result, nil
}

func (s *SysApiSrv) Update(ctx context.Context, req *schema.SysApiUpdateReq) error {
	model, err := s.SysApiRepo.Get(ctx, req.ID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("无权更新该数据")
	}

	copier.Copy(model, req)
	err = s.SysApiRepo.Update(ctx, model)
	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("更新api失败")
	}
	return err
}

func (s *SysApiSrv) Delete(ctx context.Context, req *schema.SysApiDeleteReq) error {
	rowsAffected, err := s.SysApiRepo.Delete(ctx, &req.Ids)
	if rowsAffected == 0 {
		return errors.NewDefaultResponse("无权删除该数据")
	}

	if err != nil {
		logger.WithContext(ctx).Errorf("db error:%s", err.Error())
		return errors.NewDefaultResponse("删除api失败")
	}
	return nil
}

func (s *SysApiSrv) SaveSysApi(ctx context.Context, engine *gin.Engine) {
	list := make([]schema.Router, 0)
	routers := engine.Routes()
	for _, router := range routers {
		list = append(list, schema.Router{RelativePath: router.Path, Handler: router.Handler, HttpMethod: router.Method})
	}
	for _, v := range list {
		if v.HttpMethod != "HEAD" ||
			strings.Contains(v.RelativePath, "/swagger/") ||
			strings.Contains(v.RelativePath, "/static/") {

			// 根据接口方法注释里的@Summary填充接口名称，适用于代码生成器
			// 可在此处增加配置路径前缀的if判断，只对代码生成的自建应用进行定向的接口名称填充
			jsonFile, _ := ioutil.ReadFile("docs/swagger.json")
			jsonData, _ := simplejson.NewFromReader(bytes.NewReader(jsonFile))
			urlPath := v.RelativePath

			idPatten := "(.*)/:(\\w+)" // 正则替换，把:id换成{id}
			reg, _ := regexp.Compile(idPatten)
			if reg.MatchString(urlPath) {
				urlPath = reg.ReplaceAllString(v.RelativePath, "${1}/{${2}}") // 把:id换成{id}
			}

			urlPath = strings.Replace(urlPath, "/api/v1", "", -1) // 去除基础路径

			apiTitle, _ := jsonData.Get("paths").Get(urlPath).Get(strings.ToLower(v.HttpMethod)).Get("summary").String()

			if err := s.SysApiRepo.Create(ctx, &v, apiTitle); err != nil {
				logger.WithContext(ctx).Errorf("SaveSysApi error: %s \r\n ", err.Error())
			}
		}
	}
}
