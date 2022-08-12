import mysql.connector as mysql

HOST = "ip" 
DATABASE = "Users"
USER = "user"
PASSWORD = "Password"
# connect to MySQL server
db_connection = mysql.connect(host=HOST, database=DATABASE, user=USER, password=PASSWORD)
print("Connected to:", db_connection.get_server_info())

cur = db_connection.cursor()  
sql = "INSERT INTO userinfo (username, password,email) VALUES (%s, %s, %s)" 
val = ("Lee","12345666","Lee@gmail.com")  


try:  
    cur.execute(sql,val)  
    db_connection.commit()  
      
except:  
    db_connection.rollback()  
  
print(cur.rowcount,"record inserted!")  
db_connection.close()