// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

import (
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_pulse(t *testing.T) {
	ctx := GetContextForced()
	if ctx == nil {
		return
	}
	defaultSink := getPulseDefaultSink(ctx)
	if defaultSink == nil {
		err := errors.New("failed to get default sink")
		t.Log(err)
		return
	}
	ctx.SetDefaultSink(defaultSink.Name)

	defaultSource := getPulseDefaultSource(ctx)
	if defaultSource == nil {
		err := errors.New("failed to get default source")
		t.Log(err)
		return
	}
	ctx.SetDefaultSource(defaultSource.Name)

	card, err := getCard(ctx, 0)
	require.NoError(t, err)
	if err != nil {
		t.Log(err)
		return
	}
	assert.NotNil(t, card)
}

func getPulseDefaultSink(ctx *Context) (defaultSink *Sink) {
	defaultSinkName := ctx.GetDefaultSink()
	for _, sink := range ctx.GetSinkList() {
		if sink.Name == defaultSinkName {
			defaultSink = sink
			break
		}
	}
	return
}

func getPulseDefaultSource(ctx *Context) (defaultSource *Source) {
	defaultSinkName := ctx.GetDefaultSource()
	for _, source := range ctx.GetSourceList() {
		if source.Name == defaultSinkName {
			defaultSource = source
			break
		}
	}
	return
}

func getCard(ctx *Context, index uint32) (*Card, error) {
	cardIndex, err := ctx.GetCard(index)
	if err != nil {
		err := errors.New("failed to get default source")
		return nil, err
	}
	for _, card := range ctx.GetCardList() {
		if card.Name == cardIndex.Name {
			return card, nil
		}
	}
	return nil, errors.New("failed to get default source")
}
