package map256

import (
	"fmt"

)

type Node256 struct  {
	Data interface{}
        Nodes []*Node256
}

func (this *Node256)String(key_this []byte)string{
	str:=""
	if(this==nil){
		return str;
	}
	if(this.Data!=nil){
		str=string(key_this)+":"+fmt.Sprint(this.Data)+" "
	}
	for k,_:=range this.Nodes{
		var key_node []byte=make([]byte,len(key_this)+1)
		key_node=append(key_node,key_this...)
		key_node=append(key_node,byte(k))
		str+=this.Nodes[k].String(key_node)
	}
	return str
}
func (this *Node256)Len()int{
	l:=0
	for k,_:=range this.Nodes{
		if(this.Nodes[k]!=nil){
			if(this.Nodes[k].Data!=nil){
				l++
			}
			l+=this.Nodes[k].Len()
		}
	}
	return l
}
func (this *Node256)Put(key []byte,data interface{}){

	l:=len(key)
	if(l==0){
		this.Data=data
	}else {
		nodekey :=uint8(key[0])
		if(this==nil) {
			this = &Node256{}
		}
		if(len(this.Nodes)!=256){
			this.Nodes=make([]*Node256,256)
		}

		if(this.Nodes[nodekey]==nil){
			this.Nodes[nodekey]=&Node256{}
		}
		this.Nodes[nodekey].Put(key[1:],data)
	}

}
func (this *Node256)Exist(key []byte)bool{
	return (this.Get(key)!=nil)
}
func (this *Node256)Get(key []byte)(interface{}){

	l:=len(key)
	if(l==0){
		return this.Data
	}else {
		if (l < 1) {
			return nil
		}
		nodekey := key[0]

		if (this.Nodes[nodekey] == nil) {
			return nil
		}
		return this.Nodes[nodekey].Get(key[1:])
	}


}
func (this *Node256)Del(key []byte){
	this.Put(key,nil)
}