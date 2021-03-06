package yugle

import (
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/page_processer"
	"github.com/hu17889/go_spider/core/spider"
	"time"
)

var shotOnOnePlusUrl = "https://photos.oneplus.com/cn"
var ShotOnOnePlusTaskName = "Shot on OnePlus爬虫"

type ShotOnOnePlusProcessor struct {
	processor page_processer.PageProcesser
}

type ShotOnOnePlusPipeline struct {
}

func NewShotOnOnePlusProcessor() *ShotOnOnePlusProcessor {
	return &ShotOnOnePlusProcessor{}
}

func (spider *ShotOnOnePlusProcessor) Process(p *page.Page) {
	//log.Println(p)
	query := p.GetHtmlParser()
	//获取壁纸链接
	img, imgExists := query.Find(`p[class="p-img"]`).ChildrenFiltered(`img`).First().Attr("data-src")
	if imgExists {
		p.AddField("picture", img)
	}
	//获取壁纸标题
	title := query.Find(`p[class="p-title"]`).First().Text()
	p.AddField("title", title)
	//获取壁纸作者
	author := query.Find(`#all > main > div.card.card-block.text-center.photo-simple-info > div > div.card-text > span`).First().Text()
	p.AddField("author", author)
}

func (spider *ShotOnOnePlusProcessor) Finish() {
	onePlusTask := GetTaskByTaskName(ShotOnOnePlusTaskName)
	onePlusTask.State = Idle
	onePlusTask.LastSuccess = true
	UpdateTask(onePlusTask, crons[ShotOnOnePlusTaskName])
}

func NewShotOnOnePlusPipeline() *ShotOnOnePlusPipeline {
	return &ShotOnOnePlusPipeline{}
}

func (pipeline *ShotOnOnePlusPipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	db := DbConnect()
	defer db.Close()
	picture, pictureOk := items.GetItem("picture")
	title, titleOk := items.GetItem("title")
	author, authorOk := items.GetItem("author")
	if pictureOk && titleOk && authorOk {
		shotOnOnePlusPicture := &ShotOnOnePlusPicture{Picture: picture, Title: title, Author: author, Date: time.Now().Local().String()[:10]}
		shotOnOnePlusPictureInDb := &ShotOnOnePlusPicture{}
		db.Where("picture = ?", picture).Find(shotOnOnePlusPictureInDb)
		if shotOnOnePlusPictureInDb.Picture == "" {
			db.Create(shotOnOnePlusPicture)
		}
	}
}

func CrawlShotOnOnePlusPicture() {
	onePlusTask := GetTaskByTaskName(ShotOnOnePlusTaskName)
	onePlusTask.State = Running
	SaveTask(onePlusTask)
	spider.NewSpider(NewShotOnOnePlusProcessor(), ShotOnOnePlusTaskName).
		AddUrl(shotOnOnePlusUrl, "html").
		AddPipeline(NewShotOnOnePlusPipeline()).
		SetThreadnum(1).
		Run()
}
