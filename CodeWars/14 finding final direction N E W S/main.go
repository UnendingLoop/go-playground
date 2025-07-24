package kata

var n, s, e, w string = "NORTH", "SOUTH", "EAST", "WEST"

func DirReduc(arr []string) []string {
	//["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"]
	for i := 0; i < len(arr)-1; i++ {
		if (arr[i] == n && arr[i+1] == s || arr[i] == s && arr[i+1] == n) ||
			(arr[i] == e && arr[i+1] == w || arr[i] == w && arr[i+1] == e) {
			switch {
			case len(arr) == 2:
				return []string{}
			case i == len(arr)-2:
				return DirReduc(arr[:i])
			default:
				arr = append(arr[:i], arr[i+2:]...)
				return DirReduc(arr)
			}
		}
	}
	return arr
}
