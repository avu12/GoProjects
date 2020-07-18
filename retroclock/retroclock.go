package retroclock

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func AnimateClock() {

	zero := [5]string{
		" ███ ",
		" █ █ ",
		" █ █ ",
		" █ █ ",
		" ███ ",
	}
	one := [5]string{
		" ██  ",
		"  █  ",
		"  █  ",
		"  █  ",
		" ███ ",
	}
	two := [5]string{
		" ███ ",
		"   █ ",
		" ███ ",
		" █   ",
		" ███ ",
	}
	three := [5]string{
		" ███ ",
		"   █ ",
		" ███ ",
		"   █ ",
		" ███ ",
	}
	four := [5]string{
		" █ █ ",
		" █ █ ",
		" ███ ",
		"   █ ",
		"   █ ",
	}
	five := [5]string{
		" ███ ",
		" █   ",
		" ███ ",
		"   █ ",
		" ███ ",
	}
	six := [5]string{
		" ███ ",
		" █   ",
		" ███ ",
		" █ █ ",
		" ███ ",
	}
	seven := [5]string{
		" ███ ",
		"   █ ",
		"   █ ",
		"   █ ",
		"   █ ",
	}
	eight := [5]string{
		" ███ ",
		" █ █ ",
		" ███ ",
		" █ █ ",
		" ███ ",
	}
	nine := [5]string{
		" ███ ",
		" █ █ ",
		" ███ ",
		"   █ ",
		" ███ ",
	}
	numbArray := [10][5]string{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}
	dots := [5]string{
		"     ",
		"  █  ",
		"     ",
		"  █  ",
		"     ",
	}
	empty := [5]string{
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}
	_ = dots
	_ = empty

	clock := [8][5]string{
		one, two, dots, four, five, dots, seven, eight,
	}

	//fmt.Println(time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	hour := time.Now().Hour()

	clock[0] = numbArray[hour/10]
	clock[1] = numbArray[hour%10]

	minute := time.Now().Minute()

	clock[3] = numbArray[minute/10]
	clock[4] = numbArray[minute%10]

	second := time.Now().Second()

	clock[6] = numbArray[second/10]
	clock[7] = numbArray[second%10]

	if second%2 == 0 {
		clock[2] = dots
		clock[5] = dots
	} else {
		clock[2] = empty
		clock[5] = empty
	}

	for j := 0; j < 5; j++ {
		for i := 0; i < 8; i++ {
			fmt.Print(clock[i][j])
		}
		fmt.Println()
	}
	time.Sleep(time.Second)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
