# Meizar（注意：仅供娱乐）

golang实现抓取妹子图

默认图片来源网站：[http://jandan.net](http://jandan.net)

依赖项目：[https://github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)

## 2016/6/7更新

1. 增加RuleProvider函数，所以，切换Rule时，只需要修改RuleProvider的返回值就可以
2. 增加一个新的Rule-MeiRule，该Rule用来抓取[https://meizi.us/](https://meizi.us/)图片
3. 增加pagesort参数，用来指定页码增加规则， 1为升序，0为降序，默认为0

默认的Rule还是JandanRule， 想要切换到MeiRule的方法是：
修改/rule/ruleProvider.go文件，将RuleProvider方法的返回值指定为MeiRule.
代码：

``` go
func RuleProvider() Rule {
	return &MeiRule{}
}
```

## 如何使用

window用户可以直接下载**win_exe**目录下在的zip文件，解压出一个main.exe文件。

其他平台用户，可以自行下载源码编译。

使用步骤(以windows为例)：

1. 打开命令行定位到编译文件所在目录
2. 输入命令：

> main.exe -dir D:\jandan -start 2009

## 参数说明

1. -dir 可选， 制定图片保存路径，默认保存当前目录/images/下
2. -start 可选，从多少页开始，默认从第2009页开始抓取(注意程序是按照页码倒序抓取的)
3. -cookie 可选，用户使用的cookie， chrome下可以打开[http://jandan.net](http://jandan.net)，按F12键，选择network项查看，这个参数主要在抓取时发生503错误使用的，默认不用指定
4. -pagesort 可选，指定页码增加规则， 1为升序，0为降序，默认为0

## 效果

![](./art/1.png)

^_^ 图片就不展示了， 大家可以自己运行看。

## 扩展

默认图片来源是[煎蛋](http://jandan.net)上的， 但是作为一个___，难道一个煎蛋就满足了吗？不可能！！！！
所以，我们还可以自定义抓取规则，来抓取不同网站的内容。

怎么定制？

只需要实现`Rule`接口的3个方法就可以了。例如煎蛋的抓取规则是：
``` go
// /rule/jandanRule.go
type JandanRule struct{}

func (p *JandanRule) UrlRule() (url string) {
	return "http://jandan.net/ooxx/"
}

func (p *JandanRule) PageRule(currentPage int) (page string) {
	return "page-" + strconv.Itoa(currentPage)
}

func (p *JandanRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}
```
1. 第一个方法返回我们要抓取的url
2. 第二个方法根据当前页返回url后面的页面信息
3. 第三个方法是内容匹配规则， 将匹配到的内容利用f函数返回

还有哪些你喜欢的网站呢？ 赶紧定制规则来尝试抓取吧。 或者提Issue我来添加抓取规则！
