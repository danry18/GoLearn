package gaogo 

import ( 
   "database/sql"  
   "fmt" 
   "runtime"
) 

func main() {

  // mysql扩展 https://github.com/go-sql-driver/mysql  
  mydb, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gogao?charset=utf8")   
  if err != nil { 
       fmt.Println(err)  
  } 
  defer mydb.Close()     

   type User struct {  
      usrid int32    
      usrname string 
      usrsex  int8 
      usrage int8  
  } 
    
  var myuser User       
  rows, err := mydb.Query(`SELECT id,name,sex,age FROM myuser`)
  if err != nil {
        fmt.Println(err) 
   }   
  defer rows.Close() 
 
 for rows.Next() {   
     // 遍历  
        if err := rows.Scan(&myuser.usrid, &myuser.usrname,&mysuer.usersex, &myuser.usrage); err != nil {
            fmt.Println(err)
            continue 
       } 
       fmt.Println(myuser.usrid, myuser.usrame, mysur.usrsex,myuser.usrage)
    }
    if err := rows.Err(); err != nil { 
       fmt.Println(err)
    } 

   // 查询一条记录 
    err = mydb.QueryRow(`SELECT id,name,sex,age WHERE id = ?`, 2).Scan(&myuser.usrid, &myuser.usrname, &myuser.usrsex,&myuser.usrage )
    switch {
    case err == sql.ErrNoRows:
    case err != nil:
            fmt.Println(err)
    } 
   fmt.Println(myuser.usrid, myuser.usrname) 

}
