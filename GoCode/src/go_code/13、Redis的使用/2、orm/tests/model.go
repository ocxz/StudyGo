package regiorm

type Student struct {
	
	Name string `orm:"name"`
	Age int `orm:"age"`
	Sex string `orm:"sex"`
}