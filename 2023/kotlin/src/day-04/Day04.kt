fun main() {

    fun part1(input: List<String>): Int {

        var result = 0

        for (card in input) {

            val winningNumbers = parseWinningNumbers(card)
            val cardNumbers = parseCardNumbers(card)

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

        // assign each card a count of 1
        val cardCounts = mutableMapOf<Int, Int>()
        input.forEachIndexed { index, _ -> cardCounts[index] = 1 }

        // iterate over each card
        for ((index, s) in input.withIndex()) {

            val winningNumbers = parseWinningNumbers(s)
            val cardNumbers = parseCardNumbers(s)

            val wins = cardNumbers.filter { it in winningNumbers }.size

            repeat(wins) {

                val nextIndex = index + (it + 1)
                val currentCardWins = cardCounts[index]!!

                cardCounts[nextIndex] = cardCounts[nextIndex]!! + currentCardWins
            }
        }

        return cardCounts.map { it.value }.sum()
    }

    val input = readInput("day-04/input")
    part1(input).println()
    part2(input).println()
}

fun parseWinningNumbers(card: String): List<Int> {
    return card
        .split(":")[1]
        .split("|")[0]
        .split(" ")
        .filter { it.isNotBlank() }
        .map { it.toInt() }
}

fun parseCardNumbers(card: String): List<Int> {
    return card
        .split("|")[1]
        .split(" ")
        .filter { it.isNotBlank() }
        .map { it.toInt() }
}
