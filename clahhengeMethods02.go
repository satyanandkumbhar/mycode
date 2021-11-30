package main

import (
	"fmt")

	type Virtmach struct {
		ip string
		hostname string
		diskgb int
		ram int
	}

	func (v *Virtmach)ipset(ip string) {
		v.ip = ip
	}

	func (v *Virtmach)diskexpand(gb int){
		v.diskgb = v.diskgb + gb
	}

	func (v Virtmach)display() {
		fmt.Println("Primary IP Address:", v.ip)
		fmt.Println("Hostname:", v.hostname)
		fmt.Println("Disk GB:", v.diskgb)
		fmt.Println("RAM:", v.ram)
	}

	func main() {
		vm1 := Virtmach{"10.0.0.5", "zebra", 20, 8 }
		vm1.display()                   // show the current state

		    vm1.diskexpand(3)               // increase by 3 GB

		    vm1.ipset("192.168.77.33")      // change the IP address

		    vm1.display()                   // show the changes
	}
