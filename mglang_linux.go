// Copyright (C) 2025-2026 Murilo Gomes <profmugomes.com.br>
// SPDX-License-Identifier: MIT

// Site: https://www.profmugomes.com.br

package mglang

import (
	"os"
	"strings"
)

type linuxProvider struct{}

func init() {
	Current = linuxProvider{}
}

func (linuxProvider) GetSystemLanguage() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		lang = os.Getenv("LC_MESSAGES")
	}
	if lang == "" {
		return "en" // fallback padrão
	}

	// Exemplo: "pt_BR.UTF-8" → "pt"
	parts := strings.Split(lang, ".")
	base := parts[0]
	baseParts := strings.Split(base, "_")
	return strings.ToLower(baseParts[0])
}
