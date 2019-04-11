/**
代理模式：为其它对象提供一种代理 ，以控制对这个对象的坊问。说白了，代理模式就是为了在你坊问对象时加了一层，你可以在这层里做你想做的事情
应用场景：
	(1) 当客户端对象需要访问远程主机中的对象时可以使用远程代理。

   (2) 当需要用一个消耗资源较少的对象来代表一个消耗资源较多的对象，从而降低系统开销、缩短运行时间时可以使用虚拟代理，
		例如一个对象需要很长时间才能完成加载时。

   (3) 当需要为某一个被频繁访问的操作结果提供一个临时存储空间，以供多个客户端共享访问这些结果时可以使用缓冲代理。
		通过使用缓冲代理，系统无须在客户端每一次访问时都重新执行操作，只需直接从临时缓冲区获取操作结果即可。

   (4) 当需要控制对一个对象的访问，为不同用户提供不同级别的访问权限时可以使用保护代理。

   (5) 当需要为一个对象的访问（引用）提供一些额外的操作时可以使用智能引用代理。
*/

package main

import "fmt"

type User struct {
	Name string
}

func (user *User) Eat(food string) {
	fmt.Println(user.Name, " eat ", food)
}

type UserProxy struct {
	User User
}

func (userProxy UserProxy) Eat(food string) {
	fmt.Println("begin to eat")
	userProxy.User.Eat(food)
	fmt.Println("end to eat")
}

func main() {
	userProxy := UserProxy{}
	userProxy.User = User{Name: "Arno"}
	userProxy.Eat("meat")
}
