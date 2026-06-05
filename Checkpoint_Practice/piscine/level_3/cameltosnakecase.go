package level_3

func CamelToSnakeCase(s string )string{
	//checkif  empty string
	if s == ""{
		return 
	}
	//check if only letters 
	for _, ch := range s {
		if !(ch >= 'a' && ch <= 'z') && (ch >= 'A' && ch <= 'Z'){
			return s
		}
		
	}
	//check if it ends in caps
	if s[len(s)-1] >='A' && s[len(s)-1]<='Z'{
		return s
	}
	out := make([]byte, len(s))

	//convert to snake case
	for  i:=0; i<len(s);i++{
		c:= s[i]
		if c >='A' && c<='Z'{
			out = append(out,'-')
			out = append(out,c)
		}
		out = append(out,c)
	}
}