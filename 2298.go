		synonymList := strings.Split(synonym, ",")
		first := string([]byte(synonymList[0])[1:])
		second := string([]byte(synonymList[1])[0 : len(synonymList[1])-1])
