fun main() {

    fun part1(input: List<String>): Int {

        var result = 0
        for (line in input) {

            val firstDigit = line.find { it.isDigit() }
            val lastDigit = line.findLast { it.isDigit() }

            if (firstDigit != null && lastDigit != null) {
                result += "$firstDigit$lastDigit".toInt()
            }
        }

        return result
    }

    fun part2(input: List<String>): Int {
        return input.size
    }

    val input = readInput("day-01/input")
    part1(input).println()
    part2(input).println()
}
