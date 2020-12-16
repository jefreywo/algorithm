package emoji

import "regexp"

// https://www.unicode.org/Public/emoji/13.0/emoji-sequences.txt
// https://www.unicode.org/Public/emoji/13.0/emoji-zwj-sequences.txt
// https://apps.timwhitlock.info/emoji/tables/unicode#block-1-emoticons
// 不能通过字符长度过滤emoji表情，可以通过正则匹配识别出emoji表情，使用regexp包
func FilterEmoji(str string) string {
	// 正则匹配
	var emojiRx = regexp.MustCompile(`[\x{1F000}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	return emojiRx.ReplaceAllString(str, ``)
}
