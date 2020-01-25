package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type help struct {
	program  string
	command  string
	shortcut string
	desc     string
}

var contents = []help{
	{
		program:  "vim",
		command:  "horizontal split",
		shortcut: ":sp",
		desc:     "move windows: ctrl-w, vertical split: vsp",
	},
	{
		program:  "vim",
		command:  "vertical split",
		shortcut: ":vsp",
		desc:     "move windows: ctrl-w, horizontal split: sp",
	},
	{
		program:  "vim",
		command:  "focus window left",
		shortcut: "ctrl-w h",
		desc:     "focus window right: ctrl-w l",
	},
	{
		program:  "vim",
		command:  "focus window right",
		shortcut: "ctrl-w l",
		desc:     "focus window right: ctrl-w h",
	},
	{
		program:  "vim",
		command:  "copy / yank",
		shortcut: "y",
		desc:     "paste with p",
	},
	{
		program:  "vim",
		command:  "paste",
		shortcut: "p",
		desc:     "copy with y (yank)",
	},
	{
		program:  "vim",
		command:  "go back",
		shortcut: "Ctrl-O",
	},
	{
		program:  "vim",
		command:  "go forward",
		shortcut: "Ctrl-I",
	},
	{
		program:  "vim",
		command:  "find",
		shortcut: "/<text>",
		desc:     "n for next match, N for previous match",
	},
	{
		program:  "vim",
		command:  "replace in visual selection",
		shortcut: ":s/<find>/replace/g",
		desc:     "use /gc for confirm",
	},
	{
		program:  "vim",
		command:  "replace in file",
		shortcut: ":%s/<find>/replace/g",
		desc:     "use /gc for confirm",
	},
	{
		program:  "vim",
		command:  "uppercase",
		shortcut: "gU<motion>",
		desc:     "e.g. gUw word, gUU for line, ~ cycles case on a character",
	},
	{
		program:  "vim",
		command:  "lowercase",
		shortcut: "gu<motion>",
		desc:     "e.g. guw (word), guu (line), ~ cycles case on a character",
	},
	{
		program:  "vim",
		command:  "page up",
		shortcut: "Ctrl-u",
		desc:     "Ctrl-b for full screen back, Ctrl-d for down",
	},
	{
		program:  "vim",
		command:  "page down",
		shortcut: "Ctrl-d",
		desc:     "Ctrl-f for full screen down, Ctrl-u for up",
	},
	{
		program:  "vim: NERDCommenter",
		command:  "toggle comments",
		shortcut: "\\c <space>",
		desc:     "\\ is the default <leader> key",
	},
	{
		program:  "tmux",
		command:  "new window",
		shortcut: "ctrl-b c",
	},
	{
		program:  "tmux",
		command:  "create window",
		shortcut: "ctrl-b c",
	},
	{
		program:  "tmux",
		command:  "select window",
		shortcut: "ctrl-b w",
	},
	{
		program:  "tmux",
		command:  "next window",
		shortcut: "ctrl-b n",
	},
	{
		program:  "tmux",
		command:  "kill window",
		shortcut: "ctrl-b x",
	},
	{
		program:  "tmux",
		command:  "exit window",
		shortcut: "ctrl-b x",
	},
	{
		program:  "tmux",
		command:  "detach",
		shortcut: "ctrl-b d",
	},
	{
		program:  "tmux",
		command:  "copy",
		shortcut: "ctrl-b [",
		desc:     "enter mode: ctrl-b [, use arrow keys, copy: ctrl-space, paste: ctrl-b ]",
	},
	{
		program:  "tmux",
		command:  "paste",
		shortcut: "ctrl-b ]",
		desc:     "enter mode: ctrl-b [, use arrow keys, copy: ctrl-space, paste: ctrl-b ]",
	},
}

var noise = map[string]struct{}{
	"the": struct{}{},
	"and": struct{}{},
	"of":  struct{}{},
}

func normaliseWords(s string) (words []string) {
	ww := strings.Split(strings.ToLower(s), " ")
	for i := 0; i < len(ww); i++ {
		if _, isNoisy := noise[ww[i]]; isNoisy {
			continue
		}
		words = append(words, ww[i])
	}
	return
}

func wordMatcher(q string, h []help) (op []int) {
	op = make([]int, len(h))
	if q == "" {
		return
	}
	qw := normaliseWords(q)

	// If the first word is a program, then filter the results.
	// So that only that program is included.
	programNames := map[string]struct{}{}
	for _, hh := range h {
		programNames[hh.program] = struct{}{}
	}
	_, filterByProgram := programNames[qw[0]]

	for i, hh := range h {
		if filterByProgram && strings.ToLower(hh.program) != qw[0] {
			continue
		}
		hw := normaliseWords(hh.command)
		for _, w := range qw {
			for _, hw := range hw {
				if strings.Index(hw, w) >= 0 {
					op[i]++
				}
			}
		}
	}
	return
}

func initialMatcher(q string, h []help) (op []int) {
	op = make([]int, len(h))
	for i, hh := range h {
		// If there's no query.
		if q == "" {
			continue
		}
		// If the program doesn't match.
		if q[0] != hh.program[0] {
			continue
		}
		// If there are more characters in the initials than words in the command.
		cw := strings.Split(hh.command, " ")
		if len(cw) < len(q)-1 {
			continue
		}
		// If the query matches, add one.
		for j, w := range q[1:] {
			if rune(cw[j][0]) == w {
				op[i]++
			}
		}
	}
	return
}

type helpScore struct {
	index, score int
}

func main() {
	q := strings.Join(os.Args[1:], " ")

	// Score the results.
	scores := initialMatcher(q, contents)
	for i, s := range wordMatcher(q, contents) {
		scores[i] += s
	}

	// Filter and collect the help.
	var helpScores []helpScore
	for i, s := range scores {
		if s > 0 {
			helpScores = append(helpScores, helpScore{index: i, score: s})
		}
	}

	// Sort the results.
	sort.Slice(helpScores, func(i, j int) bool {
		return helpScores[i].score > helpScores[j].score
	})

	// Output.
	if len(helpScores) == 0 {
		fmt.Printf("no results for \"%v\"\n", q)
	}
	for _, hs := range helpScores {
		h := contents[hs.index]
		fmt.Printf("%v: %v - %v %v\n", h.program, h.command, h.shortcut, surround("(", h.desc, ")"))
	}
}

func surround(prefix, text, suffix string) string {
	if text == "" {
		return text
	}
	return prefix + text + suffix
}
