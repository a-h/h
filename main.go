package main

import (
	"fmt"
	"log"
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
		command:  "advance to end of current/next word",
		shortcut: "e",
		desc:     "start of next word: w, start of previous word: B, end of previous word: b",
	},
	{
		program:  "vim",
		command:  "advance to start of next word",
		shortcut: "w",
		desc:     "start of previous word: B, end of previous word: b",
	},
	{
		program:  "vim",
		command:  "home",
		shortcut: "0",
		desc:     "end: $",
	},
	{
		program:  "vim",
		command:  "end",
		shortcut: "$",
		desc:     "home: 0",
	},
	{
		program:  "vim",
		command:  "query history of past commands",
		shortcut: "q:",
		desc:     "a list of previous commands, :q to exit",
	},
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
		shortcut: "ctrl-w h / ctrl-w <-",
		desc:     "focus window right: ctrl-w l",
	},
	{
		program:  "vim",
		command:  "focus window right",
		shortcut: "ctrl-w l / ctrl-w ->",
		desc:     "focus window right: ctrl-w h",
	},
	{
		program:  "vim",
		command:  "delete previous line",
		shortcut: ":-d",
		desc:     "delete next line: :+d, delete to line num: d<line_num>G",
	},
	{
		program:  "vim",
		command:  "delete word",
		shortcut: "dw",
		desc:     "delete inside word: diw, delete to char: dt<char>",
	},
	{
		program:  "vim",
		command:  "delete next line",
		shortcut: ":+d",
		desc:     ":-d to delete previous line, delete to line num: d<line_num>G",
	},
	{
		program:  "vim",
		command:  "delete to line",
		shortcut: "d<line_num>G",
		desc:     "dd: delete current line, :-d to delete previous line",
	},
	{
		program:  "vim",
		command:  "change in word",
		shortcut: "ciw",
		desc:     "ciw: 'w█rd' -> '', cw: 'w█rd' -> 'w'",
	},
	{
		program:  "vim",
		command:  "change word",
		shortcut: "cw",
		desc:     "ciw: 'w█rd' -> '', cw: 'w█rd' -> 'w'",
	},
	{
		program:  "vim",
		command:  "copy / yank",
		shortcut: "y",
		desc:     "paste: p, copy to system clipboard: \"+y, copy file: :%y+",
	},
	{
		program:  "vim",
		command:  "paste",
		shortcut: "p",
		desc:     "copy: y (yank), copy to system clipboard: \"+y, copy file :%y+",
	},
	{
		program:  "vim",
		command:  "copy lines / yank lines",
		shortcut: ":<from_line>,<to_line>y",
		desc:     "paste: p, copy to system clipboard: :<from_line>,<to_line>y+",
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
		program:  "vim",
		command:  "next bracket",
		shortcut: "%",
		desc:     "moves to the previous/next [], {}, ()",
	},
	{
		program:  "vim",
		command:  "previous bracket",
		shortcut: "%",
		desc:     "moves to the previous/next [], {}, ()",
	},
	{
		program:  "vim",
		command:  "search (find)",
		shortcut: "/query",
		desc:     "next: n, move to: enter",
	},
	{
		program:  "vim",
		command:  "regex find/replace execute command",
		shortcut: ":g/regex/<cmd>",
		desc:     "delete lines that contain a or b: :/g/(a|b)/dd",
	},
	{
		program:  "vim",
		command:  "clear search highlight",
		shortcut: ":noh",
		desc:     "",
	},
	{
		program:  "vim: NERDCommenter",
		command:  "toggle comments",
		shortcut: "\\c <space>",
		desc:     "\\ is the default <leader> key",
	},
	{
		program:  "vim: surround",
		command:  "surround with",
		shortcut: "ys<object><text>",
		desc:     "surround word with quotes: ysiw\", surround to end of line with quotes: ys$\"",
	},
	{
		program:  "vim: surround",
		command:  "change surrounding",
		shortcut: "cs<previous><new>",
		desc:     "switch from double to single quotes: cs\"'",
	},
	{
		program:  "vim: surround",
		command:  "delete surrounding",
		shortcut: "ds<surround>",
		desc:     "delete quotes from word: ds\"",
	},
	{
		program:  "vim",
		command:  "new tab",
		shortcut: ":tabedit <filename>",
		desc:     "next tab: gt, previous tab: gT, go to tab 1: 1gt",
	},
	{
		program:  "vim",
		command:  "next tab",
		shortcut: "gt",
		desc:     "new tab: :tabedit <filename>, previous tab: gT, go to tab 1: 1gt",
	},
	{
		program:  "vim",
		command:  "previous tab",
		shortcut: "gT",
		desc:     "new tab: :tabedit filename, next tab: gt, go to tab 1: 1gt",
	},
	{
		program:  "vim",
		command:  "run test",
		shortcut: "<leader>, t, Ctrl-S or :TestSuite",
		desc:     "The leader key is \\ by default, press enter to exit the test results.",
	},
	{
		program:  "vim",
		command:  "move line up",
		shortcut: ":m -1",
		desc:     "move line down: :m +1",
	},
	{
		program:  "vim",
		command:  "move line down",
		shortcut: ":m +1",
		desc:     "move line up: :m -1",
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
		desc:     "enter mode: ctrl-b [, use arrow keys, start copy: ctrl-space, copy to buffer: ctrl-w, paste: ctrl-b ]",
	},
	{
		program:  "tmux",
		command:  "paste",
		shortcut: "ctrl-b ]",
		desc:     "enter mode: ctrl-b [, use arrow keys, start copy: ctrl-space, copy to buffer: ctrl-w, paste: ctrl-b ]",
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

func shortcutMatcher(q string, h []help) (op []int) {
	op = make([]int, len(h))
	if q == "" {
		return
	}

	for i, hh := range h {
		if hh.shortcut == q {
			op[i]++
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

var debug = os.Getenv("H_DEBUG") == "true"

func logf(format string, a ...interface{}) {
	if debug {
		log.Printf(format, a...)
	}
}

func main() {
	q := strings.Join(os.Args[1:], " ")
	logf("query: %v", q)

	// Score the results.
	scores := initialMatcher(q, contents)
	logf("completed initial matcher")
	for i, s := range wordMatcher(q, contents) {
		scores[i] += s
	}
	logf("completed word matcher")
	for i, s := range shortcutMatcher(q, contents) {
		scores[i] += s
	}
	logf("completed shortcut matcher")

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
	logf("sorted results")

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
