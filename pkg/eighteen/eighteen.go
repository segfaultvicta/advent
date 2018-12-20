package eighteen

// Dispatch takes in a day and side and lines of input and calls the correct puzzle function,
// returning the string result of that puzzle function.
func Dispatch(day string, side string, lines []string) string {
	puzzle := day + side
	switch puzzle {
	case "1A":
		return day1sideA(lines)
	case "1B":
		return day1sideB(lines)
	case "2A":
		return day2sideA(lines)
	case "2B":
		return day2sideB(lines)
	case "3A":
		return day3sideA(lines)
	case "3B":
		return day3sideB(lines)
	case "4A":
		return day4sideA(lines)
	case "4B":
		return day4sideB(lines)
	case "5A":
		return day5sideA(lines)
	case "5B":
		return day5sideB(lines)
	case "6A":
		return day6sideA(lines)
	case "6B":
		return day6sideB(lines)
	case "7A":
		return day7sideA(lines)
	case "7B":
		return day7sideB(lines)
	case "8A":
		return day8sideA(lines)
	case "8B":
		return day8sideB(lines)
	case "9A":
		return day9sideA(lines)
	case "9B":
		return day9sideB(lines)
	case "10A":
		return day10sideA(lines)
	case "10B":
		return day10sideB(lines)
	case "11A":
		return day11sideA(lines)
	case "11B":
		return day11sideB(lines)
	case "12A":
		return day12sideA(lines)
	case "12B":
		return day12sideB(lines)
	case "13A":
		return day13sideA(lines)
	case "13B":
		return day13sideB(lines)
	case "14A":
		return day14sideA(lines)
	case "14B":
		return day14sideB(lines)
	case "15A":
		return day15sideA(lines)
	case "15B":
		return day15sideB(lines)
	case "16A":
		return day16sideA(lines)
	case "16B":
		return day16sideB(lines)
	case "17A":
		return day17sideA(lines)
	case "17B":
		return day17sideB(lines)
	case "18A":
		return day18sideA(lines)
	case "18B":
		return day18sideB(lines)
	case "19A":
		return day19sideA(lines)
	case "19B":
		return day19sideB(lines)
	case "20A":
		return day20sideA(lines)
	case "20B":
		return day20sideB(lines)
	case "21A":
		return day21sideA(lines)
	case "21B":
		return day21sideB(lines)
	case "22A":
		return day22sideA(lines)
	case "22B":
		return day22sideB(lines)
	case "23A":
		return day23sideA(lines)
	case "23B":
		return day23sideB(lines)
	case "24A":
		return day24sideA(lines)
	case "24B":
		return day24sideB(lines)
	case "25A":
		return day25sideA(lines)
	case "25B":
		return "happy christmas!"
	default:
		return "unrecognised day/side"
	}
}
