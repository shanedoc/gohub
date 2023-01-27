package paginator

import (
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/config"
	"github.com/shanedoc/gohub/pkg/logger"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//处理分页逻辑
type Paging struct {
	CurrentPage int    // 当前页
	PerPage     int    // 每页条数
	TotalPage   int    // 总页数
	TotalCount  int64  // 总条数
	NextPageURL string // 下一页的链接
	PrevPageURL string // 上一页的链接
}

//Paginator 分页操作类
type Paginator struct {
	BaseURL    string // 用以拼接 URL
	PerPage    int    // 每页条数
	Page       int    // 当前页
	Offset     int    // 数据库读取数据时 Offset 的值
	TotalCount int64  // 总条数
	TotalPage  int    // 总页数 = TotalCount/PerPage
	Sort       string // 排序规则
	Order      string // 排序顺序

	query *gorm.DB     // db query 句柄
	ctx   *gin.Context // gin context，方便调用
}

func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseURL string, perPage int) Paging {
	//初始化Paginator 实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseURL)
	//查询数据库
	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error
	//数据库出错
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}
	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

//初始化分页必须用到的属性，基于这些属性查询数据库
func (p *Paginator) initProperties(perPage int, baseURL string) {
	p.BaseURL = p.formatBaseURL(baseURL)
	p.PerPage = p.getPerPage(perPage)
	//排序参数（控制器中以验证过这些参数，可放心使用）
	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")
	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage

}

//兼容url带与不带 `?` 的情况
func (p *Paginator) formatBaseURL(baseURL string) string {
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseURL = baseURL + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseURL
}

func (p Paginator) getPerPage(perPage int) int {
	// 优先使用请求 per_page 参数
	queryPerpage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPerpage) > 0 {
		perPage = cast.ToInt(queryPerpage)
	}
	//没有传参使用默认
	if perPage <= 0 {
		perPage = config.GetInt("paging.perpage")
	}
	return perPage
}

//getTotalCount 返回的是数据库里的条数
func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

//getCurrentPage 返回当前页码
func (p Paginator) getCurrentPage() int {
	// 优先取用户请求的 page
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if page <= 0 {
		page = 1
	}
	// TotalPage 等于 0 ，意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}
	// 请求页数大于总页数，返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page
}

//拼接分页
func (p Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)
}

//getNextPageURL 返回下一页的链接
func (p Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

// getPrevPageURL 返回下一页的链接
func (p Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}
