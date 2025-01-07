package proxies

import (
	"fmt"
	"regexp"
)

// Counter 用于存储各个地区的计数
type Counter struct {
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
}

var counter Counter

// Reset 重置所有计数器为0
func ResetRenameCounter() {
	counter = Counter{}
}

func Rename(name string) string {
	// 香港
	if regexp.MustCompile(`(?i)(hk|港|hongkong|hong kong)`).MatchString(name) {
		counter.hk++
		return fmt.Sprintf("香港%d", counter.hk)
	}
	// 台湾
	if regexp.MustCompile(`(?i)(tw|台|taiwan|tai wen)`).MatchString(name) {
		counter.tw++
		return fmt.Sprintf("台湾%d", counter.tw)
	}
	// 美国
	if regexp.MustCompile(`(?i)(us|美|united states|america)`).MatchString(name) {
		counter.us++
		return fmt.Sprintf("美国%d", counter.us)
	}
	// 新加坡
	if regexp.MustCompile(`(?i)(sg|新|singapore|狮城)`).MatchString(name) {
		counter.sg++
		return fmt.Sprintf("新加坡%d", counter.sg)
	}
	// 日本
	if regexp.MustCompile(`(?i)(jp|日|japan)`).MatchString(name) {
		counter.jp++
		return fmt.Sprintf("日本%d", counter.jp)
	}
	// 英国
	if regexp.MustCompile(`(?i)(uk|英|united kingdom|britain)`).MatchString(name) {
		counter.uk++
		return fmt.Sprintf("英国%d", counter.uk)
	}
	// 加拿大
	if regexp.MustCompile(`(?i)(ca|加|canada)`).MatchString(name) {
		counter.ca++
		return fmt.Sprintf("加拿大%d", counter.ca)
	}
	// 澳大利亚
	if regexp.MustCompile(`(?i)(au|澳|australia)`).MatchString(name) {
		counter.au++
		return fmt.Sprintf("澳大利亚%d", counter.au)
	}
	// 德国
	if regexp.MustCompile(`(?i)(de|德|germany|deutschland)`).MatchString(name) {
		counter.de++
		return fmt.Sprintf("德国%d", counter.de)
	}
	// 法国
	if regexp.MustCompile(`(?i)(fr|法|france)`).MatchString(name) {
		counter.fr++
		return fmt.Sprintf("法国%d", counter.fr)
	}
	// 荷兰
	if regexp.MustCompile(`(?i)(nl|荷|netherlands)`).MatchString(name) {
		counter.nl++
		return fmt.Sprintf("荷兰%d", counter.nl)
	}
	// 俄罗斯
	if regexp.MustCompile(`(?i)(ru|俄|russia)`).MatchString(name) {
		counter.ru++
		return fmt.Sprintf("俄罗斯%d", counter.ru)
	}
	// 匈牙利
	if regexp.MustCompile(`(?i)(hu|匈|hungary)`).MatchString(name) {
		counter.hu++
		return fmt.Sprintf("匈牙利%d", counter.hu)
	}
	// 乌克兰
	if regexp.MustCompile(`(?i)(ua|乌|ukraine)`).MatchString(name) {
		counter.ua++
		return fmt.Sprintf("乌克兰%d", counter.ua)
	}
	// 波兰
	if regexp.MustCompile(`(?i)(pl|波|poland)`).MatchString(name) {
		counter.pl++
		return fmt.Sprintf("波兰%d", counter.pl)
	}
	// 韩国
	if regexp.MustCompile(`(?i)(kr|韩|korea)`).MatchString(name) {
		counter.kr++
		return fmt.Sprintf("韩国%d", counter.kr)
	}
	// 亚太地区
	if regexp.MustCompile(`(?i)(ap|亚太|asia)`).MatchString(name) {
		counter.ap++
		return fmt.Sprintf("亚太地区%d", counter.ap)
	}
	// 伊朗
	if regexp.MustCompile(`(?i)(ir|伊|iran)`).MatchString(name) {
		counter.ir++
		return fmt.Sprintf("伊朗%d", counter.ir)
	}
	// 意大利
	if regexp.MustCompile(`(?i)(it|意|italy)`).MatchString(name) {
		counter.it++
		return fmt.Sprintf("意大利%d", counter.it)
	}
	// 其他
	counter.other++
	return fmt.Sprintf("其他%d-%s", counter.other, name)
}
