# MGLang

[![License](https://img.shields.io/badge/license-MIT-5351FB)](LICENSE)

Uma biblioteca simples e leve em Go para **gerenciamento de traduções (i18n)**, com suporte a **detecção automática do idioma do sistema** e **formatação dinâmica de strings**.

Ideal para aplicações desktop, CLI ou bibliotecas que precisam de internacionalização sem dependências pesadas.

---

## ✨ Recursos

* 🌍 Detecção automática do idioma do sistema
* 🗂️ Registro simples de traduções por chave
* 🔁 Fallback automático quando a tradução não existe
* 🧩 Suporte a para textos dinâmicos
* 🚀 API minimalista e fácil de integrar

---

## 📦 Instalação

```bash
go get github.com/profmugomes/mglang
```

---

## 🚀 Uso básico

### Definindo traduções

```go
import "github.com/profmugomes/mglang"

func main() {
	mglang.Set("hello", "Olá")
	mglang.Set("Welcome %s", "Bem-vindo, %s!")
}
```

---

### Obtendo tradução simples

```go
text := mglang.T("hello")
// Resultado: Olá
```

---

### Tradução com parâmetros

```go
text := mglang.T("Welcome %s", "Murilo")
// Resultado: Bem-vindo, Murilo!
```

---

### Fallback automático

```go
text := mglang.T("not_exists")
// Resultado: not_exists
```

---

## 🌐 Detecção de idioma

```go
lang := mglang.GetLang()
fmt.Println(lang) // ex: pt, en, es
```

A função detecta o idioma do sistema operacional e retorna apenas o **código base** (`pt`, `en`, `es`, etc).

---

## 🧠 Estrutura interna

* `Set(key, value)` → registra tradução
* `T(key, args...)` → retorna texto traduzido
* `GetLang()` → detecta idioma do sistema

---

## 🧩 Compatibilidade

* Go 1.27.2+

---

## 👤 Autor

**Murilo Gomes**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2025-2026 Murilo Gomes <profmugomes.com.br>

Licensed under the [MIT](https://github.com/profmugomes/mglang/blob/main/LICENSE) license.

All contributions to the MGLang are subject to this license.
