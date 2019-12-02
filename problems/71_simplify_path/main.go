package main

// code
import "strings"

func simplifyPath(path string) string {
	var stack []string

	for i := 0; i < len(path); i++ {
		if ele := path[i]; ele == '/' {
			if len(stack) == 0 {
				stack = append(stack, "/")
			} else {
				switch stack[len(stack)-1] {
				case "/":
					continue

				case ".":
					stack = stack[0 : len(stack)-1]

				case "..":
					var countSlash int
					for len(stack) > 0 {
						if stack[len(stack)-1] == "/" {
							if countSlash == 0 {
								countSlash++
								stack = stack[0 : len(stack)-1]
							} else {
								break
							}
						} else {
							stack = stack[0 : len(stack)-1]
						}
					}
					if len(stack) == 0 {
						stack = []string{"/"}
					}

				default:
					stack = append(stack, "/")
				}
			}
		} else {
			if len(stack) == 0 {
				stack = []string{string([]byte{ele})}
			} else {
				if stack[len(stack)-1] == "/" {
					stack = append(stack, string([]byte{ele}))
				} else {
					stack[len(stack)-1] += string([]byte{ele})
				}
			}
		}
	}

	if len(stack) > 1 {
		switch stack[len(stack)-1] {
		case "/":
			stack = stack[0 : len(stack)-1]

		case ".":
			stack = stack[0 : len(stack)-2]

		case "..":
			var countSlash int
			for len(stack) > 0 {
				if stack[len(stack)-1] == "/" {
					stack = stack[0 : len(stack)-1]

					if countSlash == 0 {
						countSlash++
					} else {
						break
					}
				} else {
					stack = stack[0 : len(stack)-1]
				}
			}
		}
	}

	if len(stack) == 0 {
		stack = []string{"/"}
	}
	return strings.Join(stack, "")
}
