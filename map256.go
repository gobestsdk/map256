package map256



type Map256  struct{
	Nodes [256]*Node
}
func (this *Map256)String()string{
	str:=""
	for k,_:=range this.Nodes{

		var key_node []byte=make([]byte,1)
		key_node[0]=byte(k)
		str+=this.Nodes[k].String(key_node)
	}
	return str
}
func (this *Map256)Put(key []byte,data interface{}){
	l:=len(key)
	if(l<1){
		return
	}
	nodekey :=key[0]

	if(l>1){
		this.Nodes[nodekey].Put(key[1:],data)
	}else {
		if(this.Nodes[nodekey]==nil){
			this.Nodes[nodekey]=&Node{}
		}
		this.Nodes[nodekey].Data=data
	}
}
func (this *Map256)Exist(key []byte)bool{
	return (this.Get(key)!=nil)
}
func (this *Map256)Get(key []byte)interface{}{

	l:=len(key)
	if(l<1){
		return nil
	}
	nodekey :=key[0]

	if(l>1){
		return this.Nodes[nodekey].Get(key[1:])
	}else {
		return this.Nodes[nodekey].Data
	}
}
func (this *Map256)Del(key []byte){
	this.Put(key,nil)
}