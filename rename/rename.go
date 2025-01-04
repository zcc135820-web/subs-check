package rename

import (
	"fmt"
	"regexp"
)

// 重命名
var (
	// 香港
	hk int
	// 台湾
	tw int
	// 美国
	us int
	// 新加坡
	sg int
	// 日本
	jp int
	// 英国
	uk int
	// 加拿大
	ca int
	// 澳大利亚
	au int
	// 德国
	de int
	// 法国
	fr int
	// 荷兰
	nl int
	// 俄罗斯
	ru int
	// 匈牙利
	hu int
	// 乌克兰
	ua int
	// 波兰
	pl int
	// 韩国
	kr int
	// 亚太地区
	ap int
	// 伊朗
	ir int
	// 意大利
	it int
	// 其他
	other int
)

func Rename(name string) string {
	// 香港
	if regexp.MustCompile(`(?i)(hk|港|hongkong|hong kong)`).MatchString(name) {
		hk++
		return fmt.Sprintf("香港%d", hk)
	}
	// 台湾
	if regexp.MustCompile(`(?i)(tw|台|taiwan|tai wen)`).MatchString(name) {
		tw++
		return fmt.Sprintf("台湾%d", tw)
	}
	// 美国
	if regexp.MustCompile(`(?i)(us|美|united states|america)`).MatchString(name) {
		us++
		return fmt.Sprintf("美国%d", us)
	}
	// 新加坡
	if regexp.MustCompile(`(?i)(sg|新|singapore|狮城)`).MatchString(name) {
		sg++
		return fmt.Sprintf("新加坡%d", sg)
	}
	// 日本
	if regexp.MustCompile(`(?i)(jp|日|japan)`).MatchString(name) {
		jp++
		return fmt.Sprintf("日本%d", jp)
	}
	// 英国
	if regexp.MustCompile(`(?i)(uk|英|united kingdom|britain)`).MatchString(name) {
		uk++
		return fmt.Sprintf("英国%d", uk)
	}
	// 加拿大
	if regexp.MustCompile(`(?i)(ca|加|canada)`).MatchString(name) {
		ca++
		return fmt.Sprintf("加拿大%d", ca)
	}
	// 澳大利亚
	if regexp.MustCompile(`(?i)(au|澳|australia)`).MatchString(name) {
		au++
		return fmt.Sprintf("澳大利亚%d", au)
	}
	// 德国
	if regexp.MustCompile(`(?i)(de|德|germany|deutschland)`).MatchString(name) {
		de++
		return fmt.Sprintf("德国%d", de)
	}
	// 法国
	if regexp.MustCompile(`(?i)(fr|法|france)`).MatchString(name) {
		fr++
		return fmt.Sprintf("法国%d", fr)
	}
	// 荷兰
	if regexp.MustCompile(`(?i)(nl|荷|netherlands)`).MatchString(name) {
		nl++
		return fmt.Sprintf("荷兰%d", nl)
	}
	// 俄罗斯
	if regexp.MustCompile(`(?i)(ru|俄|russia)`).MatchString(name) {
		ru++
		return fmt.Sprintf("俄罗斯%d", ru)
	}
	// 匈牙利
	if regexp.MustCompile(`(?i)(hu|匈|hungary)`).MatchString(name) {
		hu++
		return fmt.Sprintf("匈牙利%d", hu)
	}
	// 乌克兰
	if regexp.MustCompile(`(?i)(ua|乌|ukraine)`).MatchString(name) {
		ua++
		return fmt.Sprintf("乌克兰%d", ua)
	}
	// 波兰
	if regexp.MustCompile(`(?i)(pl|波|poland)`).MatchString(name) {
		pl++
		return fmt.Sprintf("波兰%d", pl)
	}
	// 韩国
	if regexp.MustCompile(`(?i)(kr|韩|korea)`).MatchString(name) {
		kr++
		return fmt.Sprintf("韩国%d", kr)
	}
	// 亚太地区
	if regexp.MustCompile(`(?i)(ap|亚太|asia)`).MatchString(name) {
		ap++
		return fmt.Sprintf("亚太地区%d", ap)
	}
	// 伊朗
	if regexp.MustCompile(`(?i)(ir|伊|iran)`).MatchString(name) {
		ir++
		return fmt.Sprintf("伊朗%d", ir)
	}
	// 意大利
	if regexp.MustCompile(`(?i)(it|意|italy)`).MatchString(name) {
		it++
		return fmt.Sprintf("意大利%d", it)
	}
	// 其他
	other++
	return fmt.Sprintf("其他%d-%s", other, name)
}
