package map256

import (

	"fmt"
)

type Node struct  {
	Data interface{}
        Nodes [256]*Node
}

func (this *Node)String(key_this []byte)string{
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
func (this *Node)Put(key []byte,data interface{}){

	l:=len(key)
	if(l<1){
		return
	}
	nodekey :=key[0]


	if(l>1){
		this.Nodes[nodekey].Put(key,data)
	}else {
		if(this.Nodes[nodekey]==nil){
			this.Nodes[nodekey]=&Node{}
		}
		this.Nodes[nodekey].Data=data

	}
}
func (this *Node)Exist(key []byte)bool{
	return (this.Get(key)!=nil)
}
func (this *Node)Get(key []byte)(interface{}){

	l:=len(key)
	if(l<1){
		return nil
	}
	nodekey :=key[0]


	if(this.Nodes[nodekey]==nil){
		return nil
	}
	if(l>1){
		return this.Nodes[nodekey].Get(key[1:])
	}else {
		return this.Nodes[nodekey].Data
	}
}
func (this *Node)Del(key []byte){
	this.Put(key,nil)
}