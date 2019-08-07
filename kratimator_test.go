package main

import "testing"

func Test_kratimate_generals(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: " - ", want: " - "},
		{name: " a-123 ", want: " a-123 "},
		{name: " 1 - 2 ", want: " 1 - 2 "},
		{name: "<a href=\"ana-are-mere\">", want: "<a href=\"ana-are-mere\">"},
		{name: ".div(id=\"m-a\")", want: ".div(id=\"m-a\")"},
		{name: " 1 - 2 ", want: " 1 - 2 "},
	}

	for _, tt := range tests {
		t.Run("Generals: " + tt.name, func(t *testing.T) {
			if got := kratimate(tt.name); got != tt.want {
				t.Errorf("kratimate() = >%v<, want >%v<", got, tt.want)
			}
		})
	}
}

func Test_kratimate_standalones(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: " să-l ", want: " să‑l "},
		{name: "să-l ", want: "să‑l "},
		{name: " să-l", want: " să‑l"},
		{name: " să-l.", want: " să‑l."},
		{name: " să-l?", want: " să‑l?"},
		{name: " să-l!!!", want: " să‑l!!!"},
		{name: "	să-l	", want: "	să‑l	"},
		{name: "	să-l,", want: "	să‑l,"},
		{name: " să-lx", want: " să-lx"},
		{name: "[să-l ", want: "[să‑l "},
		{name: "(să-mi ", want: "(să‑mi "},
		{name: "[să-l)", want: "[să‑l)"},
		{name: "îîntr-un ", want: "îîntr-un "},
		{name: ".div(id=\"m-a\") m-a. ma M-a. ma-ma", want: ".div(id=\"m-a\") m‑a. ma M‑a. ma-ma"},
	}

	for _, tt := range tests {
		t.Run("Standalone: " + tt.name, func(t *testing.T) {
			if got := kratimate(tt.name); got != tt.want {
				t.Errorf("kratimate() = >%v<, want >%v<", got, tt.want)
			}
		})
	}
}

func Test_kratimate_terminators(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "suferința-i", want: "suferința‑i"},
		{name: "suferința-ii", want: "suferința-ii"},
		{name: "facă-se", want: "facă‑se"},
		{name: "facă-se)", want: "facă‑se)"},
		{name: "facă-se]", want: "facă‑se]"},
		{name: "facă-se\"", want: "facă-se\""},
		{name: "facă-sei", want: "facă-sei"},
		{name: "dă-mi-l", want: "dă‑mi‑l"},
		{name: "du-te ", want: "du-te "},
		{name: "duce-m-aș", want: "duce‑m‑aș"},
		{name: "mă-ta", want: "mă‑ta"},
		{name: "plecarea-ți", want: "plecarea‑ți"},
		{name: "plecarea-țiț", want: "plecarea-țiț"},
		{name: "-ul", want: "-ul"},
		{name: "ducându-l", want: "ducându‑l"},
		{name: "dă-o", want: "dă‑o"},
		{name: "0-o", want: "0-o"},
		{name: "desculț-o", want: "desculț‑o"},
		{name: "desculț-o!", want: "desculț‑o!"},
		{name: "desculț-o?", want: "desculț‑o?"},
		{name: "desculț-o\"", want: "desculț-o\""},
		{name: "desculț-o'", want: "desculț-o'"},
		{name: "đ-o'", want: "đ-o'"},
		{name: "'-o", want: "'-o"},
		{name: "desculț-o]", want: "desculț‑o]"},
		{name: "desculț-o)", want: "desculț‑o)"},
	}

	for _, tt := range tests {
		t.Run("Terminals: " + tt.name, func(t *testing.T) {
			if got := kratimate(tt.name); got != tt.want {
				t.Errorf("kratimate() = >%v<, want >%v<", got, tt.want)
			}
		})
	}
}

func Test_kratimate_customs(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: " prim-ministru ", want: " prim‑ministru "},
		{name: " prim-ministrului ", want: " prim‑ministrului "},
		{name: " prim-ministruluixxx", want: " prim‑ministruluixxx"},
		{name: "clasa a 2-a", want: "clasa a 2‑a"},
		{name: "clasa a VI-a", want: "clasa a VI‑a"},
		{name: "@22-a", want: "@22-a"},
		{name: "clasa a VI-a", want: "clasa a VI‑a"},
		{name: "clasa a UI-a", want: "clasa a UI-a"},
		{name: "măta", want: "mă‑ta"},
		{name: "măta", want: "mă‑ta"},
	}

	for _, tt := range tests {
		t.Run("Customs: " + tt.name, func(t *testing.T) {
			if got := kratimate(tt.name); got != tt.want {
				t.Errorf("kratimate() = >%v<, want >%v<", got, tt.want)
			}
		})
	}
}