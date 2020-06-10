package controllers

import (
	"fmt"
	"os"
)

type DocController struct {
	BaseController
}

func (c *DocController) Index(){
	c.TplName = "document/markdown_edit_template.html"
}

//上传附件
func (c *DocController) Upload() {
	size := c.BaseController.Ctx.Request.Form.Get("size")
	fileName := c.BaseController.Ctx.Request.Form.Get("fileName")
	chunk := c.BaseController.Ctx.Request.Form.Get("chunk")
	chunks := c.BaseController.Ctx.Request.Form.Get("chunks")
	isChunked := c.BaseController.Ctx.Request.Form.Get("isChunked")
	chunkSize := c.BaseController.Ctx.Request.Form.Get("chunkSize")
	md5value := c.BaseController.Ctx.Request.Form.Get("md5value")
	lastModifiedDate := c.BaseController.Ctx.Request.Form.Get("lastModifiedDate")
	id := c.BaseController.Ctx.Request.Form.Get("id")
	name := c.BaseController.Ctx.Request.Form.Get("name")

	fmt.Println("size: ", size)
	fmt.Println("fileName: ", fileName)
	fmt.Println("chunk: ", chunk)
	fmt.Println("chunks: ", chunks)
	fmt.Println("isChunked: ", isChunked)
	fmt.Println("chunkSize: ", chunkSize)
	fmt.Println("md5value: ", md5value)
	fmt.Println("lastModifiedDate: ", lastModifiedDate)
	fmt.Println("id: ", id)
	fmt.Println("name: ", name)
	fmt.Println("--------------------------------------------------------------------------------------: ")
	fmt.Println("--------------------------------------------------------------------------------------: ")

	file := c.Ctx.Request.MultipartForm.File
	fmt.Println(file["file"][0].Filename)
	fmt.Println(file["file"][0].Header)
	fmt.Println(file["file"][0].Size)
	file2, err := file["file"][0].Open()
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}

	bytes := make([]byte, 1024*1024*2)
	file2.Read(bytes)
	fmt.Println(len(bytes))
	fmt.Println(len(bytes))
	openFile, err := os.OpenFile("/Users/dxm/go/src/ziyoubiancheng/mbook/temp.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	openFile.Write(bytes)
	c.Data["json"] = "ok"
	c.ServeJSON()


	// 参数1：存放文件中的数据信息
	// 参数2：moreFile中存放我们上传的文件大小，FileName
	// file, moreFile, err := c.GetFile(name)
	//if err == http.ErrMissingFile { // 对比错误，看看是否真的真的将文件上传上来了～
	//	name = "editormd-image-file"
	//	file, moreFile, err = c.GetFile(name)
	//	if err == http.ErrMissingFile {
	//		c.JsonResult(1, "文件错误")
	//	}
	//}
	/*if err != nil {
		c.JsonResult(1, err.Error())
	}
	// get出file后，不要忘记了将文件流关闭掉
	defer file.Close()
	// 获取出文件的拓展名
	ext := filepath.Ext(moreFile.Filename)
	if ext == "" {
		c.JsonResult(1, "文件格式错误")
	}
	// 判断上传的文件的拓展名是否在允许的范围内
	if !common.IsAllowedFileExt(ext) {
		c.JsonResult(1, "文件类型错误")
	}
	// 查询出这本书的信息
	bookId := 0
	//如果是超级管理员，则不判断权限
	if c.Member.IsAdministrator() {
		book, err := models.NewBook().Select("identify", identify)
		if err != nil {
			c.JsonResult(1, "文档不存在或权限不足")
		}
		bookId = book.BookId
	} else {
		book, err := models.NewBookData().SelectByIdentify(identify, c.Member.MemberId)
		if err != nil {
			if err == orm.ErrNoRows {
				c.JsonResult(1, "权限错误")
			}
			c.JsonResult(6001, err.Error())
		}
		//没有编辑权限
		if book.RoleId != common.BookEditor && book.RoleId != common.BookAdmin && book.RoleId != common.BookFounder {
			c.JsonResult(1, "权限错误")
		}
		bookId = book.BookId
	}
	// 查询出这篇具体的doc信息
	if docId > 0 {
		doc, err := models.NewDocument().SelectByDocId(docId)
		if err != nil {
			c.JsonResult(1, "获取文档错误")
		}
		if doc.BookId != bookId {
			c.JsonResult(1, "获取文档错误")
		}
	}
	// 构造文件名，文件需要被保存的路径
	fileName := strconv.FormatInt(time.Now().UnixNano(), 16)
	filePath := filepath.Join(common.WorkingDirectory, "uploads", time.Now().Format("200601"), fileName+ext)
	path := filepath.Dir(filePath) // 获取到一个文件所在到目录信息
	// 创建目录，如果已经存在了，会直接忽略掉
	os.MkdirAll(path, os.ModePerm)
	// 调用beegoController中的方法，将form表单中的数据保存进给定的目的地址中
	err = c.SaveToFile(name, filePath)

	if err != nil {
		c.JsonResult(1, "保存文件失败")
	}
	attachment := models.NewAttachment() // 创建对应的附件对象，并完善附件的信息
	attachment.BookId = bookId
	attachment.Name = moreFile.Filename
	attachment.CreateAt = c.Member.MemberId
	attachment.Ext = ext
	attachment.Path = strings.TrimPrefix(filePath, common.WorkingDirectory)
	attachment.DocumentId = docId
	// Stat方法会返回文件的 FileInfo，比如name、size、mode
	if fileInfo, err := os.Stat(filePath); err == nil {
		attachment.Size = float64(fileInfo.Size())
	}
	if docId > 0 {
		attachment.DocumentId = docId
	}
	// 如果附件的后缀名是图片，todo ---
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") || strings.EqualFold(ext, ".png") || strings.EqualFold(ext, ".gif") {

		attachment.HttpPath = "/" + strings.Replace(strings.TrimPrefix(filePath, common.WorkingDirectory), "\\", "/", -1)
		if strings.HasPrefix(attachment.HttpPath, "//") {
			attachment.HttpPath = string(attachment.HttpPath[1:])
		}
		isAttach = false
	}
	// 将构造的本条accachment记录插入数据库
	err = attachment.Insert()
	// 如果保存失败了，那么将刚才上传上来的文件删除掉
	if err != nil {
		os.Remove(filePath)
		c.JsonResult(1, "文件保存失败")
	}
	if attachment.HttpPath == "" {
		attachment.HttpPath = beego.URLFor("DocController.DownloadAttachment", ":key", identify, ":attach_id", attachment.AttachmentId)

		if err := attachment.Update(); err != nil {
			c.JsonResult(1, "保存文件失败")
		}
	}
	osspath := fmt.Sprintf("projects/%v/%v", identify, fileName+filepath.Ext(attachment.HttpPath))

	osspath = "uploads/" + osspath
	if err := store.SaveToLocal("."+attachment.HttpPath, osspath); err != nil {
		beego.Error(err.Error())
	}
	attachment.HttpPath = "/" + osspath

	result := map[string]interface{}{
		"errcode":   0,
		"success":   1,
		"message":   "ok",
		"url":       attachment.HttpPath,
		"alt":       attachment.Name,
		"is_attach": isAttach,
		"attach":    attachment,
	}
	c.Ctx.Output.JSON(result, true, false)
	c.StopRun()*/
}