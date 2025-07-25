package blogService

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"siteol.com/smart/src/common/mysql/blogDB"
	"strings"
	"time"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 第三者调用资源处理
func thirdAddSource(traceID string, req *blogModel.SourceAddReq) (dbReq *blogDB.Source, errNum int, err error) {
	// 创建对象初始化
	dbReq = req.ToDbReq()
	err = blogDB.SourceTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddSource Fail . Err Is : %v", err)
		// 解析数据库错误
		return nil, 1, err
	}
	// 上传图片
	err = setSourceFile(traceID, req.FileStrArray, dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddSource UploadFIle Fail . Err Is : %v", err)
		return nil, 2, err
	}
	// 更新图片
	err = blogDB.SourceTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddSource %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return nil, 1, err
	}
	return dbReq, 0, nil
}

// 第三者调用资源处理
func thirdEditSource(traceID string, req *blogModel.SourceEditReq) (errNum int, err error) {
	dbReq, err := blogDB.SourceTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "EditSource Fail . Err Is : %v", err)
		return 1, err
	}
	// 先处理图片
	err = setSourceFile(traceID, req.FileStrArray, &dbReq)
	if err != nil && err.Error() != "" {
		log.ErrorTF(traceID, "EditSource UploadFIle Fail . Err Is : %v", err)
		return 2, err
	}
	dbReq.Name = req.Name
	dbReq.FileType = req.FileType
	err = blogDB.SourceTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditSource Fail . Err Is : %v", err)
		// 解析数据库错误
		return 3, err
	}
	return 0, nil
}

var imgMap = map[string]string{
	"data:image/jpeg;": "jpg",
	"data:image/png;":  "png",
	"data:image/gif;":  "gif",
}

// 上传图片 命名风格 id.png 缩略图 id_1.png id_2.png
func setSourceFile(traceId string, baseArray []string, dbReq *blogDB.Source) (err error) {
	if len(baseArray) == 0 {
		return errors.New("DataEmpty")
	}
	// 获取基本路径
	config, err := cacheModel.GetSysConfigCache(traceId)
	if err != nil {
		return
	}
	if config == nil || config.FileFullPath == "" {
		return errors.New("Config_Empty")
	}
	// 创建子目录
	subPath := fmt.Sprintf("%d", int(dbReq.Id/1000))
	dbReq.FilePath = subPath
	// 建立完整文件目录
	fullPath := filepath.Join(config.FileFullPath, subPath)
	if _, err := os.Stat(fullPath); err != nil {
		errM := os.MkdirAll(fullPath, 0755)
		if errM != nil {
			return errM
		}
	}
	// 读取文件格式（默认取第一个）
	fileDataArray := strings.Split(baseArray[0], "base64,")
	if len(fileDataArray) != 2 {
		return errors.New("DataFormatErr")
	}
	fileBack, ok := imgMap[fileDataArray[0]]
	if !ok {
		return errors.New("DataSupportErr")
	}
	dbReq.BackEnd = fileBack
	for i, item := range baseArray {
		itemArray := strings.Split(item, "base64,")
		// 写入图片
		imageData, err := base64.StdEncoding.DecodeString(itemArray[1])
		if err != nil {
			return errors.New("DataDecodeErr")
		}
		var fileName string
		if i == 0 {
			fileName = fmt.Sprintf("%d.%s", dbReq.Id, fileBack)
		} else {
			fileName = fmt.Sprintf("%d_%d.%s", dbReq.Id, i, fileBack)
		}
		err = os.WriteFile(filepath.Join(fullPath, fileName), imageData, 0644)
		if err != nil {
			return err
		}
	}
	// 调整状态
	dbReq.Status = constant.StatusOpen
	now := time.Now()
	dbReq.UpdateAt = &now
	dbReq.Version = fmt.Sprintf("%d", now.Unix())
	return
}

// 解析数据库错误
func checkSourceDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		//if strings.Contains(errStr, "url_uni") {
		//	// 唯一索引错误
		//	return baseModel.Fail(constant.SourceUniXxxNG)
		//}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 分页查询对象封装
func sourcePageQuery(req *blogModel.SourcePageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.Name != "" {
		query.Like("name", req.Name)
	}
	if req.FileType != "" {
		query.Eq("file_type", req.FileType)
	}
	// 模拟代码，更多函数参考Query构造器
	query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
