// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pinyin_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeneralize(t *testing.T) {
	ret := GeneralizeQuery("Firefox - Browser")
	assert.Equal(t, "firefoxbrowser", ret)
}

func Test_strSliceUniq(t *testing.T) {
	ret := strSliceUniq(nil)
	require.Nil(t, ret)

	ret = strSliceUniq([]string{"a"})
	assert.Equal(t, []string{"a"}, ret)

	ret = strSliceUniq([]string{"a", "a"})
	assert.Equal(t, []string{"a"}, ret)

	ret = strSliceUniq([]string{"a", "b"})
	assert.Equal(t, []string{"a", "b"}, ret)

	ret = strSliceUniq([]string{"b", "c", "a", "b", "c"})
	assert.Equal(t, []string{"a", "b", "c"}, ret)
}

func Test_getPyList(t *testing.T) {
	ret := getPyList("zhong")
	assert.Equal(t, []string{"zhong", "zhon", "zho", "zh", "z"}, ret)

	ret = getPyList("z")
	assert.Equal(t, []string{"z"}, ret)

	ret = getPyList("")
	require.Nil(t, ret)
}

func Test_matchAux(t *testing.T) {
	isMatch, end, n := matchAux("dong", "dong")
	assert.True(t, isMatch)
	assert.True(t, end)
	_ = n

	isMatch, end, _ = matchAux("dong", "don")
	assert.True(t, isMatch)
	assert.True(t, end)

	isMatch, end, n = matchAux("dong", "dongsh")
	assert.True(t, isMatch)
	assert.False(t, end)
	assert.Equal(t, 4, n)

	isMatch, _, _ = matchAux("dong", "sh")
	assert.False(t, isMatch)
}

func TestSplit(t *testing.T) {
	ret := Split("")
	require.Nil(t, ret)

	ret = Split("中文汉字")
	assert.Equal(t, Blocks{
		zhBlock{zh: '中', pys: []string{"zhong"}},
		zhBlock{zh: '文', pys: []string{"wen"}},
		zhBlock{zh: '汉', pys: []string{"han"}},
		zhBlock{zh: '字', pys: []string{"zi"}}},
		ret)

	ret = Split("的不")
	assert.Equal(t, Blocks{
		zhBlock{zh: '的', pys: []string{"de", "di"}},
		zhBlock{zh: '不', pys: []string{"bu", "fou", "fu"}},
	}, ret)

	ret = Split("  english word web-browser ")
	assert.Equal(t, Blocks{
		commonBlock("english"),
		commonBlock("word"),
		commonBlock("web"),
		commonBlock("browser"),
	}, ret)

	ret = Split(" hello-world! 中文 english  word ")
	assert.Equal(t, Blocks{
		commonBlock("hello"),
		commonBlock("world"),
		zhBlock{zh: '中', pys: []string{"zhong"}},
		zhBlock{zh: '文', pys: []string{"wen"}},
		commonBlock("english"),
		commonBlock("word"),
	}, ret)
}

func Test_matchBegin(t *testing.T) {
	blocks := Split("钟南")
	ret := matchBegin(blocks, "zhongnan")
	assert.True(t, ret)

	ret = matchBegin(blocks, "zn")
	assert.True(t, ret)

	ret = matchBegin(blocks, "z")
	assert.True(t, ret)

	ret = matchBegin(blocks, "n")
	assert.False(t, ret)

	ret = matchBegin(blocks, "zhongna")
	assert.True(t, ret)

	ret = matchBegin(blocks, "zhona")
	assert.True(t, ret)

	ret = matchBegin(blocks, "zhna")
	assert.True(t, ret)

	ret = matchBegin(blocks, "钟nan")
	assert.True(t, ret)

	blocks = Split("open web browser")
	ret = matchBegin(blocks, "open")
	assert.True(t, ret)

	ret = matchBegin(blocks, "op")
	assert.True(t, ret)

	ret = matchBegin(blocks, "openwebbrowser")
	assert.True(t, ret)

	ret = matchBegin(blocks, "openweb")
	assert.True(t, ret)

	ret = matchBegin(blocks, "opweb")
	assert.False(t, ret)
}

func TestMatch(t *testing.T) {
	blocks := Split("李钟南")
	ret := blocks.Match("zhongnan")
	assert.True(t, ret)

	blocks = Split("启动firefox")
	ret = blocks.Match("dofirefox")
	assert.True(t, ret)

	ret = blocks.Match("firefox")
	assert.True(t, ret)

	ret = blocks.Match("fir")
	assert.True(t, ret)

	ret = blocks.Match("firg")
	assert.False(t, ret)
}
