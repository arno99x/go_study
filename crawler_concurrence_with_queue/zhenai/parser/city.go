package parser

import (
	"go_study/crawler_concurrence_with_queue/engine"
	"regexp"
)

//<div class="content">
//	<table>
//		<tbody>
//			<tr><th><a href="http://album.zhenai.com/u/1019887461" target="_blank">曼君Lynn</a></th></tr>
//			<tr><td width="180"><span class="grayL">性别：</span>女士</td> <td><span class="grayL">居住地：</span>安徽蚌埠</td></tr>
//			<tr><td width="180"><span class="grayL">年龄：</span>34</td> <td><span class="grayL">学&nbsp;&nbsp;&nbsp;历：</span>大学本科</td> <!----></tr>
//			<tr><td width="180"><span class="grayL">婚况：</span>离异</td> <td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>160</td></tr>
//			</tbody></table> <div class="introduce">我希望我可以有幸遇到的另一半，他爱我欣赏我，他能认识到真实的自己，我们能够滋养彼此心灵，给彼此供能。如果能够引领我，我发自内心的崇拜欣赏他，听他的话，那就太好了！<br>我爱学习善于学习，很快我就可以引领你啦?<br>我希望我的他是个有信仰的人。<br>希望你有不错的的经济条件，这样我们可以过的轻松舒服一些</div></div>
const userRe = `<span class="grayL">性别：</span>([^<]+)</td>`
const userInfoRe = `<a href="http://album.zhenai.com/u/[^"]+" target="_blank"><img src="https://photo.zastatic.com/images/photo/[^>]+></a></div> <div class="content"><table><tbody><tr><th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">[^<]+</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`

func ParserCity(content []byte) engine.ParserResult {
	r := regexp.MustCompile(userInfoRe)
	userList := r.FindAllSubmatch(content, -1)

	result := engine.ParserResult{}
	//limit := 0
	for _, item := range userList {
		//limit++
		//if limit > 3 {
		//	break
		//}
		result.Items = append(result.Items, string(item[1])+";"+string(item[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(item[1]), ParserFun: func(contents []byte) engine.ParserResult {
			gender := item[2]
			var paramMap = map[string]string{"gender": string(gender)}
			return ParserProfile(contents, paramMap)
		}})
	}

	return result
}
