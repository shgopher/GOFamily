package main

var a = make([]map[string]interface{}, 10)

func add(id int) {
	t := make(map[string]interface{}, 0)
	t["M"] = 1
	a[id] = t
}

func changeData(id int) {
	t := make(map[string]interface{}, 0)
	t["N"] = 1
	a[id] = t
}
