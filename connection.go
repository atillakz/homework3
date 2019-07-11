package main

import "fmt"

func main() {

	con := ConnectionPool{}

	con.init(5)


	userid1 := "11111111111111111"
	userid2 := "222"
	userid3 := "333"

	con.getConnection(userid1)

	con.getConnection(userid2)

	con.getConnection(userid3)

	fmt.Println(con)

	con.returnConnection(userid1)

	fmt.Println(con)



}

type ConnectionPool struct {
	connections map[Connection]string
}

type Connection struct {
	id      int
	timeout string
}

func (c *ConnectionPool) getConnection(userid string) Connection {

	newConnection := Connection{}


	for key, value := range c.connections{

		if (value == "AB"){

			c.connections[key] = userid

			newConnection = key
			break
		}


	}

return newConnection
}

func (c *ConnectionPool) returnConnection(userId string) {

	for key, value := range c.connections {

		if (value == userId){

			c.connections[key] = "AB"

			break
		}
	}

}

func (c *ConnectionPool) init(size int) {


	m := make(map[Connection]string)

	for i := 1; i <= size; i++ {

		a := Connection{i,"30min"}

		m[a] = "AB"

	}

	c.connections = m
}