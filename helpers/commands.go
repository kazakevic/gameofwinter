package helpers

import (
	"kazakevic/gameofwinter/models"
	"regexp"
	"strconv"
)

/*
DetectStartGame detect start command in message
*/
func DetectStartGame(message []byte) bool {
	re := regexp.MustCompile(`(?mi)^(Start)\s(.*)`)
	match := re.MatchString(string(message))

	if match {
		return true
	}
	return false
}

/*
GetUsernameFromMessage parses username from message
*/
func GetUsernameFromMessage(message []byte) string {
	re := regexp.MustCompile(`(?mi)^(Start)\s(.*)`)
	match := re.FindStringSubmatch(string(message))

	if len(match) > 0 {
		playerName := match[2]
		return playerName
	}

	return ""
}

/*
DetectShoot looks for shoot command in message
*/
func DetectShoot(message []byte) bool {
	re := regexp.MustCompile(`(?mi)^(Shoot)\s(\d+)\s(\d+)`)
	match := re.MatchString(string(message))
	if match {
		return true
	}
	return false
}

/*
ShootXY Parses shoot coordinates from shoot command
*/
func ShootXY(message []byte) (x, y int) {

	re := regexp.MustCompile(`(?mi)^(Shoot)\s(\d+)\s(\d+)`)
	match := re.FindStringSubmatch(string(message))

	x, _ = strconv.Atoi(match[2])
	y, _ = strconv.Atoi(match[3])

	return x, y
}

/*
AnnounceWinner - sends message of winner
*/
func AnnounceWinner(player models.Player) string {
	winner := "No winner for now"
	if len(player.Username) > 1 {
		winner = "Now winner is --->" + player.Username + "\n"
	}
	return winner
}
