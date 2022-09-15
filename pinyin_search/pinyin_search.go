// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pinyin_search

import (
	"bytes"
	"sort"
	"strings"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

func GeneralizeQuery(q string) string {
	var buf bytes.Buffer
	for _, r := range q {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			continue
		} else {
			buf.WriteRune(unicode.ToLower(r))
		}
	}
	return buf.String()
}

func strSliceUniq(slice []string) []string {
	length := len(slice)
	if length == 0 {
		return nil
	} else if length == 1 {
		return slice
	} else if length == 2 {
		if slice[0] == slice[1] {
			return slice[:1]
		} else {
			return slice
		}
	} else {
		return strSliceUniqAux(slice)
	}
}

func strSliceUniqAux(in []string) []string {
	// In-place deduplicate
	sort.Strings(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		in[j] = in[i]
	}
	return in[:j+1]
}

type Blocks []block

var pinyinArgs pinyin.Args

func init() {
	pinyinArgs = pinyin.NewArgs()
	pinyinArgs.Heteronym = true
}

func Split(str string) Blocks {
	if str == "" {
		return nil
	}
	var result Blocks
	var buf bytes.Buffer

	for _, r := range str {
		pys := pinyin.SinglePinyin(r, pinyinArgs)
		pys = strSliceUniq(pys)
		if len(pys) > 0 {
			// 是汉字
			if buf.Len() > 0 {
				result = append(result, commonBlock(buf.String()))
				buf.Reset()
			}

			result = append(result, zhBlock{
				zh:  r,
				pys: pys,
			})
		} else if unicode.IsSpace(r) || unicode.IsPunct(r) {
			if buf.Len() > 0 {
				result = append(result, commonBlock(buf.String()))
				buf.Reset()
			}
		} else {
			r = unicode.ToLower(r)
			buf.WriteRune(r)
		}
	}
	if buf.Len() > 0 {
		result = append(result, commonBlock(buf.String()))
	}
	return result
}

type block interface {
	isBlock()
}

type zhBlock struct {
	zh  rune
	pys []string
}

func (zb zhBlock) isBlock() {
}

func (zb zhBlock) getList() []string {
	var result []string
	result = append(result, string(zb.zh))
	for _, py := range zb.pys {
		result = append(result, getPyList(py)...)
	}
	return result
}

func getPyList(py string) []string {
	// zhong -> zhong,zhon,zho,zh,z
	if len(py) == 0 {
		return nil
	}
	result := make([]string, len(py))
	for i := 0; i < len(py); i++ {
		result[i] = string(py[:len(py)-i])
	}
	return result
}

type commonBlock string

func (c commonBlock) isBlock() {
}

func (blocks Blocks) Match(query string) bool {
	if query == "" {
		return false
	}
	for i := 0; i < len(blocks); i++ {
		if matchBegin(blocks[i:], query) {
			return true
		}
	}
	return false
}

func matchBegin(blocks []block, query string) (result bool) {
	//log.Printf("start match %#v %#v\n", blocks, query)
	//defer func() {
	//	log.Printf("end match %#v %#v, result: %v\n", blocks, query, result)
	//}()
	if query == "" {
		return false
	}

	qIdx := 0
	blocksIdx := 0
	for {
		if blocksIdx >= len(blocks) {
			break
		}
		if qIdx >= len(query) {
			break
		}

		b := blocks[blocksIdx]
		switch block := b.(type) {
		case commonBlock:
			bbs := string(block)
			query0 := query[qIdx:]
			// commonBlock 匹配条件： 全部，或因为 query0 不足够，只能匹配部分。
			isMatch, end, n := matchAux(bbs, query0)
			if isMatch {
				if end {
					return true
				} else {
					qIdx += n
					blocksIdx++
				}

			} else {
				return false
			}

		case zhBlock:
			list := block.getList()
			query0 := query[qIdx:]
			for _, value := range list {
				// compare value and query0
				isMatch, end, n := matchAux(value, query0)
				if isMatch {
					if end {
						return true
					} else {
						if blocksIdx+1 < len(blocks) {
							if matchBegin(blocks[blocksIdx+1:], query0[n:]) {
								return true
							}
						} else {
							// blocks end
							return false
						}
					}
				}
			}
			return false
		}
	}
	return false
}

// 当 isMatch 为 true 时， end 和 n 才有意义。
// 当 end 有意义时，如果 end 为 false， n 才有意义。
func matchAux(target, query string) (isMatch bool, end bool, n int) {
	if len(query) == len(target) {
		if query == target {
			isMatch = true
			end = true
		} else {
			isMatch = false
		}

	} else if len(query) < len(target) {
		// target: dong, query: do
		if strings.HasPrefix(target, query) {
			isMatch = true
			end = true
		} else {
			isMatch = false
		}
	} else {
		// target: dong, query: dongx
		if strings.HasPrefix(query, target) {
			isMatch = true
			n = len(target)
		} else {
			isMatch = false
		}
	}
	//log.Printf("isMatch: %v, end: %v, n: %v\n", isMatch, end, n)
	return
}
