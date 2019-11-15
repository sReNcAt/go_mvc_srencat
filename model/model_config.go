package mvc_model

func dbconn()(string,string){
	return "mysql", "root:passwd@tcp(adress:3306)/database"
}