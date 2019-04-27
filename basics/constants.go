package basics

const NAME = "GOD"
const name = "God"

//the function is to escape the unused variable name
func _() {
	_ = name

	//escaping unused variable i
	var i int
	_ = i
}

const (
	ZERO = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
)
