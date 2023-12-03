fun main() {

    val nonDigitRegex = "[^0-9]".toRegex()

    fun part1(input: List<String>): Int {

        val redLimit = 12
        val greenLimit = 13
        val blueLimit = 14

        var result = 0
        game@ for ((index, game) in input.withIndex()) {

            val grabs = game.split(":").last().split(";")
            for (grab in grabs) {

                val colors = grab.split(",")
                for (color in colors) {

                    val digit = color.replace(nonDigitRegex, "").toInt()

                    if (color.contains("red") && digit > redLimit) {
                        continue@game
                    } else if (color.contains("green") && digit > greenLimit) {
                        continue@game
                    } else if (color.contains("blue") && digit > blueLimit) {
                        continue@game
                    }
                }
            }

            result += (index + 1)
        }

        return result
    }

    fun part2(input: List<String>): Int {

        var result = 0
        game@ for (game in input) {

            var redMax = 0
            var greenMax = 0
            var blueMax = 0

            val grabs = game.split(":").last().split(";")
            for (grab in grabs) {

                val colors = grab.split(",")
                for (color in colors) {

                    val digit = color.replace(nonDigitRegex, "").toInt()

                    if (color.contains("red") && digit > redMax) {
                        redMax = digit
                    } else if (color.contains("green") && digit > greenMax) {
                        greenMax = digit
                    } else if (color.contains("blue") && digit > blueMax) {
                        blueMax = digit
                    }
                }
            }

            result += (redMax * greenMax * blueMax)
        }

        return result
    }

    val input = readInput("day-02/input")
    part1(input).println()
    part2(input).println()
}
