package router

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var PASSGEN fiber.Router

const (
	lowerCase   = "abcdefghijkmnopqrstuvwxyz"
	upperCase   = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	numbers     = "1234567890"
	specialChar = "!@#$%&;:<>?/*()"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	errLowerCaseLimit   = errors.New("lower case chars should not be more than 10 and less than 3")
	errUpperCaseLimit   = errors.New("upper case chars should not be more than 10 and less than 3")
	errNumbersLimit     = errors.New("numbers should not be more than 5 and less than 2")
	errSpecialCharLimit = errors.New("special chars should not be more than 5 and less than 2")
)

func SetupPassGenRoutes() {
	PASSGEN.Get("", GetPassword)
}

func GetPassword(c *fiber.Ctx) error {
	lowerCaseStr := c.Query("lowerCase", "3")
	upperCaseStr := c.Query("upperCase", "3")
	numStr := c.Query("numbers", "2")
	specialChairStr := c.Query("specialChair", "2")

	lowerC, upperC, num, specialC, err := convertStrToNumbers(lowerCaseStr, upperCaseStr, numStr, specialChairStr)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := validateCountChars(lowerC, upperC, num, specialC); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	pass := generatePassword(lowerC, upperC, num, specialC)

	return c.Status(200).JSON(pass)
}

func randString(n int8, s string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = s[rand.Intn(len(s))]
	}
	return string(b)
}

func shuffle(src []string) string {
	final := make([]string, len(src))
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return strings.Join(final, "")
}

func generatePassword(lowerC int8, upperC int8, num int8, specialC int8) string {
	l := randString(lowerC, lowerCase)
	u := randString(upperC, upperCase)
	n := randString(num, numbers)
	s := randString(specialC, specialChar)

	pass := l + u + n + s
	splitPass := strings.Split(pass, "")

	res := shuffle(splitPass)

	return res
}

func validateCountChars(lowerC int8, upperC int8, num int8, specialC int8) error {
	if lowerC > 10 || lowerC < 3 {
		return errLowerCaseLimit
	} else if upperC > 10 || upperC < 3 {
		return errUpperCaseLimit
	} else if num > 5 || num < 2 {
		return errNumbersLimit
	} else if specialC > 5 || specialC < 2 {
		return errSpecialCharLimit
	}
	return nil
}

func convertStrToNumbers(lowerC string, upperC string, num string, specialC string) (int8, int8, int8, int8, error) {
	l, err := strconv.Atoi(lowerC)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	u, err := strconv.Atoi(upperC)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	s, err := strconv.Atoi(specialC)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return int8(l), int8(u), int8(n), int8(s), nil
}
