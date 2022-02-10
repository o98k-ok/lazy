-- For the latest version:
-- https://github.com/vitorgalvao/custom-alfred-iterm-scripts

-- Set this property to true to always open in a new window
property open_in_new_window : false

-- Handlers
on new_window()
	tell application "iTerm" to create window with default profile
end new_window

on new_tab()
	tell application "iTerm" to tell the first window to create tab with default profile
end new_tab

on call_forward()
	tell application "iTerm" to activate
end call_forward

on is_running()
	application "iTerm" is running
end is_running

on has_windows()
	if not is_running() then return false
	if windows of application "iTerm" is {} then return false
	true
end has_windows

on send_text(custom_text)
	tell application "iTerm" to tell the first window to tell current session to write text custom_text
end send_text

on split(query, delim)
	set oldDelimiters to AppleScript's text item delimiters --记录开始的去限器
    set AppleScript's text item delimiters to delim --设置分隔符
    set str2Arr to every text item of query -- 分割
    set AppleScript's text item delimiters to oldDelimiters -- 恢复原来的去限器
    return str2Arr
end split

-- Main
on alfred_script(query)

	if has_windows() then
		if open_in_new_window then
			new_window()
		else
			new_tab()
		end if
	else
		-- If iTerm is not running and we tell it to create a new window, we get two
		-- One from opening the application, and the other from the command
		if is_running() then
			new_window()
		else
			call_forward()
		end if
	end if

	-- Make sure a window exists before we continue, or the write may fail
	repeat until has_windows()
		delay 0.01
	end repeat

    -- parse args and set iterms tab name
    set arr to split(query, "#")
    set title to 1st item of arr

    send_text("tabset "&title)


    set args to 2nd item of arr
    set others to split(args, ";")

    call_forward()

    repeat with aword in others
        send_text(aword)
        delay 0.8
    end repeat
	call_forward()
end alfred_script