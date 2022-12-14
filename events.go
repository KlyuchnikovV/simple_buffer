package simple_buffer

import (
	"fmt"
	"strings"
)

func (buffer *Buffer) listenEvents() {
	for {
		select {
		case event, ok := <-buffer.KeyEvents:
			if !ok {
				panic("not ok")
			}

			if err := buffer.handleKeyEvent(event); err != nil {
				panic(err)
			}
		case event, ok := <-buffer.MouseEvents:
			if !ok {
				panic("not ok")
			}

			buffer.Selection.SetStartAndEnd(event.Start, event.End)
		}
	}
}

func (buffer *Buffer) handleKeyEvent(event KeyboardEvent) error {
	if event.Meta {
		return buffer.handleMeta(event)
	}
	if event.Alt {
		return buffer.handleAlt(event)
	}

	switch event.Key {
	case "ArrowLeft":
		buffer.CursorLeft(event.Shift)
		return nil
	case "ArrowRight":
		buffer.CursorRight(event.Shift)
		return nil
	case "ArrowUp":
		buffer.CursorUp(event.Shift)
		return nil
	case "ArrowDown":
		buffer.CursorDown(event.Shift)
		return nil
	case "Enter":
		return buffer.Append('\n')
	case "Tab":
		return buffer.Append('\t')
	case "Backspace":
		return buffer.Delete()
	case "Shift":
		return nil
	}

	if len(event.Key) < 3 {
		return buffer.Append(rune(event.Key[0]))
	}

	panic(
		fmt.Sprintf(
			"key not handled key - %s, shift %t, meta %t, option %t, control %t",
			event.Key, event.Shift, event.Meta, event.Alt, event.Ctrl,
		),
	)
}

func (buffer *Buffer) handleMeta(event KeyboardEvent) error {
	switch strings.ToLower(event.Key) {
	case "c":
		buffer.Copy()
	case "v":
		return buffer.Paste()
	case "x":
		return buffer.Cut()
	case "a":
		return buffer.SelectAll()
	case "z":
		if !event.Shift {
			buffer.Undo()
		} else {
			buffer.Redo()
		}
	}

	return nil
}

func (buffer *Buffer) handleAlt(event KeyboardEvent) error {
	switch event.Key {
	// case "ArrowLeft":
	// 	buffer.CursorLeft(event.Shift)
	// 	return nil
	// case "ArrowRight":
	// 	buffer.CursorRight(event.Shift)
	// 	return nil
	// case "ArrowUp":
	// 	buffer.MoveSelectedUp()
	// return nil
	// case "ArrowDown":
	// 	buffer.CursorDown(event.Shift)
	// 	return nil
	}

	return nil
}
