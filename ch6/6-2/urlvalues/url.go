package url 
type Values map[string][]string 
func(v Values)Get(key string)string{
	if vs:=v[key];len(vs)>0{
		return v[key][0]
	}
	return ""
}
func(v Values)Add(m string,n string){
	v[m]=append(v[m], n)
}
func (v Values)Set(m string,n []string){
	v[m]=n
}
