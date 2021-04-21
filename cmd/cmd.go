package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kavemun/toy-robot/robot"
)

const regexPlaceCmd = "PLACE [0-9],[0-9],[A-Z]+"

func Play(filename string) {
	if filename != "" {
		fmt.Printf("File mode: %v \n", filename)
		fileMode(filename)
	} else {
		fmt.Println("CLI Mode")
		cliMode()
	}
}

func parseLinesFromFile(filename string) []string {
	lines, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(lines), "\n")
}

func fileMode(filename string) {
	var new_robot *robot.Robot

	lines := parseLinesFromFile(filename)
	for _, nextCmd := range lines {
		new_robot = ParseCommand(new_robot, nextCmd)
	}

}

func ParseCommand(new_robot *robot.Robot, command string) *robot.Robot {
	var err error
	if regexp.MustCompile(regexPlaceCmd).MatchString(command) {
		// command is PLACE
		args := strings.Split(strings.Split(command, " ")[1], ",")

		pos_x, _ := strconv.Atoi(args[0])
		pos_y, _ := strconv.Atoi(args[1])
		dir := args[2]

		new_robot, err = robot.Place(pos_x, pos_y, dir)
	} else {
		new_robot, err = runCommand(command, new_robot)
	}
	if err != nil {
		fmt.Printf("Error:  %v\n", err)
	}

	return new_robot

}

func cliMode() {
	var new_robot *robot.Robot
	reader := bufio.NewReader(os.Stdin)
	for {
		nextCmd, _ := reader.ReadString('\n')
		nextCmd = strings.TrimSpace(nextCmd)
		new_robot = ParseCommand(new_robot, nextCmd)
	}
}

func runCommand(command string, my_robot *robot.Robot) (*robot.Robot, error) {
	if my_robot == nil {
		return nil, fmt.Errorf("you must first place the robot")
	}
	switch {
	case command == "MOVE":
		return robot.Move(my_robot)
	case command == "LEFT":
		return robot.RotateLeft(my_robot), nil
	case command == "RIGHT":
		return robot.RotateRight(my_robot), nil
	case command == "REPORT":
		fmt.Println(robot.ReportPosition(my_robot))
		return my_robot, nil
	case command == "" || command == "\n":
		return my_robot, nil
	}
	return my_robot, fmt.Errorf("unrecognized Command: %s", command)
}
