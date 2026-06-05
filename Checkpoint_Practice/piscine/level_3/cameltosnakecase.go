package level_3

func CamelToSnakeCase(s string )string{
	//checkif  empty string
	if s == ""{
		return ""
	}
	 
	for i := 0; i< len(s);i++ {
		c := s[i]
		//check if only letters
		if !(c >= 'a' && c <= 'z') &&  !(c >= 'A' && c <= 'Z'){
			return s
		}
		//check if it ends in caps
		if i ==len(s)-1 && c >='A' && c <='Z'{
			return s
		}
		//No consecutive uppercase letters
		if i >0 && (c >='A' &&c<='Z') && (s[i-1]>='A' && s[i-1] <='Z'){
			return s
		} 
	}
	
	
	


	out := []byte{}

	//convert to snake case
	for  i:=0; i<len(s);i++{
		c:= s[i]
		if i>0 && c >='A' && c<='Z'{
			out = append(out,'_')
		}
		out = append(out,c)
	}
	return string(out)
}