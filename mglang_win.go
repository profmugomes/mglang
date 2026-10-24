// Copyright (C) 2025-2026 Murilo Gomes <profmugomes.com.br>
// SPDX-License-Identifier: MIT

// Site: https://www.profmugomes.com.br

//go:build windows

package mglang

import (
	"strings"

	"golang.org/x/sys/windows/registry"
)

type winProvider struct{}

func init() {
	Current = winProvider{}
}

func (winProvider) GetSystemLanguage() string {
	k, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		`Control Panel\International`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return "en"
	}
	defer k.Close()

	s, _, err := k.GetStringValue("LocaleName")
	if err != nil || len(s) < 2 {
		return "en"
	}

	return strings.ToLower(s[:2])
}
