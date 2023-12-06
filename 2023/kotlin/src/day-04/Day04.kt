fun main() {

    fun part1(input: List<String>): Int {

        var result = 0

        for (card in input) {

            val winningNumbers = card
                .split(":")[1]
                .split("|")[0]
                .split(" ")
                .filter { it.isNotBlank() }
                .map { it.toInt() }

            val cardNumbers = card
                .split("|")[1]
                .split(" ")
                .filter { it.isNotBlank() }
                .map { it.toInt() }

            val wins = cardNumbers.filter { it in winningNumbers }.size

            var points = if (wins == 0) 0 else 1
            repeat(wins - 1) {
                points *= 2
            }

            result += points
        }

        return result
    }

    fun part2(input: List<String>): Int {
        return input.size
    }

    val input = readInput("day-04/input")
    part1(input).println()
    part2(input).println()
}
